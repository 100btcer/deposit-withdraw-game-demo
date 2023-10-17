// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "openzeppelin-contracts/contracts/access/Ownable.sol";

import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

// 用来部署lp代币
contract TicketToken is ERC20,Ownable {
    constructor() ERC20("Lucky Coin", "LC") {

    }

    //禁用转账功能
    function transfer(address to, uint256 amount) public virtual override returns (bool) {
        revert();
    }

    //禁用
    function transferFrom(address from, address to, uint256 amount) public virtual override returns (bool) {
        revert();
    }

    //销毁
    function burn(address account, uint256 amount) public virtual {
        _burn(account, amount);
    }

    //mint
    function mintTicketToken(address account, uint256 amount) public onlyOwner {
        _mint(account, amount);
    }
}
