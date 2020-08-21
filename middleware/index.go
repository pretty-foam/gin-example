package middleware 

import (
	"github.com/gin-gonic/gin"
)

//SetUp 中间件初始化
func SetUp(app *gin.Engine) {
	app.Use(logInit())
	app.Use(Cors())
	app.Use(sessionConfig())
}
