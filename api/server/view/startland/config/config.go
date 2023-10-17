package config

import (
	"context"
	"mm-ndj/api/param/types"
	"mm-ndj/config"
	"mm-ndj/server/dao"
)

// 获取配置
func GetConfig(ctx context.Context, svc *dao.ServiceCtx, userId int64) (*types.ConfigResp, error) {
	c := make([]config.PoolItem, 0)
	for _, v := range config.PoolConfig {
		c = append(c, v)
	}
	//for i := 0; i < 1000; i++ {
	//	err := task.BurnTicketSuccess(context.Background(), svc, "0x878ADa7AF22A35Afc5D7c4f0AF57cdE44139a91c", "0xa97f48f0d0415a164725411cba19f75041752bdff1c21b4b57440167d766d907", 1000)
	//	fmt.Println(err)
	//}
	//userPrizeCount, err := dao.GetUserPrizeCount(ctx, svc, userId, 14)
	//fmt.Println(userPrizeCount, err)
	return &types.ConfigResp{
		PoolConfig:         c,
		LpContract:         config.LpContract,
		GachaContract:      config.GachaContract,
		PlayGameCostTicket: config.PlayGameCostTicket,
	}, nil
}
