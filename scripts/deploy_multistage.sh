#!/bin/bash
# Multi-stage deployment script for blockchain logistics marketplace with multi-chain support

set -e

echo "Starting multi-stage deployment..."

# Stage 1: Build backend services
echo "Building backend services..."
go build -o bin/marketplace ./...

# Stage 2: Run backend tests
echo "Running backend tests..."
go test ./...

# Stage 3: Build frontend
echo "Building frontend..."
cd frontend
npm install
npm run build
cd ..

# Stage 4: Deploy smart contracts to multiple chains
echo "Deploying smart contracts to multiple chains..."
CHAINS=("localhost" "mainnet" "polygon" "arbitrum")
cd hardhat-tests
npm install
npx hardhat compile

for CHAIN in "${CHAINS[@]}"
do
  echo "Deploying to $CHAIN..."
  OWNER_ADDRESS=$OWNER_ADDRESS npx hardhat run --network $CHAIN scripts/deploy_upgradeable.js
done
cd ..

# Stage 5: Start backend server
echo "Starting backend server..."
./bin/marketplace &

# Stage 6: Start frontend server
echo "Starting frontend server..."
cd frontend
npm start &

echo "Multi-stage deployment completed successfully."
