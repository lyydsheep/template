package main

import (
	"your-module-name/common/enum"
	log "your-module-name/common/logger"
	"your-module-name/config"
	"your-module-name/dal/cache"
	"your-module-name/dal/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig()
	log.InitLogger()
	cache.RedisInit()
	dao.InitGormLogger()
	dao.InitDB()
}

func main() {
	if config.App.Env == enum.ModePROD {
		gin.SetMode(gin.ReleaseMode)
	}
	g := InitializeApp()
	if err := g.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
