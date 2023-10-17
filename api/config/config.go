package config

import (
	"errors"
)

// 池子id和质押天数映射
type PoolItem struct {
	PoolId      int     `json:"pool_id"`      //奖池id
	Days        int     `json:"days"`         //天数
	RewardRatio float64 `json:"reward_ratio"` //奖励系数
}

var PoolConfig = []PoolItem{
	{
		PoolId:      1,
		Days:        7,
		RewardRatio: 1.42,
	},
	{
		PoolId:      2,
		Days:        14,
		RewardRatio: 1.78,
	},
	{
		PoolId:      3,
		Days:        30,
		RewardRatio: 1.83,
	},
	{
		PoolId:      4,
		Days:        90,
		RewardRatio: 2.77,
	},
	{
		PoolId:      5,
		Days:        180,
		RewardRatio: 2.08,
	},
	{
		PoolId:      6,
		Days:        365,
		RewardRatio: 1.38,
	},
}

// LP代币合约
const LpContract = "0xEB7791dDE771d7E769d71Fb0bd97408045C06F77"

// 扭蛋合约
const GachaContract = "0x744686824Adc9121e19e7B106053753b4eCe66E9"

// 单次抽奖消耗ticket
const PlayGameCostTicket = 10

// 奖品排序
var TokenPrizeSortConfig = []string{"MIX", "RPG", "USDT", "碎片"}
var NFTPrizeSoftConfig = []string{"龙蛋（普通）", "龙蛋（稀有）", "man WL", "man AL"}

// 碎片id，中奖超限，奖励这个
var CouponId = 14

// 奖品配置
type PrizeItem struct {
	Type            int    //类型 1-token 2-NFT
	TokenId         int    //奖品token id
	Id              int    //奖品id
	Name            string //名称
	RewardAmount    int    //单次奖励数量
	Ratio           int64  //概率，放大10000倍，50%=5000
	UserWinCount    int    //个人赢取上限 -1为不限制
	PrizeTotalCount int    //奖品总上限
	Link            string //外链
	ContractAddress string //合约地址
	Decimal         int    //精度
	Icon            string //图标
	Sort            int    //排序，数字越小越靠前
}

const (
	ImxIcon    = "https://game.mixmarvel.finance/icon/imx.png"
	RpgIcon    = "https://game.mixmarvel.finance/icon/rpg.png"
	UsdtIcon   = "https://game.mixmarvel.finance/icon/usdt.png"
	CouponIcon = "https://game.mixmarvel.finance/icon/coupon.png"
	ManAlIcon  = "https://game.mixmarvel.finance/icon/manal.png"
	ManWlIcon  = "https://game.mixmarvel.finance/icon/manwl.png"
	D3Pt       = "https://game.mixmarvel.finance/icon/3dpt.png"
	D3Xy       = "https://game.mixmarvel.finance/icon/3dts.png"
)

// 奖品配置
var PrizeConfig []PrizeItem

func InitPrizeConfig() {
	PrizeConfig = ImitPrizeConfig()
}

func ImitPrizeConfig() []PrizeItem {
	var prizeConfig []PrizeItem
	prizeConfig = append(prizeConfig,
		//IMX配置
		AddPrizeConfig(1, 1, 1, "MIX", 50000, 50, 1, 3, "", Conf.ContractC.Imx, 18, ImxIcon, 1),
		AddPrizeConfig(1, 1, 2, "MIX", 10000, 100, 1, 5, "", Conf.ContractC.Imx, 18, ImxIcon, 1),
		AddPrizeConfig(1, 1, 3, "MIX", 5000, 200, 1, 10, "", Conf.ContractC.Imx, 18, ImxIcon, 1),
		AddPrizeConfig(1, 1, 4, "MIX", 1000, 500, 3, 50, "", Conf.ContractC.Imx, 18, ImxIcon, 1),
		AddPrizeConfig(1, 1, 5, "MIX", 200, 1000, 10, 500, "", Conf.ContractC.Imx, 18, ImxIcon, 1),

		//RPG配置
		AddPrizeConfig(1, 2, 6, "RPG", 100, 100, 1, 5, "", Conf.ContractC.Rpg, 18, RpgIcon, 2),
		AddPrizeConfig(1, 2, 7, "RPG", 50, 50, 3, 10, "", Conf.ContractC.Rpg, 18, RpgIcon, 2),
		AddPrizeConfig(1, 2, 8, "RPG", 10, 200, 20, 50, "", Conf.ContractC.Rpg, 18, RpgIcon, 2),
		AddPrizeConfig(1, 2, 9, "RPG", 1, 800, 50, 500, "", Conf.ContractC.Rpg, 18, RpgIcon, 2),

		//USDT配置
		AddPrizeConfig(1, 3, 10, "USDT", 100, 50, 1, 5, "", Conf.ContractC.Usdt, 18, UsdtIcon, 3),
		AddPrizeConfig(1, 3, 11, "USDT", 10, 300, 10, 100, "", Conf.ContractC.Usdt, 18, UsdtIcon, 3),
		AddPrizeConfig(1, 3, 12, "USDT", 1, 500, 20, 500, "", Conf.ContractC.Usdt, 18, UsdtIcon, 3),

		//Coupon
		AddPrizeConfig(1, 4, 13, "Coupon", 5, 1500, 50, 1000, "", Conf.ContractC.Coupon, 18, CouponIcon, 4),
		AddPrizeConfig(1, 4, 14, "Coupon", 1, 2500, -1, -1, "", Conf.ContractC.Coupon, 18, CouponIcon, 4),

		//MAN AL配置
		AddPrizeConfig(2, 5, 15, "MAN AL", 1, 900, 1, 200, "https://man.metacene.io/", "", 0, ManAlIcon, 8),
		//MAN WL配置
		AddPrizeConfig(2, 6, 16, "MAN WL", 1, 200, 1, 20, "https://man.metacene.io/", "", 0, ManWlIcon, 7),
		//云斗龙3D 普通配置
		AddPrizeConfig(2, 7, 17, "云斗龙3D 普通", 1, 900, 1, 200, "", "", 0, D3Pt, 5),
		//云斗龙3D 稀有配置
		AddPrizeConfig(2, 8, 18, "云斗龙3D 稀有", 1, 200, 1, 20, "", "", 0, D3Xy, 6),
	)
	return prizeConfig
}

func AddPrizeConfig(ty int, tokenId int, id int, name string, rewardAmount int, ratio int64, userWinCount int, prizeTotalCount int, link string, contractAddress string, decimal int, icon string, sort int) PrizeItem {
	return PrizeItem{
		Type:            ty,
		TokenId:         tokenId,
		Id:              id,
		Name:            name,
		RewardAmount:    rewardAmount,
		Ratio:           ratio,
		UserWinCount:    userWinCount,
		PrizeTotalCount: prizeTotalCount,
		Link:            link,
		ContractAddress: contractAddress,
		Decimal:         decimal,
		Icon:            icon,
		Sort:            sort,
	}
}

// 获取某一个奖品的配置
func GetPrizeConfigById(id int) (PrizeItem, error) {
	for _, v := range PrizeConfig {
		if v.Id == id {
			return v, nil
		}
	}
	return PrizeItem{}, errors.New("奖品不存在")
}
