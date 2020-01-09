package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"index.tmpl",gin.H{
		"title":"hello gin " + strings.ToLower(ctx.Request.Method)+" method",
	})
}
