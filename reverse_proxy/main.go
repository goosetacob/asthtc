package main

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/goosetacob/asthtc/proto/toolsService"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port        string
	rpcEndpoint string
)

func main() {
	// configure logrus timestamp
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	// configure rpc endpoint
	rpcEndpoint := os.Getenv("RPC_ENDPOINT")
	switch len(rpcEndpoint) {
	case 0:
		rpcEndpoint = "localhost:80"
		logrus.Printf("Environment variable RPC_ENDPOINT is undefined. Forwarding requests to %v by default\n", rpcEndpoint)
	default:
		logrus.Printf("Listenining to RPC endpoint %s", rpcEndpoint)
	}

	// configure port address
	port := os.Getenv("PORT")
	switch len(port) {
	case 0:
		port = ":8080"
		logrus.Printf("Environment variable PORT is undefined. Serving to port %v by default", port)
	default:
		port = ":" + port
		logrus.Printf("Serving on port %s", port)
	}

	// run proxy server
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := toolsService.RegisterToolsHandlerFromEndpoint(ctx, mux, rpcEndpoint, opts); err != nil {
		logrus.Fatal(err)
	}

	if err := http.ListenAndServe(port, mux); err != nil {
		logrus.Fatal(err)
	}
}
