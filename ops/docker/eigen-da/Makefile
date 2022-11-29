
# set up local geth build and import local private keys from datalayr
setup-geth:
	cd ../../../datalayr-mantle/integration/geth-node && sudo rm -rf data && ./run.sh setup 40 ./secret/private-keys.txt ./secret/geth-account-password

.PHONY: deploy-eigenda-local
deploy-eigenda-local:
# run setup, compile contracts
	@echo "RUNNING SETUP..."
	./setup.sh
	@echo "SETUP COMPLETE"
# compile contracts
	@echo "COMPILING CONTRACTS..."
	cd ../../../datalayr-mantle/contracts && ./compile.sh compile-el && ./compile.sh compile-rollup
	@echo "CONTRACTS COMPILED"

# start nodes
	@echo "STARTING GETH DOCKER CONTAINERS"
	cd ../../../datalayr-mantle/integration && docker compose up -d
	@echo "CONTAINERS STARTED"

# # deploy contracts
# 	@echo "DEPLOYING CONTRACTS"
# 	cd ../../../datalayr-mantle/contracts/eignlayr-contracts/script && forge script Deployer.s.sol --rpc-url http://0.0.0.0:8545 --skip-simulation --broadcast --private-key 0x1a5843a0c64beab41ae8b1059f30db49339019d9b31a42f61577aae58976ee68 --legacy --slow

# deploy script
	@echo "Deploying contracts and subgraph"
	cd ../../../datalayr-mantle/integration/ && go run ./scripts

#



# # Build

# binaries: maybe-compile-el maybe-compile-rollup
# 	$(MAKE) -C ./dl-disperser dl-disperser
# 	$(MAKE) -C ./dl-node dl-node
# 	$(MAKE) -C ./dl-retriever dl-retriever
# 	$(MAKE) -C ./middleware/rollup-example/challenger

# .PHONY: images
# images: maybe-compile-el maybe-compile-rollup
# 	docker compose build

# .PHONY: maybe-compile-el
# maybe-compile-el:
# 	./on-change.sh contracts/eignlayr-contracts/src data/eigenlayer_hash "make compile-el"

# .PHONY: compile-el
# compile-el:
# 	cd contracts && ./compile.sh compile-el

# .PHONY: maybe-compile-rollup
# maybe-compile-rollup:
# 	./on-change.sh contracts/datalayr-rollup-example-contracts/src data/rollup_example_hash "make compile-rollup"

# .PHONY: compile-rollup
# compile-rollup:
# 	cd contracts && ./compile.sh compile-rollup


# # Unit Test
# .PHONY: unit-test
# unit-test:
# 	ginkgo common/crypto/bls


# # Local Infrastructure



# # setup-geth:
# # 	cd integration/geth-node && sudo rm -rf data && ./setup.sh 40 ./secret/private-keys.txt

# update-host:
# 	cd integration && ./update-host.sh

# deploy-chain: destroy-chain update-host
# 	cd integration && docker compose up -d
# 	cd integration && ./wait-for http://0.0.0.0:8000 -- echo "GraphQL up" \
# 		&& ./wait-for http://0.0.0.0:8545 -- echo "Geth up"

# deploy-chain-n-explorer: destroy-chain-n-explorer update-host
# 	cd integration && docker compose -f docker-compose-with-explorer.yml up -d
# 	cd integration && ./wait-for http://0.0.0.0:8000 -- echo "GraphQL up" \
# 		&& ./wait-for http://0.0.0.0:8545 -- echo "Geth up" \
# 		&& ./wait-for http://0.0.0.0:4000 -- echo "Explorer up"

# destroy-chain:
# 	cd integration \
# 		&& docker compose down

# destroy-chain-n-explorer:
# 	cd integration \
# 		&& docker compose -f docker-compose-with-explorer.yml down

# # Local Deployment

# new-experiment:
# 	cd integration && ./new-exp.sh

# deploy-experiment: binaries new-experiment
# 	cd integration && go run ./scripts

# deploy-nodes: destroy-nodes
# 	cd integration && ./bin.sh start-detached

# destroy-nodes:
# 	cd integration && ./bin.sh stop

# deploy-nodes-docker: destroy-nodes-docker
# 	cd integration && ./docker.sh start

# destroy-nodes-docker:
# 	cd integration && ./docker.sh stop

# # Integration Test

# test-new-infra: binaries deploy-chain new-experiment
# 	ginkgo ./integration

# test-new-contract: binaries new-experiment
# 	ginkgo ./integration


# deploy-rollup:
# 	cd integration && ./new-exp.sh config.rollup.yaml
# 	cd integration && go run ./scripts

# test-new-nodes: binaries
# 	ginkgo ./integration -v -- --new-experiment=false

