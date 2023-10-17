package lottery

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"mm-ndj/api/param/types"
	"mm-ndj/config"
	"mm-ndj/constant"
	"mm-ndj/model"
	"mm-ndj/pkg/errcode"
	"mm-ndj/server/dao"
	"mm-ndj/server/task"
	"time"
)

// 获取开奖结果
func GetLotteryResult(ctx context.Context, svc *dao.ServiceCtx, userId int64, hash string) (interface{}, error) {
	prizes, err := dao.GetPrizeByUserIdAndBurnHash(ctx, svc, userId, hash)
	if err != nil {
		return nil, err
	}
	var res []*types.LotteryResultItem
	if len(prizes) == 0 {
		return []struct{}{}, nil
	}
	for _, v := range prizes {
		prize, err := getPrizeConfigById(v.PrizeId)
		if err != nil {
			return nil, err
		}
		res = append(res, &types.LotteryResultItem{
			Id:         v.Id,
			TokenId:    prize.TokenId,
			Name:       prize.Name,
			Type:       prize.Type,
			Amount:     v.Amount,
			Link:       prize.Link,
			Icon:       prize.Icon,
			CreateTime: time.Unix(v.CreateTime, 0).Format("2006-01-02 15:04:05"),
		})
	}
	return res, nil
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

// 获取奖品记录
func GetLotteryPrizeList(ctx context.Context, svc *dao.ServiceCtx, userId int64, ty *int, page, pageSize int64) (interface{}, error) {
	prizes, total, err := dao.GetLotteryPrizeList(ctx, svc, userId, ty, page, pageSize)
	if err != nil {
		return struct{}{}, err
	}
	if len(prizes) == 0 {
		return struct{}{}, nil
	}
	var list []types.LotteryResultItem
	for _, v := range prizes {
		prize, err := getPrizeConfigById(v.PrizeId)
		if err != nil {
			return struct{}{}, err
		}
		list = append(list, types.LotteryResultItem{
			Id:         v.Id,
			TokenId:    prize.TokenId,
			Name:       prize.Name,
			Type:       v.PrizeType,
			Amount:     v.Amount,
			Link:       prize.Link,
			Icon:       prize.Icon,
			CreateTime: time.Unix(v.CreateTime, 0).Format("2006-01-02 15:04:03"),
		})
	}
	return &types.LotteryResultListResp{
		Total: total,
		List:  list,
	}, nil
}

// 获取奖品记录
func GetLotteryPrizeRecordList(ctx context.Context, svc *dao.ServiceCtx, userId int64, ty *int, page, pageSize int64) (interface{}, error) {
	prizes, total, err := dao.GetLotteryPrizeRecordList(ctx, svc, userId, ty, page, pageSize)
	if err != nil {
		return struct{}{}, err
	}
	if len(prizes) == 0 {
		return struct{}{}, nil
	}
	var list []types.LotteryResultItem
	for _, v := range prizes {
		prize, err := getPrizeConfigById(v.PrizeId)
		if err != nil {
			return struct{}{}, err
		}
		list = append(list, types.LotteryResultItem{
			Id:         v.Id,
			TokenId:    prize.TokenId,
			Name:       prize.Name,
			Type:       v.PrizeType,
			Amount:     v.Amount,
			Link:       prize.Link,
			Icon:       prize.Icon,
			CreateTime: time.Unix(v.CreateTime, 0).Format("2006-01-02 15:04:03"),
		})
	}
	return &types.LotteryResultListResp{
		Total: total,
		List:  list,
	}, nil
}

// 获取奖品证明
func GetPrizeProof(ctx context.Context, svc *dao.ServiceCtx, userId int64, id int) (*types.PrizeProofResp, error) {
	prizeRecord, err := dao.GetPrizeById(ctx, svc, id)
	if err != nil {
		return nil, err
	}
	prize, err := config.GetPrizeConfigById(prizeRecord.PrizeId)
	//更新默克尔树
	cliam, err := task.NewMerkleTask(svc).Run()
	if err != nil {
		return nil, err
	}
	var proof []common.Hash
	var amount string
	for _, v := range cliam {
		if v.Id == id {
			proof = v.Proof
			amount = v.Amount
			break
		}
	}
	return &types.PrizeProofResp{
		Id:            id,
		TokenContract: prize.ContractAddress,
		Amount:        amount,
		Proof:         proof,
	}, nil
}

// 发奖
func GiveOutAward(ctx context.Context, svc *dao.ServiceCtx, userId int64, id int) (string, error) {
	prizeRes, err := dao.GetPrizeById(ctx, svc, id)
	if err != nil {
		return "", err
	}
	if prizeRes.Status != constant.Status0 {
		return "", errcode.NewCustomErr("已经领取过")
	}
	if prizeRes.UserId != userId {
		return "", errcode.NewCustomErr("无权领取")
	}
	prize, err := config.GetPrizeConfigById(prizeRes.PrizeId)
	hash, err := task.SysTransfer(ctx, svc, prize.ContractAddress, prizeRes.Amount, prizeRes.Id)
	if err != nil {
		return "", err
	}
	m := model.NewPrizeModel(ctx, svc.Db)
	prizeRes.ReceiveHash = hash
	prizeRes.Status = constant.Status1 //已领取
	prizeRes.UpdateTime = time.Now().Unix()
	err = m.UpdateByCondition(prizeRes, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
	return hash, err
}
