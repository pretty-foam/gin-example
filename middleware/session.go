package middleware 

import (
	// 导入session包
	"github.com/gin-contrib/sessions"
	// 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
	// 导入gin框架包
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)



//session配置
func sessionConfig() gin.HandlerFunc {
	key := viper.GetString("session.key")
	name := viper.GetString("session.name")
	store := cookie.NewStore([]byte(key))
	sessionMaxAge := viper.GetInt("session.maxAge")
	// 初始化基于redis的存储引擎
	// 参数说明：
	//    第1个参数 - redis最大的空闲连接数
	//    第2个参数 - 数通信协议tcp或者udp
	//    第3个参数 - redis地址, 格式，host:port
	//    第4个参数 - redis密码
	//    第5个参数 - session加密密钥
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge,
		Path: "/",
	})
	return  sessions.Sessions(name,store)
}
