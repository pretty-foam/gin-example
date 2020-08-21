package router 

import (
	"github.com/gin-gonic/gin"
)


// SetUp 初始化路由
func SetUp(app *gin.Engine)  {
	router := app.Group("/api")
	user(router)
}