package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Register(ctx *gin.Context) {
	var err error
	userName := ctx.Query("username")
	password := ctx.Query("password")

	log.Info(userName, password)

	resp, err := utils.GetUserSvrClient().Register(ctx, &pb.RegisterRequest{
		Username: userName,
		Password: password,
	})

	if err != nil {
		log.Error("register error", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", resp)

}
