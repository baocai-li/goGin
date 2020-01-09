package initRouter

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouterNoGroup() *gin.Engine {
	router := gin.Default()
	router.GET("/", retHelloGinAndMethod)
	router.POST("/", retHelloGinAndMethod)
	router.PUT("/", retHelloGinAndMethod)
	router.DELETE("/", retHelloGinAndMethod)
	router.PATCH("/", retHelloGinAndMethod)
	router.HEAD("/", retHelloGinAndMethod)
	router.OPTIONS("/", retHelloGinAndMethod)

	//get请求参数区别，见user_test.go
	//ctx *gin.Context 获取参数区别
	router.GET("/user/:name",handler.UserSave)
	//router.GET("user",handler.UserSaveByQuery)
	router.GET("user",handler.UserSaveByDefaultQuery)
	return router
}

func retHelloGinAndMethod(content *gin.Context){
	content.String(http.StatusOK,
		"hello gin " + strings.ToLower(content.Request.Method) + " method")
}
