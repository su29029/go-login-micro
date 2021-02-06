package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	logincli "login-cli/proto"
)

type LoginCli struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *LoginCli) Call(ctx context.Context, req *logincli.Request, rsp *logincli.Response) error {
	log.Info("Received LoginCli.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *LoginCli) Stream(ctx context.Context, req *logincli.StreamingRequest, stream logincli.LoginCli_StreamStream) error {
	log.Infof("Received LoginCli.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&logincli.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *LoginCli) PingPong(ctx context.Context, stream logincli.LoginCli_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&logincli.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
