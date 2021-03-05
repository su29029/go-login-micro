package handler

import (
	"context"
	"database/sql"

	usersrv "user-srv/proto"

	log "github.com/micro/micro/v3/service/logger"
	db "github.com/su29029/go-login-micro/basic/db"
)

// UserSrv user handlers
type UserSrv struct{}

// Login UserSrv login
func (e *UserSrv) Login(ctx context.Context, req *usersrv.LoginRequest, rsp *usersrv.UserSrvResponse) error {
	var username string
	log.Info("Received UserSrv.Login request")
	o := db.GetDB()
	err := o.QueryRow("select username from user where username = ? and password = ?", req.Username, req.Password).Scan(&username)
	if err != nil {
		rsp.Msg = "fail"
		rsp.Status = 401
	} else {
		rsp.Msg = "success"
		rsp.Status = 200
	}
	return nil
}

// Register UserSrv register
func (e *UserSrv) Register(ctx context.Context, req *usersrv.RegisterRequest, rsp *usersrv.UserSrvResponse) error {
	var username string
	log.Info("Received UserSrv.Register request")
	if req.Username == "" || req.Password == "" {
		rsp.Msg = "invalid parameters"
		rsp.Status = 400
		return nil
	}
	o := db.GetDB()
	err := o.QueryRow("select username from user where username = ?", req.Username).Scan(&username)
	if username != "" {
		rsp.Msg = "fail"
		rsp.Status = 401
	} else {
		if err == sql.ErrNoRows {
			_, err := o.Exec("insert into user (username, password) values (?, ?)", req.Username, req.Password)
			if err == nil {
				rsp.Msg = "success"
				rsp.Status = 200
			} else {
				rsp.Msg = "fail"
				rsp.Status = 500
			}
		} else {
			rsp.Msg = "fail"
			rsp.Status = 500
		}
	}

	return nil
}

// Validation UserSrv validation
func (e *UserSrv) Validation(ctx context.Context, req *usersrv.ValidationRequest, rsp *usersrv.UserSrvResponse) error {
	log.Info("Received UserSrv.Validation request")
	rsp.Msg = "success"
	rsp.Status = 200

	return nil
}
