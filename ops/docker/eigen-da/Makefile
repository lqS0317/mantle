
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

# run deploy script
	@echo "Deploying contracts and subgraph"
	cd ../../../datalayr-mantle/integration/ && go run ./scripts

# set ${DLR_ADDRESS} environmental variable
	@echo "exporting address of DLRepository"
	export DLR_ADDRESS="$(file < ../../../datalayr-mantle/contracts/eignlayr-contracts/data/dlRepository.addr)" && cd ./datalayr && docker compose up
