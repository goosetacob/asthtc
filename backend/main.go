package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/goosetacob/asthtc/backend/resource"

	"github.com/Sirupsen/logrus"
	pb "github.com/goosetacob/asthtc/api"
	"golang.org/x/net/context"
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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logrus.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	lis, err := net.Listen("tcp", portAddress)
	if err != nil {
		logrus.Fatalf("could not listen on port %d: %v", portAddress, err)
	}

	s := grpc.NewServer()
	pb.RegisterToolsServer(s, server{})
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
}

type server struct{}

func (server) Voweless(ctx context.Context, job *pb.VowelessJob) (*pb.Response, error) {
	vowelessPhrase, err := tool.Voweless(job.Phrase)
	if err != nil {
		return nil, err
	}

	res := &pb.Response{Phrase: vowelessPhrase}
	return res, nil
}

func (server) Aesthetic(ctx context.Context, job *pb.AestheticJob) (*pb.Response, error) {
	aestheticPhrase, err := tool.Aesthetic(job.Phrase)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Phrase: aestheticPhrase}, nil
}

func (server) DeBruijn(ctx context.Context, job *pb.DeBruijnJob) (*pb.Response, error) {
	return nil, fmt.Errorf("not implemented")
}
