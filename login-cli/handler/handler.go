package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/micro/micro/v3/service"

	"github.com/gin-gonic/gin"

	proto "github.com/su29029/go-login-micro/login-srv/proto"
)

var (
	loginClient proto.LoginSrvService
)

type Error struct {
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	srv := service.New()
	loginClient = proto.NewLoginSrvService("go.micro.login.srv", srv.Client())
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", loginHandler)

	r.POST("/register", registerHandler)

	r.GET("/login", loginHandler)

	return r
}

func loginHandler(c *gin.Context) {
	rsp, err := loginClient.Call(context.Background(), &proto.Request{
		Say: "1",
	})
	fmt.Println(rsp)
	rsp2, err := loginClient.Login(context.Background(), &proto.LoginRequest{
		Username: "su29029",
		Password: "123456a",
	})
	fmt.Println("1")
	fmt.Println(rsp2)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
