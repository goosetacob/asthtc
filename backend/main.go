package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/goosetacob/asthtc/backend/resource"
	"github.com/goosetacob/asthtc/proto/toolsService"
	"google.golang.org/grpc"
)

var portAddress string

func init() {
	// configure logrus timestamp
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	// configure port address
	port := os.Getenv("PORT")
	switch len(port) {
	case 0:
		logrus.Println("Environment variable PORT is undefined. Using port :80 by default")
		portAddress = ":80"
	case 1:
		logrus.Printf("Environment variable PORT=\"%s\"", port)
		portAddress = ":" + port
	}
}

func main() {
	lis, err := net.Listen("tcp", portAddress)
	defer lis.Close()
	if err != nil {
		logrus.Fatalf("could not listen on port %d: %v", portAddress, err)
	}

	s := grpc.NewServer()
	toolsService.RegisterToolsServer(s, server{})
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
}

type server struct{}

func (server) Voweless(ctx context.Context, job *toolsService.VowelessJob) (*toolsService.Response, error) {
	vowelessPhrase, err := tool.Voweless(job.Phrase)
	if err != nil {
		return nil, err
	}

	res := &toolsService.Response{Phrase: vowelessPhrase}
	return res, nil
}

func (server) Aesthetic(ctx context.Context, job *toolsService.AestheticJob) (*toolsService.Response, error) {
	aestheticPhrase, err := tool.Aesthetic(job.Phrase)
	if err != nil {
		return nil, err
	}

	return &toolsService.Response{Phrase: aestheticPhrase}, nil
}

func (server) DeBruijn(ctx context.Context, job *toolsService.DeBruijnJob) (*toolsService.Response, error) {
	return nil, fmt.Errorf("not implemented")
}
