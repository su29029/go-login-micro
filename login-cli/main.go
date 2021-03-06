package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"

	"github.com/su29029/go-login-micro/login-cli/handler"
)

func main() {

	router := handler.InitRouter()

	service := web.NewService(
		web.Name("go.micro.login.web"),
		web.Version("latest"),
		web.Address(":18088"),
		web.Handler(router),
	)
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// init handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
