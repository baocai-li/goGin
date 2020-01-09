package handler

import (
	"GinHello/model"
	"GinHello/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}

func UserSaveByQuery(ctx *gin.Context) {
	username := ctx.Query("name")
	age := ctx.Query("age")
	ctx.String(http.StatusOK, "用户"+username+",年龄:"+age+"已经保存")
}

func UserSaveByDefaultQuery(ctx *gin.Context) {
	username := ctx.Query("name")
	age := ctx.DefaultQuery("age", "20")
	ctx.String(http.StatusOK, "用户"+username+",年龄:"+age+"已经保存")
}

func UserRegister(ctx *gin.Context) {
	var user model.UserModel
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.String(http.StatusBadRequest, "输入的数据不合法")
		log.Println("err ->", err.Error())
	}
	//email := ctx.PostForm("email")
	//password := ctx.DefaultPostForm("password","Wa123456")
	//passwordAgain := ctx.DefaultPostForm("password-again","Wa123456")
	id := user.Save()
	log.Println("id is", id)
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(ctx *gin.Context) {
	var user model.UserModel
	if e := ctx.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		ctx.SetCookie("user_cookie", string(u.Id), 1000, "/", "localhost", false, true)
		log.Panicln("登录成功", u.Email)
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
			"id":    u.Id,
		})
	}
}

func UserProfile(ctx *gin.Context) {
	id := ctx.Query("id")
	var user model.UserModel
	i, e := strconv.Atoi(id)
	u, err := user.QueryById(i)
	if e != nil || err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	ctx.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误 ", err.Error())
	}
	file, e := context.FormFile("avatar-file")
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("文件上传错误", e.Error())
	}
	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	//log.Panicln("path:avatar:",path)
	log.Println("path:avatar:",path)
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	e = context.SaveUploadedFile(file, path+"/"+fileName)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}
	avatarUrl := "http://localhost:8080/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.Update(user.Id)
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	context.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(int(user.Id)))
}
