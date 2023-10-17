package types

type RegisterResp struct {
	Token  string `json:"token"`   //token
	UserId int64  `json:"user_id"` //用户user_id
}

type WalletLoginReq struct {
	Address string `json:"address" binding:"required,len=42"` //钱包地址
	Message string `json:"message" binding:"required"`        //消息
	Sign    string `json:"sign" binding:"required"`           //签名
}

type UserInfoResp struct {
	UserId        int64   `json:"user_id"`        //用户id
	Address       string  `json:"address"`        //钱包地址
	Ticket        float64 `json:"-"`              //ticket数量
	DepositAmount float64 `json:"deposit_amount"` //已质押总额
}
