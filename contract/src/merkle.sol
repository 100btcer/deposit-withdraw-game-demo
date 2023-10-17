// SPDX-License-Identifier: MIT
pragma solidity >=0.8.13;
 
import "openzeppelin-contracts/contracts/utils/cryptography/MerkleProof.sol";
 
 
contract MerkleVerify {
 
    //记录当前默克尔树
    bytes32 public merkleRoot;

    constructor() {
        merkleRoot = 0xbe27f4c175c0bbc6394076c150619580a3ef12fb294ccb31195ce56dfcbf2a9d;
    }
 
    //限制每个默克尔树，每个用户只能提现一次
    mapping(bytes32 => mapping(uint256 => uint256)) private merkleRootRecord;
 
    //计算 默克尔树。
    function withdraw(address _token,uint256 _amount,uint256 _id, bytes32[] calldata _proofs)
    external
    {
 
        //每个默克尔树，每个用户只能提现一次
        // require(!_isClaimed(_index), "Multiple withdrawal is prohibited");
 
        //验证默克尔树
        bytes32 _node = keccak256(abi.encodePacked(msg.sender,_token, _amount,_id));
        require(MerkleProof.verify(_proofs, merkleRoot, _node), "Validation failed");
        // _setClaimed(_index);
 
    }
 
    //验证这个证明是否用过
    function _isClaimed(uint256 index) public view returns (bool) {
        uint256 claimedWordIndex = index / 256;
        uint256 claimedBitIndex = index % 256;
        uint256 claimedWord = merkleRootRecord[merkleRoot][claimedWordIndex];
        uint256 mask = (1 << claimedBitIndex);
        return claimedWord & mask == mask;
    }
 
    //添加用过的证明
    function _setClaimed(uint256 index) private {
        uint256 claimedWordIndex = index / 256;
        uint256 claimedBitIndex = index % 256;
        merkleRootRecord[merkleRoot][claimedWordIndex] = merkleRootRecord[merkleRoot][claimedWordIndex] | (1 << claimedBitIndex);
    }
}