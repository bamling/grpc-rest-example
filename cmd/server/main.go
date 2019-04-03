package main

import (
	"flag"
	"github.com/bamling/grpc-rest-example/internal/server"
	log "github.com/sirupsen/logrus"
)

var (
	grpcPort = flag.Int("grpc-port", 8082, "the gRPC server bind port")
	httpPort = flag.Int("http-port", 8080, "the http server bind port")
)

func main() {
	// parse command line params
	flag.Parse()

	// configure logging in JSON format
	log.SetFormatter(&log.JSONFormatter{})


	log.Infof("starting gRPC server on port %d and http server on port %d", *grpcPort, *httpPort)
	if err := server.New().Start(*grpcPort, *httpPort); err != nil {
		log.WithError(err).WithFields(log.Fields{
			"grpc_port": *grpcPort,
			"http_port": *httpPort,
		}).Fatal("failed to start server")
	}
	log.Info("stopped server gracefully...")
}
