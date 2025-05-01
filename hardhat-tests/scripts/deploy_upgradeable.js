const { ethers, upgrades } = require("hardhat");

async function main() {
  const MarketplaceV1 = await ethers.getContractFactory("MarketplaceV1");
  console.log("Deploying MarketplaceV1...");
  const marketplace = await upgrades.deployProxy(MarketplaceV1, [process.env.OWNER_ADDRESS], { initializer: 'initialize' });
  await marketplace.deployed();
  console.log("MarketplaceV1 deployed to:", marketplace.address);

  // ProxyAdmin and TransparentUpgradeableProxy are handled by OpenZeppelin upgrades plugin
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
