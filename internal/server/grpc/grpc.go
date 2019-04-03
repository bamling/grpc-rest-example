package grpc

import (
	"context"
	"fmt"
	"github.com/bamling/grpc-rest-example/pkg/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

func StartServer(ctx context.Context, service echo.EchoServiceServer, port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrapf(err, "failed to start listener on port %s", port)
	}

	server := grpc.NewServer()
	echo.RegisterEchoServiceServer(server, service)

	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	return server.Serve(listen)
}
