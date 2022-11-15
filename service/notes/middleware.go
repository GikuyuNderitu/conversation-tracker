package main

import (
	"context"

	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func middleWare() []grpc.ServerOption {
	logger := logrus.New()
	// logger.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	logger.Formatter.(*logrus.TextFormatter).PadLevelText = true
	logger.Formatter.(*logrus.TextFormatter).QuoteEmptyFields = true
	loggingEntry := logrus.NewEntry(logger)
	var logOpts []grpc_logrus.Option
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_logrus.UnaryServerInterceptor(loggingEntry, logOpts...),
			grpc_logrus.PayloadUnaryServerInterceptor(loggingEntry, loggingDecider()),
		),
	}
}

func loggingDecider() grpc_logging.ServerPayloadLoggingDecider {
	return func(
		ctx context.Context,
		fullMethodName string,
		servingObject interface{},
	) bool {
		return true
	}
}
