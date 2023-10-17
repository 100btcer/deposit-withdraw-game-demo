package task

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mm-ndj/config"
	"mm-ndj/constant"
	"mm-ndj/model"
	"mm-ndj/pkg/lock"
	"mm-ndj/pkg/lottery"
	"mm-ndj/server/dao"
	"time"
)

// 监控到数据之后，记录数据并开奖
func BurnTicketSuccess(ctx context.Context, svc *dao.ServiceCtx, address string, hash string, amount float64) error {
	lockKey := fmt.Sprintf("lottery:lock")
	redisLock := lock.NewRedisLock(svc.Rds.GetRedisClient(), 5)
	ok, err := redisLock.Lock(lockKey)
	if err != nil {
		config.Logger.Error("BurnTicketSuccess", zap.Error(err))
		return err
	}
	defer redisLock.UnLock(lockKey)
	if !ok {
		return errors.New("访问频繁，请稍后再试")
	}

	if amount < config.PlayGameCostTicket {
		config.Logger.Error("BurnTicketSuccess amount不足", zap.Error(err))
		return errors.New("amount不足")
	}
	times := amount / config.PlayGameCostTicket
	for i := 0; i < int(times); i++ {
		err := BurnTicketSuccess2(ctx, svc, address, hash, config.PlayGameCostTicket)
		if err != nil {
			config.Logger.Error("BurnTicketSuccess", zap.Error(err))
			return err
		}
	}
	return nil
}
func BurnTicketSuccess2(ctx context.Context, svc *dao.ServiceCtx, address string, hash string, amount float64) error {
	userInfo, err := dao.GetUserByAddress(ctx, svc, address)
	if err != nil {
		config.Logger.Error("BurnTicketSuccess2", zap.Error(err))
		return err
	}
	userId := userInfo.Id
	//获取开奖结果
	prize, err := getPlayLotteryResult(ctx, svc, userId)
	if err != nil {
		config.Logger.Error("BurnTicketSuccess2", zap.Error(err))
		return err
	}
	//记录数据
	//记录奖品
	prizeM := model.NewPrizeRecordModel(ctx, svc.Db)
	err = prizeM.Insert(&model.PrizeRecord{
		UserId:     userId,
		PrizeId:    prize.Id,
		Amount:     float64(prize.RewardAmount),
		BurnHash:   hash,
		PrizeType:  prize.Type,
		Status:     constant.Status0,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	})
	if err != nil {
		config.Logger.Error("BurnTicketSuccess2", zap.Error(err))
		return err
	}
	//更新奖品总数
	err = updateAwardAmount(ctx, svc, userId, prize.Id, prize.Type, prize.TokenId, float64(prize.RewardAmount), prize.Sort)
	if err != nil {
		config.Logger.Error("BurnTicketSuccess2", zap.Error(err))
		return err
	}
	return nil
}

// 更新未领取奖品总量
func updateAwardAmount(ctx context.Context, svc *dao.ServiceCtx, userId int64, prizeId int, prizeType int, tokenId int, amount float64, sort int) error {
	prizeM := model.NewPrizeModel(ctx, svc.Db)
	//查询未领取奖励
	prizeRes, err := dao.GetPrizeByTokenId(ctx, svc, userId, tokenId, constant.Status0)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		//奖品不存在，新增
		err = prizeM.Insert(&model.Prize{
			UserId:      userId,
			PrizeId:     prizeId,
			TokenId:     tokenId,
			Amount:      amount,
			Status:      0,
			UpdateTime:  0,
			CreateTime:  time.Now().Unix(),
			PrizeType:   prizeType,
			ReceiveHash: "",
			Sort:        sort,
		})
		if err != nil {
			return err
		}
	} else {
		//更新
		prizeRes.Amount += amount
		prizeRes.UpdateTime = time.Now().Unix()
		err = prizeM.UpdateByCondition(prizeRes, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", prizeRes.Id)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取游戏结果
func getPlayLotteryResult(ctx context.Context, svc *dao.ServiceCtx, userId int64) (*config.PrizeItem, error) {
	rateMap := getPrizeRatio()
	l := lottery.NewLottery(rateMap)
	resPrizeId := l.RandomType(time.Now().UnixNano())
	//中奖奖品
	resPrize, err := getPrizeConfigById(resPrizeId)
	if err != nil {
		//中奖奖品配置不存在
		config.Logger.Error("getPlayLotteryResult", zap.Error(err))
		return nil, err
	}
	//查询用户该奖品中奖次数
	userPrizeCount, err := dao.GetUserPrizeCount(ctx, svc, userId, resPrizeId)
	if err != nil {
		config.Logger.Error("getPlayLotteryResult", zap.Error(err))
		return nil, err
	}
	//新的中奖奖品id
	newResPrizeId := resPrizeId
	//碎片奖品id
	prize4Id := config.CouponId
	if resPrize.UserWinCount != -1 && userPrizeCount >= int64(resPrize.UserWinCount) {
		//超限，奖励碎片
		//fmt.Println("超过个人限制", userPrizeCount, resPrize.UserWinCount)
		newResPrizeId = prize4Id
	}
	//查询奖励已发放总量
	prizeCount, err := dao.GetPrizeCount(ctx, svc, resPrizeId)
	if err != nil {
		config.Logger.Error("getPlayLotteryResult", zap.Error(err))
		return nil, err
	}
	if resPrize.PrizeTotalCount != -1 && prizeCount >= int64(resPrize.PrizeTotalCount) {
		//超限，奖励碎片
		//fmt.Println("超过总量限制", prizeCount, resPrize.PrizeTotalCount)
		newResPrizeId = prize4Id
	}
	//获取最终的奖品数据
	newResPrize, err := getPrizeConfigById(newResPrizeId)
	if err != nil {
		config.Logger.Error("getPlayLotteryResult", zap.Error(err))
		return nil, err
	}
	return &newResPrize, nil
}

// 获取奖品id=>ratio配置
func getPrizeRatio() map[int]int64 {
	ratio := make(map[int]int64, 0)
	prizeConfig := config.PrizeConfig
	for _, v := range prizeConfig {
		ratio[v.Id] = v.Ratio
	}
	return ratio
}

// 获取某一个奖品的配置
func getPrizeConfigById(id int) (config.PrizeItem, error) {
	for _, v := range config.PrizeConfig {
		if v.Id == id {
			return v, nil
		}
	}
	return config.PrizeItem{}, errors.New("奖品不存在")
}
