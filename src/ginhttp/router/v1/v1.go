package v1

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/class/getclasslst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/system/getconfig"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/system/getxieyi"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/system/image"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/system/video"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/user/login"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/user/register"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/wupin/gethotwupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/wupin/getsearch"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/global/wupin/getwupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/bag/adminaddbag"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/bag/admingetbag"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/daohuo/admindaohuo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/fahuo/adminacceptfahuoquxiao"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/fahuo/adminfahuodengji"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/fahuo/adminfahuoquxiao"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/fahuo/changeinfo/adminfahuochangeshop"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/fahuo/changeinfo/adminfahuochuangeuser"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/getter/admingetbuyrecordbypage"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/getter/admingetbuyrecordinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/getter/admingetuserbuyrecordlstbypage"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/pay/adminpeoplepay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/pay/adminquxiaopay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/tuihuo/admintuihuoaccept"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/tuihuo/admintuihuodaohuo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/tuihuo/admintuihuodengji"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/buy/tuihuo/admintuihuoshenqing"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/class/adminaddclass"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/class/adminchangeclassname"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/class/adminchangeclassshow"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/class/admingetclass"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/class/admingetclasslst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/class/adminupdateclass"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/fileupload/adminimageupload"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/fileupload/adminvideoupload"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/msg/admingetmsg"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/user/adminadduser"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/user/admingetuserinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/user/admingetuserlst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/user/admingetusermsg"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/user/adminupdateuseravtar"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/user/adminupdateuserinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/user/adminupdateuserpassword"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/user/adminupdateuserphone"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/wupin/adminaddwupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/wupin/admingetwupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/wupin/admingetwupinlst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/wupin/adminupdatewupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/xieyi/admingetxieyi"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/admin/xieyi/adminupdatexieyi"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/rootadmin/admindeleteconfig"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/rootadmin/admingetconfig"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/rootadmin/adminrestartserver"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/rootadmin/adminstopserver"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/rootadmin/adminupdateconfigpic"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/rootadmin/adminupdateconfigstring"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/bag/addbag"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/bag/getbaglst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/fahuo/daohuo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/fahuo/fahuochangeuser"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/fahuo/fahuoquxiaoshenqing"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/fahuo/pingjia"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/fahuo/quxiaopay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/getter/getbuyrecord"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/getter/getbuyrecordlst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/getter/getbuyrecordlstbypage"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/tuihuo/tuihuodengji"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/buy/tuihuo/tuihuoshenqing"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/edit/updateuseravtar"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/edit/updateuserinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/edit/updateuserpassword"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/getter/getuserinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/kefu/sendmsg"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/pay/bagpay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/pay/newpay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/secret/user/pay/repay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/v1/test/secret/user/pay/testpay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/middleware"
	"github.com/gin-gonic/gin"
)

func Api(api *gin.RouterGroup) {
	api.Use(middleware.DBReady(), middleware.MustFormData(), middleware.XTokenMiddleware(), middleware.AdminUser(), middleware.MustAccept(), middleware.ReturnContentJson())

	globalApiV1(api)
	secretApiV1(api)
	testApiV1(api)
}

func Resource(api *gin.RouterGroup) {
	api.GET("/image", image.Handler) // ${/resource}/v1/image
	api.GET("/video", video.Handler) // ${/resource}/v1/video
}

func globalApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/global")

	loginAndRegisterApiV1(api)
	configApiV1(api)
	classApiV1(api)
	wupinApiV1(api)
	xieyiApiV1(api)
}

func loginAndRegisterApiV1(api *gin.RouterGroup) {
	api.GET("/login", login.Handler)
	api.GET("/register", register.Handler)
}

func configApiV1(api *gin.RouterGroup) {
	api.GET("/config", getconfig.Handler)
}

func classApiV1(api *gin.RouterGroup) {
	api.GET("/class", getclasslst.Handler)
}

func wupinApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/wupin")
	api.GET("/info", getwupin.Handler)
	api.GET("/list", getsearch.Handler)
	api.GET("/hot", gethotwupin.Handler)
	api.GET("/search", getsearch.Handler)
}

func xieyiApiV1(api *gin.RouterGroup) {
	api.GET("/xieyi", getxieyi.Handler)
}

func secretApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/")
	api.Use(middleware.MustXTokenMiddleware())

	userApiV1(api)
	adminApiV1(api)
	rootAdminApiV1(api)
}

func userApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/user")

	userEditApiV1(api)
	userBuyRecordApiV1(api)
	userBagApiV1(api)
	userKefuApiV1(api)

	api.GET("/info", getuserinfo.Handler)
}

func userKefuApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/kefu")
	api.POST("/sendmsg", sendmsg.Handler)
}

func userEditApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/edit")

	api.POST("/info", updateuserinfo.Handler)
	api.POST("/password", updateuserpassword.Handler)
	api.POST("/avatar", updateuseravtar.Handler)
}

func userBuyRecordApiV1(apiV1 *gin.RouterGroup) {
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

func userFahuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/fahuo")

	api.POST("/change/user", fahuochangeuser.Handler)
	api.POST("/quxiao", fahuoquxiaoshenqing.Handler)
}

func userDaohuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/daohuo")
	api.POST("/queren", daohuo.Handler)
	api.POST("/pingjia", pingjia.Handler)
}

func userTuihuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/tuihuo")

	api.POST("/shenqing", tuihuoshenqing.Handler)
	api.POST("/dengji", tuihuodengji.Handler)
}

func userPayApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/zhifu")
	api.POST("/new", newpay.Handler)
	api.POST("/bag", bagpay.Handler)
	api.POST("/repay", repay.Handler)
	api.POST("/quxiao", quxiaopay.Handler)
}

func userBagApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/bag")

	api.POST("/add", addbag.Handler)
	api.GET("/list", getbaglst.Handler)
}

func adminApiV1(apiV1 *gin.RouterGroup) {
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
}

func adminFileUploadApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/file/upload")

	api.POST("/image", adminimageupload.Handler)
	api.POST("/video", adminvideoupload.Handler)
}

func adminMsgApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/kefu/msg")

	api.GET("/list", admingetmsg.Handler)
}

func adminXieyiApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/xieyi")

	api.GET("/info", admingetxieyi.Handler)
	api.GET("/edit", adminupdatexieyi.Handler)
}

func adminClassApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/class")
	adminUpdateClassApiV1(api)
	adminAddClassApiV1(api)

	api.GET("/info", admingetclass.Handler)
	api.GET("/list", admingetclasslst.Handler)
}

func adminUpdateClassApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/edit")
	api.POST("/name", adminchangeclassname.Handler)
	api.POST("/show", adminchangeclassshow.Handler)
	api.POST("/all", adminupdateclass.Handler)
}

func adminAddClassApiV1(api *gin.RouterGroup) {
	api.POST("/add", adminaddclass.Handler)
}

func adminUserApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/user")
	api.Use(middleware.MustAdminUserMiddleware())

	adminUpdateUserApiV1(api)
	adminUserBagApiV1(api)
	adminUserBuyRecordApiV1(api)
	adminUserMsgApiV1(api)

	api.GET("/info", admingetuserinfo.Handler)
	api.POST("/add", adminadduser.Handler)
}

func adminUserListApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/user")
	api.GET("/list", admingetuserlst.Handler) // 不需要MustAdmin
}

func adminBuyRecordApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/buyrecord")
	api.GET("/list", admingetbuyrecordbypage.Handler)
}

func adminUserMsgApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/kefu/msg")
	api.GET("/list", admingetusermsg.Handler)
}

func adminUpdateUserApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/edit")

	api.POST("/info", adminupdateuserinfo.Handler)
	api.POST("/password", adminupdateuserpassword.Handler)
	api.POST("/avatar", adminupdateuseravtar.Handler)
	api.POST("/phone", adminupdateuserphone.Handler)
}

func adminWupinApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/wupin")

	api.POST("/add", adminaddwupin.Handler)
	api.POST("/edit", adminupdatewupin.Handler)

	api.GET("/info", admingetwupin.Handler)
	api.GET("/list", admingetwupinlst.Handler)
}

func adminUserBagApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/bag")

	api.POST("/add", adminaddbag.Handler)
	api.GET("/list", admingetbag.Handler)
}

func adminUserBuyRecordApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/buyrecord")
	adminFahuoApiV1(api)
	adminDaohuoApiV1(api)
	adminPayApiV1(api)
	adminTuihuoApiV1(api)

	api.GET("/info", admingetbuyrecordinfo.Handler)
	api.GET("/list", admingetuserbuyrecordlstbypage.Handler)
}

func adminTuihuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/tuihuo")

	api.POST("/tongyi", admintuihuoaccept.Handler)
	api.POST("/daohuo", admintuihuodaohuo.Handler)
	api.POST("/dengji", admintuihuodengji.Handler)
	api.POST("/shenqing", admintuihuoshenqing.Handler)
}

func adminPayApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/zhifu")

	api.POST("/people", adminpeoplepay.Handler)
	api.POST("/quxiao", adminquxiaopay.Handler)
}

func adminFahuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/fahuo")
	adminFahuoChangeInfoApiV1(api)

	api.POST("/quxiao/tongyi", adminacceptfahuoquxiao.Handler)
	api.POST("/quxiao", adminfahuoquxiao.Handler)
	api.POST("/dengji", adminfahuodengji.Handler)
}

func adminFahuoChangeInfoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/change")

	api.POST("/user", adminfahuochuangeuser.Handler)
	api.POST("/shop", adminfahuochangeshop.Handler)
}

func adminDaohuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/daohuo")

	api.POST("/queren", admindaohuo.Handler)
}

func rootAdminApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/admin/root")
	api.Use(middleware.MustRootAdminXTokenMiddleware())

	rootAdminConfigApiV1(api)
	rootAdminHttpServer(api)
}

func rootAdminConfigApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/config")

	api.POST("/info", admingetconfig.Handler)
	api.POST("/delete", admindeleteconfig.Handler)
	api.POST("/update/pic", adminupdateconfigpic.Handler)
	api.POST("/update/string", adminupdateconfigstring.Handler)
}

func rootAdminHttpServer(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/httpserver")

	api.POST("/stop", adminstopserver.Handler)
	api.POST("/restart", adminrestartserver.Handler)
}

func testApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/test")
	api.Use(middleware.TestApiMiddleware())

	testGlobalApi(api)
	testSecretApiV1(api)
}

func testGlobalApi(apiV1 *gin.RouterGroup) {
	_ = apiV1.Group("/global")
}

func testSecretApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/")
	api.Use(middleware.MustXTokenMiddleware())

	testUserApiV1(api)
}

func testUserApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/user")

	testUserPayApiV1(api)
}

func testUserPayApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/zhifu")

	api.POST("/try", testpay.Handler)
}
