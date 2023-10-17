package types

import "mm-ndj/config"

type ConfigResp struct {
	PoolConfig         []config.PoolItem `json:"pool_config"`           //奖池配置
	LpContract         string            `json:"lp_contract"`           //lp代币合约地址
	PlayGameCostTicket int               `json:"play_game_cost_ticket"` //单次游戏花费ticket数量
	GachaContract      string            `json:"gacha_contract"`        //扭蛋合约
}
