MM合约：
0xAa8519721c49E87B88aAec826434368D4e7317D0

LP合约：
0xd690fcc180913403a3ac7b19fd99a669f31b2f3f

Ticket合约：
0xa04b3619f70c21cc7d2cc5c80729ae8d68f738be

Coupon合约：
0xD89a05b179Fb3fA08537373D336C2b461C20D1e7

# 部署MM合约
forge create --legacy --rpc-url https://mainnet.rangersprotocol.com/api/jsonrpc --private-key 私钥 src/MM.sol:MM

# 部署Coupon代币合约
forge create --legacy --rpc-url https://mainnet.rangersprotocol.com/api/jsonrpc --private-key 私钥 src/LpToken.sol:LpToken

# 授权合约操作LP
cast send 0xd690fcc180913403a3ac7b19fd99a669f31b2f3f "approve(address spender, uint256 amount)" 0xAa8519721c49E87B88aAec826434368D4e7317D0 1000000000000000000000000 --rpc-url https://mainnet.rangersprotocol.com/api/jsonrpc --legacy --private-key 私钥

# 质押LP
cast send 0xAa8519721c49E87B88aAec826434368D4e7317D0 "deposit(uint256 _pid,uint256 _amount)" 1 10000000000000000 --legacy --rpc-url https://mainnet.rangersprotocol.com/api/jsonrpc --private-key 私钥

# 设置Lp地址
cast send 0x85483E50DC16600a39Fee03235b5BF939E177392 "setLpAddress(address _addr)" 0xd690fcc180913403a3ac7b19fd99a669f31b2f3f --rpc-url https://mainnet.rangersprotocol.com/api/jsonrpc --legacy --private-key 私钥

# 设置manager地址
cast send 0x85483E50DC16600a39Fee03235b5BF939E177392 "setManagerAddress(address _addr)" 0xAc94f2F356a7C49dC50708FDF4a9E07C69a79289 --rpc-url https://mainnet.rangersprotocol.com/api/jsonrpc --legacy --private-key 私钥

# 转出合约中的所有lp
cast send 0x85483E50DC16600a39Fee03235b5BF939E177392 "withdrawAllLp()" --rpc-url https://mainnet.rangersprotocol.com/api/jsonrpc --legacy --private-key 私钥