package routes

import (
	"gatewaysvr/config"
	"gatewaysvr/controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func InitRouterAndServe() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/register", controller.Register)
	}

	if err := r.Run(":" + strconv.Itoa(config.GetGlobalConfig().SvrConfig.Port)); err != nil {
		log.Error("start server err:" + err.Error())
	}
}
