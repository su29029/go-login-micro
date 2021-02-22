module github.com/su29029/go-login-micro/login-cli

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.8.3
	github.com/micro/micro/v3 v3.0.0
	github.com/su29029/go-login-micro/login-srv v0.0.0-00010101000000-000000000000
)

replace github.com/su29029/go-login-micro/basic => ../basic

replace github.com/su29029/go-login-micro/login-srv => ../login-srv
