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
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUser)
		router.GET("categories", v1.GetCategory)
		router.GET("category/:id", v1.GetCateInfo)
		router.GET("articles", v1.GetArticle)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)
		router.GET("cate/artlist/:id", v1.GetArticleByCate)
		router.GET("article/:id", v1.GetArticleInfo)
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
