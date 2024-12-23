package modeltype

type PayUrlKey string

const (
	PayFromTypeKey    PayUrlKey = "type"
	PayBuyRecordIdKey PayUrlKey = "buyRecordId"
	PayRedirectKey    PayUrlKey = "redirect"
	PayTypeKey        PayUrlKey = "paytype"
)

type PayFromType int

const (
	AliPay    PayFromType = 1
	WeChatPay PayFromType = 2
)

var PayFromTypeToName = map[PayFromType]string{
	AliPay:    "alipay",
	WeChatPay: "wechat",
}

type PayType int

const (
	NewPay PayType = 1
	BagPay PayType = 2
	Repay  PayType = 3
)

var PayTypeToName = map[PayType]string{
	NewPay: "newpay",
	BagPay: "shoppingbagpay",
	Repay:  "repay",
}
