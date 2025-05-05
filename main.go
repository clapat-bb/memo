package main

import (
	"fmt"

	"github.com/clapat-bb/memo/config"
	"github.com/clapat-bb/memo/logger"
	"github.com/clapat-bb/memo/model"
	"github.com/clapat-bb/memo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	logger.Log.Info("log system start!")

	model.InitDB()

	r := gin.Default()
	router.SetupRouter(r)
	port := config.Config.Server.Port
	addr := fmt.Sprintf(":%d", port)
	logger.Log.Infof("servers start... listen addr: ", addr)
	r.Run(addr)
	fmt.Println(r.Routes())
}
