package router

import (
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/class/getclasslst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/system/getconfig"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/system/getxieyi"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/system/image"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/system/video"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/user/login"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/user/register"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/wupin/gethotwupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/wupin/getsearch"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/global/wupin/getwupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/bag/adminaddbag"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/bag/admingetbag"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/daohuo/admindaohuo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/fahuo/adminacceptfahuoquxiao"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/fahuo/adminfahuodengji"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/fahuo/adminfahuoquxiao"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/fahuo/changeinfo/adminfahuochangeshop"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/fahuo/changeinfo/adminfahuochuangeuser"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/getter/admingetbuyrecordbypage"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/getter/admingetbuyrecordinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/getter/admingetuserbuyrecordlstbypage"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/pay/adminpeoplepay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/pay/adminquxiaopay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/tuihuo/admintuihuoaccept"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/tuihuo/admintuihuodaohuo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/tuihuo/admintuihuodengji"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/buy/tuihuo/admintuihuoshenqing"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/class/adminaddclass"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/class/adminchangeclassname"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/class/adminchangeclassshow"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/class/admingetclass"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/class/admingetclasslst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/class/adminupdateclass"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/fileupload/adminimageupload"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/fileupload/adminvideoupload"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/msg/admingetmsg"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/user/adminadduser"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/user/admingetuserinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/user/admingetuserlst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/user/admingetusermsg"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/user/adminupdateuseravtar"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/user/adminupdateuserinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/user/adminupdateuserpassword"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/user/adminupdateuserphone"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/wupin/admingetwupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/wupin/admingetwupinlst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/wupin/adminupdatewupin"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/xieyi/admingetxieyi"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/admin/xieyi/adminupdatexieyi"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/rootadmin/admindeleteconfig"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/rootadmin/admingetconfig"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/rootadmin/adminupdateconfigpic"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/rootadmin/adminupdateconfigstring"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/bag/addbag"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/bag/getbaglst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/buy/fahuo/daohuo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/buy/fahuo/fahuochangeuser"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/buy/fahuo/fahuoquxiaoshenqing"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/buy/getter/getbuyrecord"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/buy/getter/getbuyrecordlst"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/buy/getter/getbuyrecordlstbypage"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/buy/tuihuo/tuihuodengji"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/buy/tuihuo/tuihuoshenqing"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/edit/updateuseravtar"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/edit/updateuserinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/edit/updateuserpassword"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/getter/getuserinfo"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/kefu/sendmsg"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/pay/bagpay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/pay/newpay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/secret/user/pay/repay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/test/secret/user/pay/testpay"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	if !config.IsReady() {
		panic("config is not ready")
	}
	baseApi(engine)
}

func baseApi(engine *gin.Engine) {
	api := engine.Group(config.Config().Yaml.Http.BaseAPI)

	apiV1(api)
}

func apiV1(baseApi *gin.RouterGroup) {
	api := baseApi.Group("/v1")

	globalApiV1(api)
	secretApiV1(api)
	testApiV1(api)
	resourceApiV1(api)
}

func resourceApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/fl")
	middleware.ResourceUse(api)

	api.GET("/img", image.Handler) // baseApi/v1/fl/img
	api.GET("/vio", video.Handler) // baseApi/v1/fl/vio
}

func globalApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/gl")
	middleware.GlobalUse(api)

	loginAndRegisterApiV1(api)
	classApiV1(api)
	configApiV1(api)
	wupinApiV1(api)
	xieyiApiV1(api)
}

func loginAndRegisterApiV1(api *gin.RouterGroup) {
	api.GET("/lg", login.Handler)
	api.GET("/rg", register.Handler)
}

func configApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/cfg")
	api.GET("/i", getconfig.Handler)
}

func classApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/cls")
	api.GET("/lst", getclasslst.Handler)
}

func wupinApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/wp")
	api.GET("/i", getwupin.Handler)
	api.GET("/lst", getsearch.Handler)
	api.GET("/lst/h", gethotwupin.Handler)
	api.GET("/lst/s", getsearch.Handler)
}

func xieyiApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/xy")
	api.GET("/i", getxieyi.Handler)
}

func secretApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/sr")
	middleware.SecretUse(api)

	userApiV1(api)
	adminApiV1(api)
	rootAdminApiV1(api)
}

func userApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ur")
	userEditApiV1(api)
	userBuyRecordApiV1(api)
	userBagApiV1(api)
	userKefuApiV1(api)

	api.GET("/i", getuserinfo.Handler)
}

func userKefuApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/kf")

	api.POST("/msg", sendmsg.Handler)
}

func userEditApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ed")

	api.POST("/i", updateuserinfo.Handler)
	api.POST("/p", updateuserpassword.Handler)
	api.POST("/a", updateuseravtar.Handler)
}

func userBuyRecordApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/br")
	userFahuoApiV1(api)
	userDaohuoApiV1(api)
	userTuihuoApiV1(api)
	userPayApiV1(api)

	api.GET("/i", getbuyrecord.Handler)
	api.GET("/lst", getbuyrecordlst.Handler)
	api.GET("/lst/i", getbuyrecordlst.Handler)
	api.GET("/lst/p", getbuyrecordlstbypage.Handler)
}

func userFahuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/fho")

	api.POST("/chu", fahuochangeuser.Handler)
	api.POST("/qx", fahuoquxiaoshenqing.Handler)
}

func userDaohuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/dho")

	api.POST("/cfm", daohuo.Handler)
}

func userTuihuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/dho")

	api.POST("/sq", tuihuoshenqing.Handler)
	api.POST("/dj", tuihuodengji.Handler)
}

func userPayApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/py")

	api.POST("/n", newpay.Handler)
	api.POST("/b", bagpay.Handler)
	api.POST("/r", repay.Handler)
}

func userBagApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/bg")

	api.POST("/ad", addbag.Handler)
	api.GET("/lst", getbaglst.Handler)
}

func adminApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ad")
	api.Use(middleware.MustAdminXTokenMiddleware())

	adminUserApiV1(api)
	adminWupinApiV1(api)
	adminBagApiV1(api)
	adminBuyRecordApiV1(api)
	adminClassApiV1(api)
	adminClassApiV1(api)
	adminXieyiApiV1(api)
	adminMsgApiV1(api)
	adminFileUploadApiV1(api)
}

func adminFileUploadApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/fu")

	api.POST("/img", adminimageupload.Handler)
	api.POST("/vio", adminvideoupload.Handler)
}

func adminMsgApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/m")

	api.GET("/lst", admingetmsg.Handler)
}

func adminXieyiApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/xy")

	api.GET("/i", admingetxieyi.Handler)
	api.GET("/ed", adminupdatexieyi.Handler)
}

func adminClassApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/cls")
	adminUpdateClassApiV1(api)
	adminAddClassApiV1(api)

	api.GET("/i", admingetclass.Handler)
	api.GET("/lst", admingetclasslst.Handler)
}

func adminUpdateClassApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ed")
	api.POST("/n", adminchangeclassname.Handler)
	api.POST("/s", adminchangeclassshow.Handler)
	api.POST("/c", adminupdateclass.Handler)
}

func adminAddClassApiV1(api *gin.RouterGroup) {
	api.POST("/n", adminaddclass.Handler)
}

func adminUserApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ur")
	adminUpdateUserApiV1(api)

	api.GET("/i", admingetuserinfo.Handler)
	api.GET("/lst", admingetuserlst.Handler)
	api.GET("/m", admingetusermsg.Handler)

	api.POST("/n", adminadduser.Handler)
}

func adminUpdateUserApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ed")

	api.POST("/i", adminupdateuserinfo.Handler)
	api.POST("/p", adminupdateuserpassword.Handler)
	api.POST("/a", adminupdateuseravtar.Handler)
	api.POST("/ph", adminupdateuserphone.Handler)
}

func adminWupinApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/wp")

	api.POST("/n", adminadduser.Handler)
	api.POST("/ed", adminupdatewupin.Handler)

	api.GET("/i", admingetwupin.Handler)
	api.GET("/lst", admingetwupinlst.Handler)
}

func adminBagApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/bg")

	api.POST("/ad", adminaddbag.Handler)
	api.GET("/lst", admingetbag.Handler)
}

func adminBuyRecordApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/br")
	adminFahuoApiV1(api)
	adminDaohuoApiV1(api)
	adminPayApiV1(api)
	adminTuihuoApiV1(api)

	api.GET("/i", admingetbuyrecordinfo.Handler)
	api.GET("/lst", admingetuserbuyrecordlstbypage.Handler)
	api.GET("/lst/u", admingetuserbuyrecordlstbypage.Handler)
	api.GET("/lst/a", admingetbuyrecordbypage.Handler)
}

func adminTuihuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/tho")

	api.POST("/ac", admintuihuoaccept.Handler)
	api.POST("/dho", admintuihuodaohuo.Handler)
	api.POST("/dj", admintuihuodengji.Handler)
	api.POST("/sq", admintuihuoshenqing.Handler)
}

func adminPayApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/py")

	api.POST("/pp", adminpeoplepay.Handler)
	api.POST("/qx", adminquxiaopay.Handler)
}

func adminFahuoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/fho")
	adminFahuoChangeInfoApiV1(api)

	api.POST("/acqx", adminacceptfahuoquxiao.Handler)
	api.POST("/dj", adminfahuodengji.Handler)
	api.POST("/qx", adminfahuoquxiao.Handler)
}

func adminFahuoChangeInfoApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ch")

	api.POST("/ur", adminfahuochuangeuser.Handler)
	api.POST("/sp", adminfahuochangeshop.Handler)
}

func adminDaohuoApiV1(apiV1 *gin.RouterGroup) {
	apiV1.POST("/dho", admindaohuo.Handler)
}

func rootAdminApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/rt")
	api.Use(middleware.MustRotAdminXTokenMiddleware())

	rootAdminConfigApiV1(api)
}

func rootAdminConfigApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/cfg")

	api.POST("/i", admingetconfig.Handler)
	api.POST("/d", admindeleteconfig.Handler)
	api.POST("/up", adminupdateconfigpic.Handler)
	api.POST("/us", adminupdateconfigstring.Handler)
}

func testApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ts")

	testGlobalApi(api)
	testSecretApiV1(api)
}

func testGlobalApi(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/gl")
	middleware.TestGlobalUse(api)
}

func testSecretApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/sr")
	middleware.TestSecretUse(api)

	testUserApiV1(api)
}

func testUserApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/ur")
	testUserPayApiV1(api)
}

func testUserPayApiV1(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/py")

	api.POST("/p", testpay.Handler)
}
