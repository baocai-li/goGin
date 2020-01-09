package initRouter

import (
	"GinHello/handler"
	"GinHello/handler/article"
	"GinHello/middleware"
	"GinHello/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	//router := gin.Default()
	router := gin.New()
	// 添加自定义的 logger 中间件
	router.Use(middleware.Logger(), gin.Recovery())

	if mode := gin.Mode();
	mode == gin.TestMode{
		router.LoadHTMLGlob("./../templates/*")
	}else {
		router.LoadHTMLGlob("templates/*")
	}
	//router.StaticFile("/favicon.ico","./favicon.ico")
	router.Static("/statics","./statics")
	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))
	index := router.Group("/")
	{
		index.Any("", handler.Index)

	}
	userRouter := router.Group("/user")
	{
		//userRouter.GET("/:name",handler.UserSave)
		//userRouter.GET("",handler.UserSaveByQuery)
		userRouter.POST("/register",handler.UserRegister)
		userRouter.POST("/login",handler.UserLogin)
		userRouter.POST("/update", middleware.Auth(),handler.UpdateUserProfile)
		userRouter.GET("/profile",middleware.Auth(),handler.UserProfile)

	}
	articleRouter := router.Group("")
	{
		// 通过获取单篇文章
		articleRouter.GET("/article/:id", article.GetOne)
		// 获取所有文章
		articleRouter.GET("/articles", article.GetAll)
		// 添加一篇文章
		articleRouter.POST("/article", article.Insert)
		articleRouter.DELETE("/article/:id", article.DeleteOne)
	}

	//get请求参数区别，见user_test.go
	//ctx *gin.Context 获取参数区别
	//router.GET("/user/:name",handler.UserSave)
	//router.GET("user",handler.UserSaveByQuery)
	//router.GET("user",handler.UserSaveByDefaultQuery)
	return router
}

func retHelloGinAndMethodGroup(content *gin.Context){
	content.String(http.StatusOK,
		"hello gin " + strings.ToLower(content.Request.Method) + " method")
}
