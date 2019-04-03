package http

import (
	"context"
	"fmt"
	"github.com/bamling/grpc-rest-example/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

func StartServer(ctx context.Context, grpcPort, httpPort int) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := echo.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", grpcPort), opts); err != nil {
		return errors.Wrapf(err, "failed to register endpoint handler with port %d", grpcPort)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: mux,
	}

	// graceful shutdown of http server
	go func() {
		<-ctx.Done()
		httpCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(httpCtx)
	}()

	return server.ListenAndServe()
}
