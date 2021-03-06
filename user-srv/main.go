package main

import (
	"user-srv/handler"
	pb "user-srv/proto"

	"github.com/su29029/go-login-micro/basic"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {

	basic.Init()
	// Create service
	srv := service.New(
		service.Name("go.micro.user.srv"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUserSrvHandler(srv.Server(), new(handler.UserSrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
