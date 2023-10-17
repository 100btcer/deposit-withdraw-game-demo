MM合约：
0x519fFF7186C9fdE86eC6e26F73E60946fc57dc04

0xD3dd314272943453DffdC30bF7729dcC102E8D9A


0x744686824Adc9121e19e7B106053753b4eCe66E9
lp代币转入合约了，可以用来作为奖励

0xc06Dd89786c2e93f85a24CE40c3FA872CCb49723
正式测试

0x570146345596e97ADdcAd2B46ceCD2d94433c1f5
小数点没问题

range链MM合约：
0x431E7d1189e2453cB4db6105C152B98Fe4cf2aaC

range链Reward代币合约：
0xAa09f16231661bc9c73f997b996bA071ab8c1193

range链Lp代币合约：
0x2D2E7dC4b204f6b24eF0Cb719D473757651488ed


# 部署MM合约
forge create --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥 src/MM.sol:MM

# 部署LP代币
forge create --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥 src/LpToken.sol:LpToken

# 授权合约操作LP
cast send 0x2D2E7dC4b204f6b24eF0Cb719D473757651488ed "approve(address spender, uint256 amount)" 0xEB7791dDE771d7E769d71Fb0bd97408045C06F77 1000000000000000000000000 --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --legacy --private-key 私钥

# 质押LP
cast send 0xEB7791dDE771d7E769d71Fb0bd97408045C06F77 "deposit(uint256 _pid,uint256 _amount)" 8 8 --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥

# 测试领奖
cast send 0xEB7791dDE771d7E769d71Fb0bd97408045C06F77 "claim(address _token,uint256 _amount,uint256 _id, bytes32[] calldata _proofs)" 0xAa09f16231661bc9c73f997b996bA071ab8c1193 0x3e8 646 ["0x358ad291d2ab3225c359167df65b8c3b400b109a444bf47abbd37cf2937676c1","0x424269a30bfcab2d23023bd20b3385afdb85a0893cd460466bd5af36a2db1efa"] --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥

# 赎回
cast send 0xEB7791dDE771d7E769d71Fb0bd97408045C06F77 "withdraw()" --legacy --rpc-url https://robin.rangersprotocol.com/api/jsonrpc --private-key 私钥