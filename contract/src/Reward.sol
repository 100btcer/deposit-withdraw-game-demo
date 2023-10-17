// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "openzeppelin-contracts/contracts/access/Ownable.sol";
import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/MerkleProof.sol";
import "openzeppelin-contracts/contracts/utils/math/SafeMath.sol";
// import "merkle/merkle.sol";

// 奖励领取合约
contract Reward is Ownable {

     using SafeMath for uint;

    //默克尔树根
    bytes32 public merkleRoot;

    //领取记录
    mapping(address => mapping(address => mapping(uint256 => bool))) public isClaimed;  //用户地址，token合约地址，uint256代表中心化数据的id

    //发奖地址
    address private giveOutAddress;

    constructor() {
        giveOutAddress = msg.sender;
    }

    //判断发奖地址
    modifier checkGiveOutAddress() {
        require(msg.sender == giveOutAddress, 'FiveOutAddress invalid!');
        _;
    }

    //设置默克尔树根
    function setMerkleRoot(bytes32 _merkleRoot) public onlyOwner {
        merkleRoot = _merkleRoot;
    }
    
    //领取token
    function claim(address _token,uint256 _amount,uint256 _id, bytes32[] calldata _proofs) public {
        require(isClaimed[msg.sender][_token][_id] == false,"Already claimed!");

        bytes32 _node = keccak256(abi.encodePacked(msg.sender,_token, _amount,_id));
        require(MerkleProof.verify(_proofs, merkleRoot, _node), "Validation failed!");

        isClaimed[msg.sender][_token][_id] = true;
        require(
            IERC20(_token).transfer(msg.sender, _amount),
            'Transfer failed!'
        );
    }

    //系统发奖
    function sysTransfer(address _token,uint256 _amount,uint256 _id) public checkGiveOutAddress {
        require(isClaimed[msg.sender][_token][_id] == false,"Already claimed!");
        isClaimed[msg.sender][_token][_id] = true;
        require(
            IERC20(_token).transfer(msg.sender, _amount),
            'Transfer failed!'
        );
    }

    //设置发奖地址
    function setGiveOutAddress(address _addr) public onlyOwner {
        giveOutAddress = _addr;
    }

    //获取发奖地址
    function getGiveOutAddress() public view returns(address) {
        return giveOutAddress;
    }
}
