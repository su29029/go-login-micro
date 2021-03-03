package handler

import (
	"context"
	"net/http"

	"github.com/micro/micro/v3/service"

	"github.com/gin-gonic/gin"

	proto "github.com/su29029/go-login-micro/user-srv/proto"
)

var (
	loginClient proto.UserSrvService
)

type Error struct {
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	srv := service.New()
	loginClient = proto.NewUserSrvService("go.micro.user.srv", srv.Client()) // create a new object of UserSrvService, not open up the service.
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", loginHandler)

	r.POST("/register", registerHandler)

	// test loginHandler conveniently
	r.GET("/login", loginHandler)

	return r
}

func loginHandler(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.ShouldBind(&user)
	rsp, err := loginClient.Login(context.Background(), &proto.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "internal server error",
		})
	}
	switch rsp.Status {
	case 200:
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  rsp.Msg,
		})
	case 400:
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "invalid parameters",
		})
	case 401:
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "wrong username or password",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "internal server error",
		})
	}
}

func registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success register",
	})
}
