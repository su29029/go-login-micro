package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	loginsrv "login-srv/proto"
)

type LoginSrv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *LoginSrv) Call(ctx context.Context, req *loginsrv.Request, rsp *loginsrv.Response) error {
	log.Info("Received LoginSrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *LoginSrv) Stream(ctx context.Context, req *loginsrv.StreamingRequest, stream loginsrv.LoginSrv_StreamStream) error {
	log.Infof("Received LoginSrv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&loginsrv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *LoginSrv) PingPong(ctx context.Context, stream loginsrv.LoginSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&loginsrv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
