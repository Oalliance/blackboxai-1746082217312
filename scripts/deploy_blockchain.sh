#!/bin/bash
# Deployment script for blockchain smart contracts and marketplace

set -e

echo "Starting blockchain smart contract deployment..."

# Set owner address environment variable before running this script
if [ -z "$OWNER_ADDRESS" ]; then
  echo "Error: OWNER_ADDRESS environment variable is not set."
  exit 1
fi

# Navigate to hardhat tests directory
cd hardhat-tests

# Install dependencies
npm install

# Compile smart contracts
npx hardhat compile

# Deploy upgradeable proxy contract
OWNER_ADDRESS=$OWNER_ADDRESS npx hardhat run --network mainnet scripts/deploy_upgradeable.js

echo "Blockchain smart contract deployment completed."
