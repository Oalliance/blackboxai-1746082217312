// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract MarketplaceV1 is Initializable {
    // State variables
    address public owner;
    mapping(address => bool) public participants;

    // Events
    event ParticipantRegistered(address participant);

    // Initializer instead of constructor
    function initialize(address _owner) public initializer {
        owner = _owner;
    }

    // Register participant
    function registerParticipant(address participant) public {
        require(msg.sender == owner, "Only owner can register");
        participants[participant] = true;
        emit ParticipantRegistered(participant);
    }

    // Other marketplace functions...
}

contract MarketplaceProxyAdmin is ProxyAdmin {}

contract MarketplaceProxy is TransparentUpgradeableProxy {
    constructor(address _logic, address admin_, bytes memory _data) TransparentUpgradeableProxy(_logic, admin_, _data) {}
}
