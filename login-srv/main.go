package main

import (
	"login-srv/handler"
	pb "login-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("go.micro.login.srv"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterLoginSrvHandler(srv.Server(), new(handler.LoginSrv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
