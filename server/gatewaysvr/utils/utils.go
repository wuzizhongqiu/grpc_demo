package utils

import (
	"context"
	"fmt"
	"gatewaysvr/config"
	log "github.com/sirupsen/logrus"
	"github.com/wuzi/grpc_demo/proto"

	"time"
	// 必须要导入这个包，否则grpc会报错
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
)

var (
	UserSvrClient proto.UserServiceClient
)

func NewSvrConn(svrName string) (*grpc.ClientConn, error) {
	consulInfo := config.GetGlobalConfig().ConsulConfig
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, svrName),
		// grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		// grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		log.Errorf("NewSvrConn with svrname %s err:%v", svrName, err)
		return nil, err
	}
	log.Info("NewSvrConn success")
	return conn, nil
}

func GetUserSvrClient() proto. {
	return UserSvrClient
}

func NewUserSvrClient(svrName string) UserServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return NewUserServiceClient(conn)
}

func InitSvrConn() {
	UserSvrClient = NewUserSvrClient(config.GetGlobalConfig().SvrConfig.UserSvrName)
}
