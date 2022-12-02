#!/bin/bash

# Clone the submodules
git submodule update --init --recursive

# Install subgraph in datalayr_mantle an come back to our repo
cd ../../../datalayr-mantle/subgraph && yarn && cd ../../ops/docker/eigen-da

# Create necessary .env files
echo '
# EIGENLAYER DEPLOYMENT
RPC_URL=http://18.216.241.146:31000
PRIVATE_KEY=0x54b9e633e01bc3a105561e25a473671fbb6ba21ff5842ddc9b62760c6f508f71

# CHAIN CONFIG
GETH_CHAINID=40525
GETH_UNLOCK_ADDRESS=3aa273f6c6df3a8498ebdcdd241a2575e08cde64

#EXPERIMENT IDENTIFIER
ENVIRONMENT="dev"

# # GLOBAL CONFIG
HOSTNAME="192.168.144.1"
TIMEOUT=3000ms
DLR_ADDRESS=1
CHAIN_PROVIDER=http://${HOSTNAME}:8545
CHAIN_ID=40525
GRAPH_PROVIDER=http://${HOSTNAME}:8000/subgraphs/name/datalayr
G1_PATH=data/kzg/g1.point
G2_PATH=data/kzg/g2.point
SRS_TABLE_PATH=data/kzg/SRSTables
ORDER=3000
CHALLENGE_ORDER=3000' > ./cluster/.env

echo '
GETH_CHAINID=40525
GETH_UNLOCK_ADDRESS=3aa273f6c6df3a8498ebdcdd241a2575e08cde64
EXPERIMENT=""
HOST_IP=0.0.0.0' > ./.env

# Clone g1, g2 points from s3
cd ./datalayr
mkdir -p ./data/kzg
wget --no-check-certificate --no-proxy https://datalayr-testnet.s3.amazonaws.com/g1.point.3000 -O ./data/kzg/g1.point
wget --no-check-certificate --no-proxy https://datalayr-testnet.s3.amazonaws.com/g2.point.3000 -O ./data/kzg/g2.point

# Run setup script for geth-node
cd ../../../../datalayr-mantle/ && make setup-geth