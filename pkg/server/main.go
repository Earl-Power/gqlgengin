package server

import (
	"github.com/Earl-Power/gqlgengin/internal/handlers"
	"github.com/gin-gonic/gin"
)

//var host, port, mode string
//
//func init() {
//	host = utils.MustGet("GQL_SERVER_HOST")
//	port = utils.MustGet("GQL_SERVER_PORT")
//	mode = utils.MustGet("GQL_SERVER_MODE")
//}

func Run() {
	host := "localhost"
	port := "7000"
	//gin run models gin.ReleaseMode & gin.DebugMode
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/test", handlers.Heartbeat())
	r.Run(host + ":" + port)
}
