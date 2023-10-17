package constant

// 资产变更类型
const (
	AssetLogOptAdd  = iota + 1 // 增加
	AssetLogOptRedu            // 扣除
)

// 资产类型
const (
	AssetTypeTicket = iota + 1 // ticket
)

// 资产变更业务类型
const (
	AssetBizTypeRecharge  = iota + 1 // 充值
	AssetBizTypeAdConsume            // 广告消耗
)

// 质押状态
const (
	TradeStatus0 = iota //未交易
	TradeStatus1        //待确认
	TradeStatus2        //已确认
	TradeStatus3        //已失败
)

const (
	Status0 = iota
	Status1
)
