package main

import (
	"login-cli/handler"
	pb "login-cli/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("login-cli"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterLoginCliHandler(srv.Server(), new(handler.LoginCli))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
