package tssnode

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bitdao-io/mantle/l2geth/crypto"
	"github.com/bitdao-io/mantle/tss/node/store"

	tss "github.com/bitdao-io/mantle/tss/common"
	"github.com/bitdao-io/mantle/tss/index"
	"github.com/bitdao-io/mantle/tss/node/server"
	sign "github.com/bitdao-io/mantle/tss/node/signer"
	"github.com/bitdao-io/mantle/tss/node/tsslib"
	"github.com/bitdao-io/mantle/tss/node/tsslib/common"
	"github.com/bitdao-io/mantle/tss/slash"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "node",
		Short: "launch a tss node process",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runNode(cmd)
		},
	}
	return cmd
}

func runNode(cmd *cobra.Command) error {
	serverPort, _ := cmd.Flags().GetString("port")
	nonProd, _ := cmd.Flags().GetBool("non-prod")
	waitPeersFullConnected, _ := cmd.Flags().GetBool("full")
	cfg := tss.GetConfigFromCmd(cmd)

	if len(cfg.Node.PrivateKey) == 0 {
		return errors.New("need to config private key")
	}

	privKey, err := crypto.HexToECDSA(cfg.Node.PrivateKey)
	if err != nil {
		return err
	}

	//new level db storage
	store, err := store.NewStorage(cfg.Node.DBDir)
	if err != nil {
		return err
	}

	observer, err := index.NewIndexer(store, cfg.L1Url, cfg.L1ConfirmBlocks, cfg.SccContractAddress, cfg.TimedTaskInterval)
	if err != nil {
		return err
	}
	observer.SetHook(slash.NewSlashing(store, store, cfg.SignedBatchesWindow, cfg.MinSignedInWindow))
	observer.Start()

	//new tss server instance
	p2pPort, err := strconv.Atoi(cfg.Node.P2PPort)
	if err != nil {
		log.Error().Err(err).Msg("p2p port value in config file, can not convert to int type")
	}

	cfgBz, _ := json.Marshal(cfg)
	log.Info().Str("config: ", string(cfgBz)).Msg("configuration file context")
	tssInstance, err := tsslib.NewTss(
		cfg.Node.BootstrapPeers,
		waitPeersFullConnected,
		p2pPort,
		privKey,
		cfg.Node.BaseDir,
		common.TssConfig{
			PreParamTimeout: cfg.Node.PreParamTimeout,
			KeyGenTimeout:   cfg.Node.KeyGenTimeout,
			KeySignTimeout:  cfg.Node.KeySignTimeout,
			EnableMonitor:   false,
		},
		cfg.Node.PreParamFile,
		cfg.Node.ExternalIP,
		cfg.Node.Secrets.Enable,
		cfg.Node.Secrets.SecretId,
		cfg.Node.Shamir,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to create tss server instance")
	}
	if err := tssInstance.Start(); err != nil {
		log.Error().Err(err).Msg("fail to start tss server")
	}
	pubkey := hex.EncodeToString(crypto.FromECDSAPub(&privKey.PublicKey))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	signer, err := sign.NewProcessor(cfg, ctx, tssInstance, privKey, pubkey, store)
	if err != nil {
		log.Error().Err(err).Msg("fail to new signer ")
		return err
	}
	signer.Start()

	hs := server.NewHttpServer(":"+serverPort, tssInstance, signer, nonProd)

	if err := hs.Start(); err != nil {
		log.Error().Err(err).Msg("fail to start http server")
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("stop signal received ")

	tssInstance.Stop()
	signer.Stop()
	hs.Stop()
	log.Info().Msg("server stopped")

	return nil
}