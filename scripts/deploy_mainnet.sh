#!/bin/bash
# Script to deploy the smart contract to the Ethereum mainnet using Hardhat

set -e

if [ -z "$OWNER_ADDRESS" ]; then
  echo "Error: OWNER_ADDRESS environment variable is not set."
  exit 1
fi

echo "Deploying MarketplaceV1 smart contract to Ethereum mainnet..."

# Ensure dependencies are installed
cd hardhat-tests
npm install

# Compile contracts
npx hardhat compile

# Deploy contract using the deploy_upgradeable.js script
OWNER_ADDRESS=$OWNER_ADDRESS npx hardhat run --network mainnet scripts/deploy_upgradeable.js

echo "Deployment completed."
