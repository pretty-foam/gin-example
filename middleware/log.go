package middleware 

import (
	"io"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
)

var(
	logFilePath="./log"
	logFileName="system.log.%Y%m%d"
)

//日志
func logInit() gin.HandlerFunc {
	path := path.Join(logFilePath,logFileName)
	file ,err := rotatelogs.New(path)
	if err !=nil {
	   println("error:",err)
	}
	gin.ForceConsoleColor()
	gin.DefaultWriter = io.MultiWriter(file,os.Stdout)
	return gin.Logger()
}