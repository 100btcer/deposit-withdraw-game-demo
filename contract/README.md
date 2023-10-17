# 部署lp代币合约
forge create --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥 src/ERC20.sol:LPToken

# 部署质押合约
forge create --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥 src/Deposit.sol:Deposit

# 调用铸造
cast send 0xaa09f16231661bc9c73f997b996ba071ab8c1193 "mintVeLP(address _addr,uint256 _num)" 0x878ADa7AF22A35Afc5D7c4f0AF57cdE44139a91c 100 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥


lp合约地址：0xeb7791dde771d7e769d71fb0bd97408045c06f77
质押合约地址：0x60bb2d1603122f0fc79d3be4c27aab8ac09cdb53

# 调用测试铸造velp
cast send 0x0bf4d240db005ada49e62e4f17115781efa5efd4 "testMint(uint256 _amount)" 100 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 授权操作lp
cast send 0xeb7791dde771d7e769d71fb0bd97408045c06f77 "approve(address spender, uint256 amount)" 0x60bb2d1603122f0fc79d3be4c27aab8ac09cdb53 100000000000000000000999999999999999 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 授权操作velp
cast send 0x60bb2d1603122f0fc79d3be4c27aab8ac09cdb53 "approve(address spender, uint256 amount)" 0x60bb2d1603122f0fc79d3be4c27aab8ac09cdb53 100000000000000000000 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 调用存款合约
cast send 0x60bb2d1603122f0fc79d3be4c27aab8ac09cdb53 "deposit(uint256 _depositAmount)" 1000000 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 调用合约取款
cast send 0x60bb2d1603122f0fc79d3be4c27aab8ac09cdb53 "draw(uint256 _drawAmount)" 1000000 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 调用合约查询存款余额
cast send 0xab796520783970ad7e2ffd6d97bed97c8b5df0fa "burn(address account, uint256 amount)" 0x878ADa7AF22A35Afc5D7c4f0AF57cdE44139a91c --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥





-----------------------
# 部署一个无法转账的代币
forge create --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥 src/LpToken.sol:LpToken

# 销毁一个无法转账的代币
cast send 0x076c7dfbdc8b372ed4af46142421d3c9a5c737a9 "burn(address account, uint256 amount)" 0x878ADa7AF22A35Afc5D7c4f0AF57cdE44139a91c 1000 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥


----------------------------------------------------
MM合约：
0xCd28DB4b2128eBa811B8b7F718C9c34f2CE69df0

# ticket合约
0x12047767B9f8CabF9fdCEE1572748A88Db1523f0
cast send 0x12047767B9f8CabF9fdCEE1572748A88Db1523f0 "mintTicketToken(address account, uint256 amount)" 0x878ADa7AF22A35Afc5D7c4f0AF57cdE44139a91c 123456 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥


# 部署MM合约
forge create --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥 src/MM.sol:MM

# 授权合约操作LP
cast send 0xEB7791dDE771d7E769d71Fb0bd97408045C06F77 "approve(address spender, uint256 amount)" 0xCd28DB4b2128eBa811B8b7F718C9c34f2CE69df0 100000000000000000000000 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 质押LP
cast send 0x89B41b393e085bBD5264FA2028C26274e6eb31b6 "deposit(uint256 _pid,uint256 _amount)" 2 40000000000 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 获取用户在某个池的质押
cast send 0x956e3afd2180cce33b82827b0f718b85abe3cdd3 "getDepositByPoolId(uint256 _pid)" 1 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 获取合约lp代币余额
cast send 0x956e3afd2180cce33b82827b0f718b85abe3cdd3 "getLpBalance()" --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 添加池子
cast send 0x956e3afd2180cce33b82827b0f718b85abe3cdd3 "addPool(uint256 _pid,uint256 secondTotal)" 4 60 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 赎回
cast send 0x956e3afd2180cce33b82827b0f718b85abe3cdd3 "withdraw()" --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 销毁ticket
cast send 0x956e3afd2180cce33b82827b0f718b85abe3cdd3 "burn(address account, uint256 amount)" 0x878ADa7AF22A35Afc5D7c4f0AF57cdE44139a91c 1000 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥


# 发行一个奖品token
forge create --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥 src/ERC20.sol:LPToken

# 给合约转账10000个奖品token
cast send 0x0b101adc1a24f7b3e188626b31905dd94c02937e "transfer(address to, uint256 amount)" 0x956e3afd2180cce33b82827b0f718b85abe3cdd3 10000000000000000000000 --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 测试领取奖励


# 设置默克尔树根
cast send 0xCd28DB4b2128eBa811B8b7F718C9c34f2CE69df0 "setMerkleRoot(bytes32 _merkleRoot)" 0xc652ccd5278e7c25c53ec67ca9d2cc6810c321dbab67ab12312deb0b29f2c98f --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥

# 测试领奖
cast send 0xCd28DB4b2128eBa811B8b7F718C9c34f2CE69df0 "claim(address _token,uint256 _amount,uint256 _id, bytes32[] calldata _proofs)" 0xEB7791dDE771d7E769d71Fb0bd97408045C06F77 100 8 ["0xffda24c35ac90eb2d29f04d549403996b1408f53bff094fe9b4b0d156d9142b7"] --rpc-url https://data-seed-prebsc-1-s1.binance.org:8545 --private-key 私钥


0x190


# abi生成golang文件
abigen --abi abi/MM.abi --pkg MM --type MM --out mm.go