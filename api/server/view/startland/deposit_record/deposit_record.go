package deposit_record

import (
	"context"
	"mm-ndj/api/param/types"
	"mm-ndj/server/dao"
	"time"
)

// 添加质押记录
func DepositRecordAdd(ctx context.Context, svc *dao.ServiceCtx, userId int64, hash string, amount float64, poolId int) error {
	return dao.DepositRecordAdd(ctx, svc, userId, hash, amount, poolId)
}

// 获取质押记录
func DepositRecordList(ctx context.Context, svc *dao.ServiceCtx, userId int64, page, pageSize int64) (*types.DepositRecordListResp, error) {
	list, total, err := dao.DepositRecordList(ctx, svc, userId, page, pageSize)
	if err != nil {
		return nil, err
	}
	item := make([]*types.DepositRecordListItem, len(list))
	for k, v := range list {
		item[k] = &types.DepositRecordListItem{
			Id:            v.Id,
			UserId:        v.UserId,
			CreateTime:    time.Unix(v.CreateTime, 0).Format("2006-01-02"),
			LockDay:       v.LockDay,
			DepositAmount: v.DepositAmount,
			Hash:          v.Hash,
		}
	}
	return &types.DepositRecordListResp{
		Total: total,
		List:  item,
	}, nil
}
