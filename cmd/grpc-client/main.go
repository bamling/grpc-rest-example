package main

import (
	"context"
	"flag"
	"github.com/bamling/grpc-rest-example/pkg/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

var (
	serverAddr  = flag.String("server-addr", "localhost:8082", "the gRPC server address (ip:port)")
	echoMessage = flag.String("echo-message", "Hello world!", "the echo message sent to the server")
)

func main() {
	// parse command line params
	flag.Parse()

	// configure logging in JSON format
	log.SetFormatter(&log.JSONFormatter{})

	// attempt to create insecure gRPC connection
	connection, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Fatal("failed to establish connection with address: %s", *serverAddr)
	}
	defer func() {
		_ = connection.Close()
	}()

	client := echo.NewEchoServiceClient(connection)

	// attempt to send request with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.Echo(ctx, &echo.EchoMessage{Value: *echoMessage})
	if err != nil {
		log.WithError(err).Fatal("failed to send echo request to server")
	}

	log.Infof("Response: %#v", *response)
}
