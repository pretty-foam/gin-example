package router 

import (
	"test/controller"
	"github.com/gin-gonic/gin"
)



func user(app *gin.RouterGroup)  {
	router := app.Group("/user")
	userCtr := controller.User{}
	router.POST("/get",userCtr.Get)
	router.GET("/captcha",userCtr.Captcha)
	router.GET("/captchaVerify",userCtr.CaptchaVerify)
	router.GET("/login",userCtr.Setjwt)
	router.GET("/test",userCtr.Test)
}