package server

import (
	"context"
	"github.com/bamling/grpc-rest-example/internal/server/grpc"
	"github.com/bamling/grpc-rest-example/internal/server/http"
	"github.com/bamling/grpc-rest-example/internal/service"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type Server struct{}

func New() Server {
	return Server{}
}

func (s Server) Start(grpcPort, httpPort int) error {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

		select {
		case sig := <-signals:
			log.Warnf("received signal %s, shutting down...", sig.String())
			cancel()
		case <-ctx.Done():
			log.Debug("closing signal handler")
			return ctx.Err()
		}

		return nil
	})
	group.Go(func() error { return grpc.StartServer(ctx, service.NewEcho(), grpcPort) })
	group.Go(func() error { return http.StartServer(ctx, grpcPort, httpPort) })

	return group.Wait()
}
