package main

import (
	"context"
	"github.com/Akongstad/gRPC-time/Time"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":9000"
)

type server struct {
	Time.UnimplementedTimeServiceServer
}

func (s *server) GetTime(ctx context.Context, r *Time.Message) (*Time.Message, error) {
	log.Printf("Received msg %r", r.GetBody())
	return &Time.Message{Body: time.Now().String()}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}
	s := server{}
	grpcServer := grpc.NewServer()

	Time.RegisterTimeServiceServer(grpcServer, &s)

	log.Printf("server listening at %v", lis.Addr())
	//fmt.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed serve server: %v", err)
	}
}
