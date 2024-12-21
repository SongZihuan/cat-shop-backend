package modeltype

type BuyStatus int

const (
	WaitPayCheck      BuyStatus = 1
	PayCheckFail      BuyStatus = 2
	WaitFahuo         BuyStatus = 3
	WaitShouHuo       BuyStatus = 4
	WaitPingJia       BuyStatus = 5
	YiPingJia         BuyStatus = 6
	TuiHuoCheck       BuyStatus = 7
	WaitTuiHuoFahuo   BuyStatus = 8
	WaitTuiHuoShouHuo BuyStatus = 9
	TuiHuoFail        BuyStatus = 10
	TuiHuo            BuyStatus = 11
	CheckQuXiao       BuyStatus = 12
	QuXiao            BuyStatus = 13
)

var StatusToName = map[BuyStatus]string{
	WaitPayCheck:      "已下单，支付代确认",
	PayCheckFail:      "已下单，支付失败",
	WaitFahuo:         "待发货",
	WaitShouHuo:       "已发货",
	WaitPingJia:       "已收货",
	YiPingJia:         "已评价",
	TuiHuoCheck:       "申请退货审核中",
	WaitTuiHuoFahuo:   "申请退货成功，等待寄回",
	WaitTuiHuoShouHuo: "申请退货成功，待收货",
	TuiHuoFail:        "申请退货失败",
	TuiHuo:            "已退货",
	CheckQuXiao:       "申请取消订单审核中",
	QuXiao:            "已取消订单",
}
