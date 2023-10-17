package types

type DepositRecordAddReq struct {
	Hash   string  `json:"hash"`    //交易hash
	Amount float64 `json:"amount"`  //质押数量
	PoolId int     `json:"pool_id"` //质押池子id
}

type DepositRecordListReq struct {
	Page     int64 `form:"page"`      //页码
	PageSize int64 `form:"page_size"` //每页数量
}

type DepositRecordListResp struct {
	Total int64                    `json:"total"` //总数
	List  []*DepositRecordListItem `json:"list"`  //数据
}

type DepositRecordListItem struct {
	Id            int64   `json:"id"`
	UserId        int64   `json:"user_id"`
	CreateTime    string  `json:"create_time"`
	LockDay       int     `json:"lock_day"`
	DepositAmount float64 `json:"deposit_amount"`
	Hash          string  `json:"hash"`
}

type WithdrawRecordAddReq struct {
	Hash   string  `json:"hash"`   //交易hash
	Amount float64 `json:"amount"` //提款数量
}
