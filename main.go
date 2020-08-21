package main

import (
	"fmt"
	"test/docs"
	"test/middleware"
	// "test/models"
	"test/router"
	"test/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)




func init() {
	//读取配置
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err !=nil{
		fmt.Println("获取配置错误:"+err.Error())
	} 
	//设置swagger
	docs.SwaggerInfo.Host =util.GetIP()+":"+viper.GetString("server.port")
	docs.SwaggerInfo.BasePath = "/api"
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com
// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email me@razeen.me
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	app := gin.New()
	//数据库初始化
	// models.SetUp()
	//中间件初始化 
	middleware.SetUp(app)
	//路由初始化
	router.SetUp(app)
	//静态路径初始化
	app.Static("/public","./public")
	//swagger初始化
	app.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("Swagger Api At:http://"+docs.SwaggerInfo.Host+"/swagger/index.html")
	app.Run(util.GetIP()+":"+viper.GetString("server.port"))
}

