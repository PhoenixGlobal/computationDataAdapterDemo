package main

import (
	"computationDataAdapterDemo/conf"
	_ "computationDataAdapterDemo/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Init init http
func InitHttp() {
	engine := gin.Default()
	outerRouter(engine)
	// swagger address: http://localhost:8088/swagger/index.html
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// listen and serve on 0.0.0.0:8088
	//engine.Run("0.0.0.0:8088")
	if err := engine.Run(conf.Conf.SwagConf.Addr); err != nil {
		panic(err)
	}
}

func outerRouter(e *gin.Engine) {
	e.GET("/monitor/ping", ping)
	noAuthGroup := e.Group("/api")
	{
		noAuthGroup .GET("/data", GetData)
	}
}

// Ping is
func ping(ctx *gin.Context) {

}