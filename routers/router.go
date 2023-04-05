package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	//后台页面托管
	//r.LoadHTMLFiles("static/admin/index.html")
	//r.Static("admin/static", "static/admin/static")
	//r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
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
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUser)
		router.GET("categories", v1.GetCategory)
		router.GET("category/:id", v1.GetCateInfo)
		router.GET("articles", v1.GetArticle)
		router.GET("cate/artlist/:id", v1.GetArticleByCate)
		router.GET("article/:id", v1.GetArticleInfo)
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
