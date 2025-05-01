#!/bin/bash
# Script to compile smart contracts using Hardhat

set -e

echo "Installing dependencies..."
cd hardhat-tests
npm install

echo "Compiling smart contracts..."
npx hardhat compile

echo "Smart contract compilation completed."
