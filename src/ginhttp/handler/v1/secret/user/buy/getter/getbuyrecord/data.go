package getbuyrecord

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type Query struct {
	ID uint `form:"id"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type _wp struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Pic        string `json:"pic"`
	ClassID    uint   `json:"classId"`
	ClassOf    *Class `json:"classOf"`
	Tag        string `json:"tag,omitempty"`
	HotPrice   int64  `json:"hotPrice,omitempty"`
	RealPrice  int64  `json:"realPrice"`
	Info       string `json:"info"`
	Ren        string `json:"ren"`
	Phone      string `json:"phone"`
	Email      string `json:"email,omitempty"`
	Wechat     string `json:"wechat,omitempty"`
	Location   string `json:"location"`
	BuyTotal   int64  `json:"buytotal"`
	BuyDaohuo  int64  `json:"buydaohuo"`
	BuyGood    int64  `json:"buygood"`
	BuyPrice   int64  `json:"buyprice"`
	BuyPingjia int64  `json:"buypingjia"`
	BuyJian    int64  `json:"buyjian"`
	Hot        bool   `json:"hot"`
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
	UserID             uint                `json:"userId"`
	WupinID            uint                `json:"wupinId"`
	ClassID            uint                `json:"classId"`
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
	Down               bool                `json:"down"`
}

func NewData(record *model.BuyRecord) BuyRecord {
	if record.IsWupinSale() {
		return BuyRecord{
			ID:                 record.ID,
			UserID:             record.UserID,
			WupinID:            record.WupinID,
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
			Wupin: &Wupin{
				ID:      record.Wupin.ID,
				Name:    record.WupinName,
				Pic:     record.WupinPic,
				ClassID: record.WupinClassID,
				ClassOf: &Class{
					ID:   record.WupinClassID,
					Name: record.WupinClass.Name,
				},
				Tag:        utils.GetSQLNullString(record.WupinTag),
				HotPrice:   record.WupinHotPrice.ToInt64(),
				RealPrice:  record.WupinRealPrice.ToInt64(),
				Info:       record.WupinInfo,
				Ren:        record.WupinRen,
				Phone:      record.WupinPhone,
				Email:      utils.GetSQLNullString(record.WupinEmail),
				Wechat:     utils.GetSQLNullString(record.WupinWeChat),
				Location:   record.WupinLocation,
				BuyTotal:   record.WupinBuyTotal.ToInt64(),
				BuyDaohuo:  record.WupinBuyDaoHuo.ToInt64(),
				BuyGood:    record.WupinBuyGood.ToInt64(),
				BuyPrice:   record.WupinBuyPrice.ToInt64(),
				BuyPingjia: record.WupinBuyPingJia.ToInt64(),
				BuyJian:    record.WupinBuyJian.ToInt64(),
				Hot:        record.WupinHot,
			},
			NowWupin: &NowWupin{
				ID:      record.Wupin.ID,
				Name:    record.Wupin.Name,
				Pic:     record.Wupin.Pic,
				ClassID: record.ClassID,
				ClassOf: &Class{
					ID:   record.Class.ID,
					Name: record.Class.Name,
				},
				Tag:        utils.GetSQLNullString(record.Wupin.Tag),
				HotPrice:   record.Wupin.HotPrice.ToInt64(),
				RealPrice:  record.Wupin.RealPrice.ToInt64(),
				Info:       record.Wupin.Info,
				Ren:        record.Wupin.Ren,
				Phone:      record.Wupin.Phone,
				Email:      utils.GetSQLNullString(record.Wupin.Email),
				Wechat:     utils.GetSQLNullString(record.Wupin.WeChat),
				Location:   record.Wupin.Location,
				BuyTotal:   record.Wupin.BuyTotal.ToInt64(),
				BuyDaohuo:  record.Wupin.BuyDaoHuo.ToInt64(),
				BuyGood:    record.Wupin.BuyGood.ToInt64(),
				BuyPrice:   record.Wupin.BuyPrice.ToInt64(),
				BuyPingjia: record.Wupin.BuyPingjia.ToInt64(),
				BuyJian:    record.Wupin.BuyJian.ToInt64(),
				Hot:        record.Wupin.Hot,
			},
			Down: false,
		}
	} else {
		return BuyRecord{
			ID:                 record.ID,
			UserID:             record.UserID,
			WupinID:            record.WupinID,
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
			Wupin: &Wupin{
				ID:      record.Wupin.ID,
				Name:    record.WupinName,
				Pic:     record.WupinPic,
				ClassID: record.WupinClassID,
				ClassOf: &Class{
					ID:   record.WupinClassID,
					Name: record.WupinClass.Name,
				},
				Tag:        utils.GetSQLNullString(record.WupinTag),
				HotPrice:   record.WupinHotPrice.ToInt64(),
				RealPrice:  record.WupinRealPrice.ToInt64(),
				Info:       record.WupinInfo,
				Ren:        record.WupinRen,
				Phone:      record.WupinPhone,
				Email:      utils.GetSQLNullString(record.WupinEmail),
				Wechat:     utils.GetSQLNullString(record.WupinWeChat),
				Location:   record.WupinLocation,
				BuyTotal:   record.WupinBuyTotal.ToInt64(),
				BuyDaohuo:  record.WupinBuyDaoHuo.ToInt64(),
				BuyGood:    record.WupinBuyGood.ToInt64(),
				BuyPrice:   record.WupinBuyPrice.ToInt64(),
				BuyPingjia: record.WupinBuyPingJia.ToInt64(),
				BuyJian:    record.WupinBuyJian.ToInt64(),
				Hot:        record.WupinHot,
			},
			NowWupin: &NowWupin{
				ID:      record.Wupin.ID,
				Name:    fmt.Sprintf("%s（以下架）", record.Wupin.Name),
				Pic:     record.Wupin.Pic,
				ClassID: record.ClassID,
				ClassOf: &Class{
					ID:   record.Class.ID,
					Name: record.Class.Name,
				},
				Tag:        "",
				HotPrice:   0,
				RealPrice:  0,
				Info:       record.Wupin.Info,
				Ren:        "无",
				Phone:      "000-0000-0000",
				Email:      "",
				Wechat:     "",
				Location:   "无",
				BuyTotal:   record.Wupin.BuyTotal.ToInt64(),
				BuyDaohuo:  record.Wupin.BuyDaoHuo.ToInt64(),
				BuyGood:    record.Wupin.BuyGood.ToInt64(),
				BuyPrice:   record.Wupin.BuyPrice.ToInt64(),
				BuyPingjia: record.Wupin.BuyPingjia.ToInt64(),
				BuyJian:    record.Wupin.BuyJian.ToInt64(),
				Hot:        false,
			},
			Down: true,
		}
	}
}

func NewJsonData(record *model.BuyRecord) data.Data {
	return data.NewSuccessWithData(NewData(record))
}
