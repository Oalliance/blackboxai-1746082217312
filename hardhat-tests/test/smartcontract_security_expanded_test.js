const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("SmartContract Security Expanded Tests", function () {
  let SmartContract;
  let smartContract;
  let owner;
  let addr1;
  let addr2;

  beforeEach(async function () {
    SmartContract = await ethers.getContractFactory("UpgradeableMarketplace");
    [owner, addr1, addr2] = await ethers.getSigners();
    smartContract = await SmartContract.deploy();
    await smartContract.deployed();
  });

  it("Should deploy the smart contract", async function () {
    expect(smartContract.address).to.properAddress;
  });

  it("Should allow minting tokens", async function () {
    await expect(smartContract.mintToken(addr1.address, 1000))
      .to.emit(smartContract, "Transfer")
      .withArgs(ethers.constants.AddressZero, addr1.address, 1000);
  });

  it("Should allow transferring tokens", async function () {
    await smartContract.mintToken(owner.address, 1000);
    await smartContract.transferToken(owner.address, addr1.address, 500);
    const balance = await smartContract.balanceOf(addr1.address);
    expect(balance).to.equal(500);
  });

  it("Should prevent transferring more tokens than balance", async function () {
    await smartContract.mintToken(owner.address, 100);
    await expect(
      smartContract.transferToken(owner.address, addr1.address, 200)
    ).to.be.revertedWith("Insufficient balance");
  });

  it("Should prevent unauthorized minting", async function () {
    await expect(
      smartContract.connect(addr1).mintToken(addr1.address, 1000)
    ).to.be.revertedWith("Ownable: caller is not the owner");
  });

  it("Should allow owner to pause and unpause contract", async function () {
    await smartContract.pause();
    expect(await smartContract.paused()).to.be.true;
    await smartContract.unpause();
    expect(await smartContract.paused()).to.be.false;
  });

  it("Should prevent token transfers when paused", async function () {
    await smartContract.mintToken(owner.address, 1000);
    await smartContract.pause();
    await expect(
      smartContract.transferToken(owner.address, addr1.address, 100)
    ).to.be.revertedWith("Pausable: paused");
  });

  it("Should allow only owner to upgrade contract", async function () {
    const NewContract = await ethers.getContractFactory("UpgradeableMarketplace");
    const newContract = await NewContract.deploy();
    await newContract.deployed();

    await expect(
      smartContract.connect(addr1).upgradeTo(newContract.address)
    ).to.be.revertedWith("Ownable: caller is not the owner");

    await smartContract.upgradeTo(newContract.address);
    expect(await smartContract.getImplementation()).to.equal(newContract.address);
  });

  // Add more security-related tests as needed
});
