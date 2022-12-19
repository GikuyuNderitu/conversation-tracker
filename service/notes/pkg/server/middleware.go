package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

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

func gatewayLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("[Path]: %s", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			fmt.Println("In options block")
			fmt.Printf("Path in options block: [Path] -> %s\n", r.URL.Path)
			fmt.Println(r.Header.Get("Access-Control-Request-Method"))
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length")
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", time.Hour*24*30))

		handler.ServeHTTP(w, r)
	})
}
