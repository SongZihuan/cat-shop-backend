package getbuyrecordlstbypage

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
)

type Query struct {
	Page     int `json:"page"`
	PageSize int `json:"pagesize"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type _wp struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	ClassID   uint   `json:"classid"`
	ClassOf   *Class `json:"classOf"`
	Tag       string `json:"tag,omitempty"`
	HotPrice  int64  `json:"hotPrice,omitempty"`
	RealPrice int64  `json:"realPrice"`
	Info      string `json:"info"`
	Ren       string `json:"ren"`
	Phone     string `json:"phone"`
	Email     string `json:"email,omitempty"`
	Wechat    string `json:"wechat,omitempty"`
	Location  string `json:"location"`
	BuyTotal  int64  `json:"buytotal"`
	BuyDaohuo int64  `json:"buydaohuo"`
	BuyGood   int64  `json:"buygood"`
}

type NowWupin _wp
type Wupin _wp

type LocationForUser struct {
	Name     string `form:"name"`
	Phone    string `form:"phone"`
	Location string `form:"location"`
	Wechat   string `form:"wechat"`
	Email    string `form:"email"`
	Remark   string `form:"remark"`
}

type BuyRecord struct {
	ID                 uint                `json:"id"`
	UserID             uint                `json:"user_id"`
	WupinID            uint                `json:"wupin_id"`
	ClassID            uint                `json:"class_id"`
	Num                modeltype.Total     `json:"num"`
	Price              modeltype.Price     `json:"price"`
	TotalPrice         modeltype.Price     `json:"total_price"`
	Time               int64               `json:"time"`
	FukuanTime         int64               `json:"fukuantime"`
	FahuoTime          int64               `json:"fahuotime"`
	ShouhuoTime        int64               `json:"shouhuotime"`
	PingjiaTime        int64               `json:"pingjiatime"`
	TuiHuoShenQingTime int64               `json:"tuihuoshenqingtime"`
	DengjituihuoTime   int64               `json:"dengjituihuotime"`
	QuerentuihuoTime   int64               `json:"querentuihuotime"`
	TuohuoTime         int64               `json:"tuohuotime"`
	Quxiaotime         int64               `json:"quxiaotime"`
	Status             modeltype.BuyStatus `json:"status"`
	Kuaidi             string              `json:"kuaidi"`
	KuaidiNum          string              `json:"kuaidinum"`
	BackKuaidi         string              `json:"backkuaidi"`
	BackKuaidiNum      string              `json:"backkuaidinum"`
	IsGood             bool                `json:"isgood"`
	User               LocationForUser     `json:"user"`
	Shop               LocationForUser     `json:"shop"`
	Wupin              *Wupin              `json:"wupin"`
	NowWupin           *NowWupin           `json:"nowwupin"`
}

func NewBuyRecord(record *model.BuyRecord) BuyRecord {
	var class *Class
	var nwp *NowWupin
	var wp *Wupin

	if record.WuPin.ClassID >= 0 && record.WuPin.Class != nil && record.WuPin.Class.Show {
		class = &Class{
			ID:   record.WuPin.ClassID,
			Name: record.WuPin.Class.Name,
		}
	} else {
		class = &Class{
			ID:   modeltype.ClassEmptyID,
			Name: modeltype.ClassEmptyName,
		}
	}

	wp = &Wupin{
		ID:        record.WuPin.ID,
		Name:      record.WupinName,
		Pic:       record.WupinPic,
		ClassID:   class.ID,
		ClassOf:   class,
		Tag:       utils.GetSQLNullString(record.WupinTag),
		HotPrice:  modeltype.GetPrice(record.WupinHotPrice),
		RealPrice: modeltype.GetPrice(record.WupinRealPrice),
		Info:      record.WupinInfo,
		Ren:       record.WupinRen,
		Phone:     record.WupinPhone,
		Email:     utils.GetSQLNullString(record.WupinEmail),
		Wechat:    utils.GetSQLNullString(record.WupinWeChat),
		Location:  record.WupinLocation,
		BuyTotal:  modeltype.GetTotal(record.WupinBuyTotal),
		BuyDaohuo: modeltype.GetTotal(record.WupinBuyDaoHuo),
		BuyGood:   modeltype.GetTotal(record.WupinBuyGood),
	}

	if record.WuPinID > 0 && record.WuPin != nil {
		nwp = &NowWupin{
			ID:        record.WuPin.ID,
			Name:      record.WuPin.Name,
			Pic:       record.WuPin.Pic,
			ClassID:   class.ID,
			ClassOf:   class,
			Tag:       utils.GetSQLNullString(record.WuPin.Tag),
			HotPrice:  modeltype.GetPrice(record.WuPin.HotPrice),
			RealPrice: modeltype.GetPrice(record.WuPin.RealPrice),
			Info:      record.WuPin.Info,
			Ren:       record.WuPin.Ren,
			Phone:     record.WuPin.Phone,
			Email:     utils.GetSQLNullString(record.WuPin.Email),
			Wechat:    utils.GetSQLNullString(record.WuPin.WeChat),
			Location:  record.WuPin.Location,
			BuyTotal:  modeltype.GetTotal(record.WuPin.BuyTotal),
			BuyDaohuo: modeltype.GetTotal(record.WuPin.BuyDaoHuo),
			BuyGood:   modeltype.GetTotal(record.WuPin.BuyGood),
		}
	} else {
		nwp = &NowWupin{
			ID: modeltype.WupinEmptyID,
		}
	}

	return BuyRecord{
		ID:                 record.ID,
		UserID:             record.UserID,
		WupinID:            record.WuPinID,
		ClassID:            record.ClassID,
		Num:                record.Num,
		Price:              record.Price,
		TotalPrice:         record.TotalPrice,
		Time:               record.XiaDanTime.Unix(),
		FukuanTime:         utils.GetSQLNullTimeUnix(record.FuKuanTime),
		FahuoTime:          utils.GetSQLNullTimeUnix(record.FaHuoTime),
		ShouhuoTime:        utils.GetSQLNullTimeUnix(record.ShouHuoTime),
		PingjiaTime:        utils.GetSQLNullTimeUnix(record.PingJiaTime),
		TuiHuoShenQingTime: utils.GetSQLNullTimeUnix(record.TuiHuoShenQingTime),
		DengjituihuoTime:   utils.GetSQLNullTimeUnix(record.DengJiTuiHuoTime),
		QuerentuihuoTime:   utils.GetSQLNullTimeUnix(record.QueRenTuiHuoTime),
		TuohuoTime:         utils.GetSQLNullTimeUnix(record.TuiHuoTime),
		Quxiaotime:         utils.GetSQLNullTimeUnix(record.QuXiaoTime),
		Status:             record.Status,
		Kuaidi:             utils.GetSQLNullString(record.FaHuoKuaiDi),
		KuaidiNum:          utils.GetSQLNullString(record.FaHuoKuaiDiNum),
		BackKuaidi:         utils.GetSQLNullString(record.TuiHuoKuaiDi),
		BackKuaidiNum:      utils.GetSQLNullString(record.TuiHuoKuaiDiNum),
		IsGood:             record.IsGood.Bool,
		User: LocationForUser{
			Name:     record.UserName,
			Phone:    record.UserPhone,
			Location: record.UserLocation,
			Wechat:   utils.GetSQLNullString(record.UserWeChat),
			Email:    utils.GetSQLNullString(record.UserEmail),
			Remark:   utils.GetSQLNullString(record.UserRemark),
		},
		Shop: LocationForUser{
			Name:     record.ShopName,
			Phone:    record.ShopPhone,
			Location: record.ShopLocation,
			Wechat:   utils.GetSQLNullString(record.ShopWeChat),
			Email:    utils.GetSQLNullString(record.ShopEmail),
			Remark:   utils.GetSQLNullString(record.ShopRemark),
		},
		Wupin:    wp,
		NowWupin: nwp,
	}
}

type Data struct {
	List     []BuyRecord `json:"list"`
	Total    int         `json:"total"`
	MaxCount int         `json:"maxpage"`
}

func NewData(res []model.BuyRecord, maxcount int) Data {
	list := make([]BuyRecord, len(res))
	for _, v := range res {
		list = append(list, NewBuyRecord(&v))
	}

	return Data{
		List:     list,
		Total:    len(res),
		MaxCount: maxcount,
	}
}

func NewJsonData(res []model.BuyRecord, maxcount int) data.Data {
	return data.NewSuccessWithData(NewData(res, maxcount))
}
