package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	//后台页面托管
	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.POST("user/add", v1.AddUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//分文章模块的用户接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		auth.POST("/upload", v1.UpLoad)

		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)

		// 评论设置
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)

		//招聘模块
		auth.GET("recruitments", v1.GetRecruitmentList)
		auth.POST("recruitment/add", v1.AddRecruitment)
	}
	router := r.Group("api/v1")
	{
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUser)
		router.GET("categories", v1.GetCategory)
		router.GET("category/:id", v1.GetCateInfo)
		router.GET("articles", v1.GetArticle)

		//登录控制
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)
		router.GET("cate/artlist/:id", v1.GetArticleByCate)

		router.GET("article/:id", v1.GetArticleInfo)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)

		//发送邮箱验证码
		router.GET("sendverifycode", v1.SendEmail)

	}
	r.Run(utils.HttpPort)
}
