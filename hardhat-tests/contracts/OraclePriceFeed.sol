// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract OraclePriceFeed is Ownable {
    struct Oracle {
        bool isActive;
        address oracleAddress;
    }

    mapping(address => Oracle) public oracles;
    address[] public oracleList;

    uint256 public minPrice;
    uint256 public maxPrice;

    uint256 public lastUpdateTime;
    uint256 public updateTimeLock; // in seconds

    uint256 public currentPrice;

    // Mapping of oracle address to their last submitted price
    mapping(address => uint256) public lastPrices;

    // Event emitted when price is updated
    event PriceUpdated(uint256 newPrice, uint256 timestamp);

    // Event emitted when oracle is added or removed
    event OracleStatusChanged(address oracle, bool isActive);

    constructor(uint256 _minPrice, uint256 _maxPrice, uint256 _updateTimeLock) {
        minPrice = _minPrice;
        maxPrice = _maxPrice;
        updateTimeLock = _updateTimeLock;
    }

    // Add or update oracle status
    function setOracle(address _oracle, bool _isActive) external onlyOwner {
        if (!oracles[_oracle].isActive && _isActive) {
            oracleList.push(_oracle);
        }
        oracles[_oracle] = Oracle(_isActive, _oracle);
        emit OracleStatusChanged(_oracle, _isActive);
    }

    // Remove oracle from list (deactivate)
    function removeOracle(address _oracle) external onlyOwner {
        oracles[_oracle].isActive = false;
        emit OracleStatusChanged(_oracle, false);
    }

    // Submit price from oracle with signature verification
    // For simplicity, assume msg.sender is the oracle address
    function submitPrice(uint256 _price) external {
        require(oracles[msg.sender].isActive, "Not an active oracle");
        require(block.timestamp >= lastUpdateTime + updateTimeLock, "Update time lock active");
        require(_price >= minPrice && _price <= maxPrice, "Price out of bounds");

        lastPrices[msg.sender] = _price;

        // Aggregate prices from all active oracles
        uint256 sum = 0;
        uint256 count = 0;
        for (uint256 i = 0; i < oracleList.length; i++) {
            address oracleAddr = oracleList[i];
            if (oracles[oracleAddr].isActive && lastPrices[oracleAddr] > 0) {
                sum += lastPrices[oracleAddr];
                count++;
            }
        }
        require(count > 0, "No prices submitted");

        uint256 aggregatedPrice = sum / count;

        // Update current price and timestamp
        currentPrice = aggregatedPrice;
        lastUpdateTime = block.timestamp;

        emit PriceUpdated(currentPrice, lastUpdateTime);
    }

    // Get current aggregated price
    function getCurrentPrice() external view returns (uint256) {
        return currentPrice;
    }

    // Set min and max price thresholds
    function setPriceThresholds(uint256 _minPrice, uint256 _maxPrice) external onlyOwner {
        require(_minPrice < _maxPrice, "Invalid thresholds");
        minPrice = _minPrice;
        maxPrice = _maxPrice;
    }

    // Set update time lock duration
    function setUpdateTimeLock(uint256 _updateTimeLock) external onlyOwner {
        updateTimeLock = _updateTimeLock;
    }
}
