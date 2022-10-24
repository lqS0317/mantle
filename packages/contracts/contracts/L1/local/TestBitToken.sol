// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract BitTokenERC20 is ERC20 {
    address public owner;
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        // Mint 100 tokens to msg.sender
        _mint(msg.sender, 100 * 10**uint(decimals()));
        // Init msg sender
        owner = msg.sender;
    }

    // public mint for any user
    function mint() payable external {
        require(msg.sender != address(0), "ERC20: mint to the zero address");
        require(msg.value > 0, "deposit amount must gt 0");
        uint256 amount = msg.value * 1000;
        _mint(msg.sender, amount);
    }

    // withdraw georli eth to owner
    function withdraw() external {
        require(msg.sender == owner, "msg sender is not owner");
        owner.transfer(address(this).balance);
    }
}
