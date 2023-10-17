// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "openzeppelin-contracts/contracts/access/Ownable.sol";

import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

import "openzeppelin-contracts/contracts/utils/math/SafeMath.sol";
import "./TicketToken.sol";
import "./Reward.sol";

contract MM is Ownable,Reward {

    using SafeMath for uint;

    //矿池信息
    struct PoolInfo {
        uint256 id;  //矿池id
        uint256 lockDay;  //锁仓天数
        uint256 rewardRatio;  //奖励系数
        uint256 decimal;  //小数精度，例如精度为100，则0.25按照整数25，计算ticket时会除以100
    }

    //存款信息
    struct DepositInfo {
        uint256 poolId;  //奖池id
        uint256 amount;  //存款数额
        uint256 unlockTimestemp;  //解锁到期时间
        bool isWithdraw;  //是否取出
    }

    //矿池
    mapping(uint256 => PoolInfo) private poolInfo;

    //用户存款
    mapping(address => DepositInfo[]) private userInfo;

    //lp代币合约
    address private lpAddress = 0xd690FcC180913403A3AC7B19fD99a669f31B2f3F;
    
    //管理员 需要硬编码到合约中，尽量不要是合约发布者地址
    address private manager;

    //ticket合约地址
    address private ticketTokenAddress;

    constructor() {
        manager = msg.sender;
        initPool();
        //创建ticketToken合约
        TicketToken ticketToken = new TicketToken();
        ticketTokenAddress = address(ticketToken);
    }

    uint private unlocked = 1;
    modifier lock() {
        require(unlocked == 1, 'LOCKED');
        unlocked = 0;
        _;
        unlocked = 1;
    }

    //添加矿池
    function addPool(uint256 _pid,uint256 lockDay,uint256 rewardRatio,uint256 decimal) public onlyOwner {
        poolInfo[_pid] = PoolInfo({
            id:_pid,
            lockDay:lockDay,
            rewardRatio:rewardRatio,
            decimal:decimal
        });
    }

    //初始化矿池
    function initPool() private {
        addPool(1,7,989088,9); //7天
        addPool(2,14,1149292,9); //14天
        addPool(3,30,1323427,9); //30天
        addPool(4,90,1741351,9); //90天
        addPool(5,180,1448804,9); //180天
        addPool(6,360,961226,9); //365天
        addPool(7,1,961226,9);  //测试1填
    }

    //获取每个矿池数据
    function getPoolById(uint256 _pid) public view returns(PoolInfo memory) {
        require(poolInfo[_pid].id > 0,"Pool invalid!");
        PoolInfo memory poolInfo = poolInfo[_pid];
        return poolInfo;
    }

    //铸造ticketToken
    // function _mintVeLP(address _addr,uint256 _num) internal {
    //     ITicketToken(ticketTokenAddress).mintTicketToken(_addr, _num);
    // }

    //存款
    function deposit(uint256 _pid,uint256 _amount) public {
        require(poolInfo[_pid].lockDay > 0,"Pool invalid!");
        
        DepositInfo[] storage user = userInfo[msg.sender];
        uint256 lockDay = poolInfo[_pid].lockDay;
        uint256 rewardRatio = poolInfo[_pid].rewardRatio;
        uint256 lockTimestamp = block.timestamp.add(lockDay.mul(86400));
        uint256 poolDecimal = poolInfo[_pid].decimal;
        user.push(DepositInfo({
            poolId:_pid,
            amount:_amount,
            unlockTimestemp:lockTimestamp,
            isWithdraw:false
        }));
        uint256 rewardAmount = _amount.mul(lockDay).mul(rewardRatio).div(10**poolDecimal);
        //划转用户代币
        IERC20(lpAddress).transferFrom(msg.sender,address(this),_amount);
        // 铸造ticket代币，需要计算比例
        uint256 mintTicketAmount = rewardAmount;
        ITicketToken(ticketTokenAddress).mintTicketToken(msg.sender,mintTicketAmount);

        emit Deposit(msg.sender,_pid,_amount,mintTicketAmount,lockTimestamp);
    }

    //取款
    function withdraw() public {
        DepositInfo[] storage depositData = userInfo[msg.sender];
        uint256 nowTime = block.timestamp;
        uint256 validAmount;
        uint256 length = depositData.length;
        for (uint i = 0;i < length;i++) {
            if (nowTime >= depositData[i].unlockTimestemp && !depositData[i].isWithdraw) {
                    validAmount = validAmount.add(depositData[i].amount);
                    depositData[i].isWithdraw = true;
                }
        }
        require(validAmount > 0,"ValidAmount insufficient");
        IERC20(lpAddress).transfer(msg.sender,validAmount);
        emit Withdraw(msg.sender,validAmount);
    }

    //获取lp余额
    function getLpBalance() public view returns(uint256) {
        uint256 balance = IERC20(lpAddress).balanceOf(address(this));
        return balance;
    }

    //一次性取出所以LP代币
    function withdrawAllLp() public onlyOwner {
        uint256 balance = IERC20(lpAddress).balanceOf(address(this));
        require(balance > 0,"Balance insufficient");
        IERC20(lpAddress).transfer(manager,balance);
    }

    //获取用户在某个池子质押数量
    function getDepositByPoolId(uint256 _pid) public view returns(uint256,uint256,uint256) {
        uint256 num;
        uint256 amount;
        uint256 validWithdrawAmount;
        DepositInfo[] memory depositData = userInfo[msg.sender];
        uint256 length = depositData.length;
        uint256 nowTime = block.timestamp;
        for (uint i=0;i<length;i++) {
            if (depositData[i].poolId == _pid) {
                num = num.add(1);
                amount = amount.add(depositData[i].amount);
                if (nowTime >= depositData[i].unlockTimestemp && depositData[i].isWithdraw == false) {
                    validWithdrawAmount = validWithdrawAmount.add(depositData[i].amount);
                }
            }
        }
        return (num,amount,validWithdrawAmount);
    }

    //获取可赎回的lp总额
    function getValidWithdrawLp() public view returns(uint256) {
        uint256 validWithdrawAmount;
        DepositInfo[] memory depositData = userInfo[msg.sender];
        uint depositLen = depositData.length;
        uint256 nowTime = block.timestamp;
        if (depositLen == 0) {
            return 0;
        }
        for (uint i = 0;i < depositLen;i++) {
            if (nowTime >= depositData[i].unlockTimestemp && depositData[i].isWithdraw == false) {
                validWithdrawAmount = validWithdrawAmount.add(depositData[i].amount);
            }
        }
        return validWithdrawAmount;
    }

    //查询合约代币余额
    function balanceOf(address _token) public view returns(uint256) {
        uint256 balance = ERC20(_token).balanceOf(address(this));
        return balance;
    }

    //取出所有代币
    function withdrawToken(address _token) public onlyOwner {
        uint256 balance = ERC20(_token).balanceOf(address(this));
        ERC20(_token).transfer(manager,balance);
    }

    //设置lp代币地址
    function setLpAddress(address _addr) public onlyOwner {
        lpAddress = _addr;
    }

    //获取lpAddress地址
    function getLpAddress() public view returns(address) {
        return lpAddress;
    }

    //设置管理员地址
    function setManagerAddress(address _addr) public onlyOwner {
        manager = _addr;
    }

    //获取管理员地址
    function getManagerAddress() public view returns(address) {
        return manager;
    }

    event Deposit(address indexed user, uint256 indexed pid, uint256 depositAmount,uint256 ticketAmount,uint256 lockTimestamp);
    event Withdraw(address indexed user, uint256 amount);

}

interface ITicketToken {
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event Transfer(address indexed from, address indexed to, uint256 value);

    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
    function burn(address account, uint256 amount) external;
    function decimals() external view returns (uint8);
    function decreaseAllowance(address spender, uint256 subtractedValue) external returns (bool);
    function increaseAllowance(address spender, uint256 addedValue) external returns (bool);
    function mintTicketToken(address account, uint256 amount) external;
    function name() external view returns (string memory);
    function owner() external view returns (address);
    function renounceOwnership() external;
    function symbol() external view returns (string memory);
    function totalSupply() external view returns (uint256);
    function transfer(address to, uint256 amount) external returns (bool);
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
    function transferOwnership(address newOwner) external;
}

