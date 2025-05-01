#!/bin/bash
# Script to verify deployed smart contracts on Etherscan using Hardhat

# Usage: ./verify_contract.sh <contract_address> <network>
# Example: ./verify_contract.sh 0x1234567890abcdef mainnet

CONTRACT_ADDRESS=$1
NETWORK=$2

if [ -z "$CONTRACT_ADDRESS" ] || [ -z "$NETWORK" ]; then
  echo "Usage: $0 <contract_address> <network>"
  exit 1
fi

npx hardhat verify --network $NETWORK $CONTRACT_ADDRESS
