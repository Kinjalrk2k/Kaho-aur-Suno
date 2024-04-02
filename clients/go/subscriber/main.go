package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/kinjalrk2k/Kaho-aur-Suno/lib/go/proto/broker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("provide the topic name as a command line argument")
	}

	topic := os.Args[1]
	log.Printf("Subscribing to Topic: %v\n", topic)
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	client := broker.NewBrokerClient(conn)
	stream, err := client.Suno(context.Background(), &broker.SunoRequest{Topic: topic})
	if err != nil {
		log.Fatal("failed to stream: ", err)
	}

	for {
		baat, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("failed to read stream", err)
		}
		log.Println(baat)
	}
}
