package service

import (
	"context"
	"github.com/bamling/grpc-rest-example/pkg/api"
	log "github.com/sirupsen/logrus"
)

type echoService struct{}

func NewEcho() echo.EchoServiceServer {
	return echoService{}
}

func (e echoService) Echo(context context.Context, echoMessage *echo.EchoMessage) (*echo.EchoMessage, error) {
	log.Infof("received RCP request: %#v", echoMessage)
	return echoMessage, nil
}
