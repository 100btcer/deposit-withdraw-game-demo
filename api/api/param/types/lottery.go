package types

import "github.com/ethereum/go-ethereum/common"

type LotteryResultReq struct {
	Hash string `form:"hash"` //交易hash
}

type LotteryResultItem struct {
	Id         int64   `json:"id"`          //奖品id
	TokenId    int     `json:"token_id"`    //奖品token id
	Name       string  `json:"name"`        //奖品名称
	Type       int     `json:"type"`        //奖品类型 1-token 2-NFT
	Amount     float64 `json:"amount"`      //奖励数额
	Link       string  `json:"link"`        //外链
	Icon       string  `json:"icon"`        //icon
	CreateTime string  `json:"create_time"` //创建时间
}

type LotteryResultListReq struct {
	Page     int64 `form:"page"`      //页码
	PageSize int64 `form:"page_size"` //每页数量
	Type     *int  `form:"type"`      //类型 1-token 2-NFT，不传则获取全部记录
}

type LotteryResultListResp struct {
	Total int64               `json:"total"` //总数
	List  []LotteryResultItem `json:"list"`  //列表
}

type PrizeProofReq struct {
	Id int `form:"id"` //奖品id
}

type PrizeProofResp struct {
	Id            int           `json:"id"`
	TokenContract string        `json:"token_contract"`
	Amount        string        `json:"amount"`
	Proof         []common.Hash `json:"proof"`
}
