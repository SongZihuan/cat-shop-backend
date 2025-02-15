package v1

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/global/class/getclasslst"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/global/system/getconfig"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/global/system/getxieyi"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/global/user/login"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/global/user/register"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/global/wupin/gethotwupin"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/global/wupin/getsearch"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/global/wupin/getwupin"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbag/adminaddbag"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbag/admingetbag"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/admindaohuo"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/adminfahuo/adminacceptfahuoquxiao"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/adminfahuo/adminfahuochexiao"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/adminfahuo/adminfahuodengji"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/adminfahuo/changeinfo/adminfahuochangeshop"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/adminfahuo/changeinfo/adminfahuochuangeuser"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/admingetter/admingetbuyrecordbypage"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/admingetter/admingetbuyrecordinfo"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/admingetter/admingetuserbuyrecordlstbypage"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/adminpay/adminpeoplepay"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/adminpay/adminquxiaopay"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/admintuihuo/admintuihuoaccept"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/admintuihuo/admintuihuodaohuo"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/admintuihuo/admintuihuodengji"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminbuy/admintuihuo/admintuihuoshenqing"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminclass/adminaddclass"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminclass/admingetclass"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminclass/admingetclasslst"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminclass/adminupdateclass"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminfileupload/adminimageupload"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminfileupload/adminvideoupload"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/admingetconfig"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/admingetmsg"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/admingetxieyi"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminuser/adminadduser"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminuser/admingetuserinfo"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminuser/admingetuserlst"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminuser/admingetusermsg"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminuser/adminupdateuseravtar"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminuser/adminupdateuserinfo"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminuser/adminupdateuserpassword"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminuser/adminupdateuserphone"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminwupin/adminaddwupin"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminwupin/admingetwupin"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminwupin/admingetwupinlst"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/admin/adminwupin/adminupdatewupin"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/rootadmin/admindeleteconfig"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/rootadmin/adminrestartserver"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/rootadmin/adminstopserver"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/rootadmin/adminupdateconfigpic"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/rootadmin/adminupdateconfigstring"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/rootadmin/adminxieyi/adminupdatexieyi"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/bag/addbag"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/bag/getbaglst"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/daohuo"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/fahuo/fahuochangeuser"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/fahuo/fahuoquxiaoshenqing"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/getter/getbuyrecord"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/getter/getbuyrecordlst"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/getter/getbuyrecordlstbypage"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/pingjia"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/quxiaopay"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/tuihuo/tuihuodengji"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/buy/tuihuo/tuihuoshenqing"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/edit/updateuseravtar"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/edit/updateuserinfo"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/edit/updateuserpassword"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/getter/getuserinfo"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/kefu/sendmsg"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/pay/bagpay"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/pay/newpay"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/secret/user/pay/repay"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/api/test/secret/user/pay/testpay"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/ping"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/resource/image"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/resource/video"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/middleware"
)

func Ping(api *ginplus.Router) {
	api.GET("/hello", ping.Handler)
}

func Api(api *ginplus.Router) {
	api.Use(middleware.MustFormData(), middleware.XTokenMiddleware(), middleware.MustAccept(), middleware.ReturnContentJson())

	globalApiV1(api)
	secretApiV1(api)
	testApiV1(api)
}

func Resource(api *ginplus.Router) {
	api.GET("/image", image.Handler) // ${/resource}/v1/image
	api.GET("/video", video.Handler) // ${/resource}/v1/video
}

func globalApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/global")

	loginAndRegisterApiV1(api)
	configApiV1(api)
	classApiV1(api)
	wupinApiV1(api)
	xieyiApiV1(api)
}

func loginAndRegisterApiV1(api *ginplus.Router) {
	api.GET("/login", login.Handler)
	api.GET("/register", register.Handler)
}

func configApiV1(api *ginplus.Router) {
	api.GET("/config", getconfig.Handler)
}

func classApiV1(api *ginplus.Router) {
	api.GET("/class", getclasslst.Handler)
}

func wupinApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/wupin")
	api.GET("/info", getwupin.Handler)
	api.GET("/list", getsearch.Handler)
	api.GET("/hot", gethotwupin.Handler)
	api.GET("/search", getsearch.Handler)
}

func xieyiApiV1(api *ginplus.Router) {
	api.GET("/xieyi", getxieyi.Handler)
}

func secretApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/")
	api.Use(middleware.MustXTokenMiddleware())

	userApiV1(api)
	adminApiV1(api)
	rootAdminApiV1(api)
}

func userApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/user")

	userEditApiV1(api)
	userBuyRecordApiV1(api)
	userBagApiV1(api)
	userKefuApiV1(api)

	api.GET("/info", getuserinfo.Handler)
}

func userKefuApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/kefu")
	api.POST("/sendmsg", sendmsg.Handler)
}

func userEditApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/edit")

	api.POST("/info", updateuserinfo.Handler)
	api.POST("/password", updateuserpassword.Handler)
	api.POST("/avatar", updateuseravtar.Handler)
}

func userBuyRecordApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/buyrecord")
	userFahuoApiV1(api)
	userDaohuoApiV1(api)
	userTuihuoApiV1(api)
	userPayApiV1(api)

	api.GET("/info", getbuyrecord.Handler)
	api.GET("/list", getbuyrecordlst.Handler)
	api.GET("/list/infinite", getbuyrecordlst.Handler)
	api.GET("/lst/page", getbuyrecordlstbypage.Handler)
}

func userFahuoApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/fahuo")

	api.POST("/change/user", fahuochangeuser.Handler)
	api.POST("/quxiao", fahuoquxiaoshenqing.Handler)
}

func userDaohuoApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/daohuo")
	api.POST("/queren", daohuo.Handler)
	api.POST("/pingjia", pingjia.Handler)
}

func userTuihuoApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/tuihuo")

	api.POST("/shenqing", tuihuoshenqing.Handler)
	api.POST("/dengji", tuihuodengji.Handler)
}

func userPayApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/zhifu")
	api.POST("/new", newpay.Handler)
	api.POST("/bag", bagpay.Handler)
	api.POST("/repay", repay.Handler)
	api.POST("/quxiao", quxiaopay.Handler)
}

func userBagApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/bag")

	api.POST("/add", addbag.Handler)
	api.GET("/list", getbaglst.Handler)
	api.GET("/list/infinite", getbaglst.Handler)
}

func adminApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/admin")
	api.Use(middleware.MustAdminXTokenMiddleware())

	adminFileUploadApiV1(api)
	adminUserApiV1(api)
	adminUserListApiV1(api)
	adminWupinApiV1(api)
	adminClassApiV1(api)
	adminXieyiApiV1(api)
	adminMsgApiV1(api)
	adminBuyRecordApiV1(api)
	adminConfigApiV1(api)
}

func adminConfigApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/config")
	api.POST("/info", admingetconfig.Handler)
}

func adminFileUploadApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/file/upload")

	api.POST("/image", adminimageupload.Handler)
	api.POST("/video", adminvideoupload.Handler)
}

func adminMsgApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/kefu/msg")

	api.GET("/list", admingetmsg.Handler)
}

func adminXieyiApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/xieyi")

	api.GET("/info", admingetxieyi.Handler)
}

func adminClassApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/class")
	adminUpdateClassApiV1(api)
	adminAddClassApiV1(api)

	api.GET("/info", admingetclass.Handler)
	api.GET("/list", admingetclasslst.Handler)
}

func adminUpdateClassApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/edit")
	api.POST("/all", adminupdateclass.Handler)
}

func adminAddClassApiV1(api *ginplus.Router) {
	api.POST("/add", adminaddclass.Handler)
}

func adminUserApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/user")
	api.Use(middleware.MustAdminUserMiddleware(), middleware.AdminHasUserPermission())

	adminUpdateUserApiV1(api)
	adminUserBagApiV1(api)
	adminUserBuyRecordApiV1(api)
	adminUserMsgApiV1(api)

	api.GET("/info", admingetuserinfo.Handler)
	api.POST("/add", adminadduser.Handler)
}

func adminUserListApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/user")
	api.GET("/list", admingetuserlst.Handler) // 不需要MustAdmin
}

func adminBuyRecordApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/buyrecord")
	api.GET("/list", admingetbuyrecordbypage.Handler)
}

func adminUserMsgApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/kefu/msg")
	api.GET("/list", admingetusermsg.Handler)
}

func adminUpdateUserApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/edit")

	api.POST("/info", adminupdateuserinfo.Handler)
	api.POST("/password", adminupdateuserpassword.Handler)
	api.POST("/avatar", adminupdateuseravtar.Handler)
	api.POST("/phone", adminupdateuserphone.Handler)
}

func adminWupinApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/wupin")

	api.POST("/add", adminaddwupin.Handler)
	api.POST("/edit", adminupdatewupin.Handler)

	api.GET("/info", admingetwupin.Handler)
	api.GET("/list", admingetwupinlst.Handler)
}

func adminUserBagApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/bag")

	api.POST("/add", adminaddbag.Handler)
	api.GET("/list", admingetbag.Handler)
}

func adminUserBuyRecordApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/buyrecord")
	adminFahuoApiV1(api)
	adminDaohuoApiV1(api)
	adminPayApiV1(api)
	adminTuihuoApiV1(api)

	api.GET("/info", admingetbuyrecordinfo.Handler)
	api.GET("/list", admingetuserbuyrecordlstbypage.Handler)
}

func adminTuihuoApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/tuihuo")

	api.POST("/check", admintuihuoaccept.Handler)
	api.POST("/daohuo", admintuihuodaohuo.Handler)
	api.POST("/dengji", admintuihuodengji.Handler)
	api.POST("/shenqing", admintuihuoshenqing.Handler)
}

func adminPayApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/zhifu")

	api.POST("/people", adminpeoplepay.Handler)
	api.POST("/quxiao", adminquxiaopay.Handler)
}

func adminFahuoApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/fahuo")
	adminFahuoChangeInfoApiV1(api)

	api.POST("/quxiao/check", adminacceptfahuoquxiao.Handler)
	api.POST("/quxiao", adminfahuochexiao.Handler)
	api.POST("/chexiao", adminfahuochexiao.Handler)
	api.POST("/dengji", adminfahuodengji.Handler)
}

func adminFahuoChangeInfoApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/change")

	api.POST("/user", adminfahuochuangeuser.Handler)
	api.POST("/shop", adminfahuochangeshop.Handler)
}

func adminDaohuoApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/daohuo")

	api.POST("/queren", admindaohuo.Handler)
}

func rootAdminApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/admin/root")
	api.Use(middleware.MustRootAdminXTokenMiddleware())

	rootAdminConfigApiV1(api)
	rootAdminHttpServer(api)
	rootAdminXieyiApiV1(api)
}

func rootAdminXieyiApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/xieyi")

	api.GET("/edit", adminupdatexieyi.Handler)
}

func rootAdminConfigApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/config")

	api.POST("/delete", admindeleteconfig.Handler)
	api.POST("/update/pic", adminupdateconfigpic.Handler)
	api.POST("/update/string", adminupdateconfigstring.Handler)
}

func rootAdminHttpServer(apiV1 *ginplus.Router) {
	api := apiV1.Group("/httpserver")

	api.POST("/stop", adminstopserver.Handler)
	api.POST("/restart", adminrestartserver.Handler)
}

func testApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/test")
	api.Use(middleware.TestApiMiddleware())

	testGlobalApi(api)
	testSecretApiV1(api)
}

func testGlobalApi(apiV1 *ginplus.Router) {
	_ = apiV1.Group("/global")
}

func testSecretApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/")
	api.Use(middleware.MustXTokenMiddleware())

	testUserApiV1(api)
}

func testUserApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/user")

	testUserPayApiV1(api)
}

func testUserPayApiV1(apiV1 *ginplus.Router) {
	api := apiV1.Group("/zhifu")

	api.POST("/try", testpay.Handler)
}
