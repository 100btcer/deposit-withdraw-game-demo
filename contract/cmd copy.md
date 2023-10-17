
range链MM合约：
0x1deFB5b39B9dACBF75479347B34Dd8d7011f6901

ticket合约：
0x2c3b9ce8a9b40ec39a7b0cc4d3ccf2374b2fb4b8

range链Reward代币合约：
0xAa09f16231661bc9c73f997b996bA071ab8c1193

range链Lp代币合约：
0x2D2E7dC4b204f6b24eF0Cb719D473757651488ed


# 部署MM合约
forge create --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥 src/MM.sol:MM

# 部署LP代币
forge create --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥 src/LpToken.sol:LpToken

# 授权合约操作LP
cast send 0x2D2E7dC4b204f6b24eF0Cb719D473757651488ed "approve(address spender, uint256 amount)" 0x1deFB5b39B9dACBF75479347B34Dd8d7011f6901 1000000000000000000000000 --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --legacy --private-key 私钥

# 系统发奖
cast send 0x2D2E7dC4b204f6b24eF0Cb719D473757651488ed "sysTransfer(address _token,uint256 _amount,uint256 _id)" 0x1deFB5b39B9dACBF75479347B34Dd8d7011f6901 1000000000000000000000000 --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --legacy --private-key 私钥

# 质押LP
cast send 0x1deFB5b39B9dACBF75479347B34Dd8d7011f6901 "deposit(uint256 _pid,uint256 _amount)" 7 7000000000000000000000 --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥

# 测试领奖
cast send 0x431E7d1189e2453cB4db6105C152B98Fe4cf2aaC "claim(address _token,uint256 _amount,uint256 _id, bytes32[] calldata _proofs)" 0xAa09f16231661bc9c73f997b996bA071ab8c1193 0x3635c9adc5dea00000 648 ["0x25e71c306cb80287ee00b61607ac2320e62b6b69ded97440eb62911517157b58","0x94f537342fc1ef728bd65376e304c5f9a0bdb3bab1394583a058686ee3d397e0","0x68b8e5dc1d4dd2c2e05e769025d3b5d25caf3ed399ba8518ece88cbb2f7594aa","0x4ed930ad7419c7df2989f68c7a4983d9f4899cd14488f4a11b060f6522901a95","0xcec97a7a28138a232b93207d89055bbe379c1ecb3fae6eb077f35b7d5cf973d1","0xd7f317e6e37a983ce7ebccc3a06a1b1c5964fd18fc5da18a75ccfdf512be08fb","0x6d685d384cda391755cc46ba11a60cb2f01324240a6c51b1193f6f173a633516","0x52df1b82461fb7a15c79debc32a752218c6bad2855dc794a51226b33edd8021e","0x59aa1667fdcf1a911f8f91cbbbd217606a4263562efe820b3083b9d59c09a64b","0x846859275ecba37f3bd23c4e0408e836bc06ccf6cd64511b09a5d522544ac60a"] --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥

# 赎回
cast send 0x431E7d1189e2453cB4db6105C152B98Fe4cf2aaC "withdraw()" --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥



cast send 0x4acbc2b1235da10c28e09fd461384bae2767e24e "burn(address account, uint256 amount)" 0x878ADa7AF22A35Afc5D7c4f0AF57cdE44139a91c 20000000000000000000 --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥



# 销毁ticket
cast send 0x2c3b9ce8a9b40ec39a7b0cc4d3ccf2374b2fb4b8 "burn(address account, uint256 amount)" 0x878ADa7AF22A35Afc5D7c4f0AF57cdE44139a91c 40000000000000000000 --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥

# 提取所有代币
cast send 0x1deFB5b39B9dACBF75479347B34Dd8d7011f6901 "withdrawToken(address _token)" 0xAa09f16231661bc9c73f997b996bA071ab8c1193 --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥