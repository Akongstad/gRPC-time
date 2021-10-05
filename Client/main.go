package main

import (
	"context"
	"github.com/Akongstad/gRPC-time/Time"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = ":9000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := Time.NewTimeServiceClient(conn)
	log.Printf("My time is", time.Now())
	time.Sleep(time.Second * 1)
	clientMessage := "Client says: Give time pls"
	log.Printf(clientMessage)
	request := Time.Message{Body: clientMessage}
	r, err := c.GetTime(context.Background(), &request)
	if err != nil {
		log.Fatalf("could not find get time: %v", err)
	}
	log.Printf("Server says: %s", r)
}
