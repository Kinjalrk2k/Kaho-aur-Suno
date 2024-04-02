package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/kinjalrk2k/Kaho-aur-Suno/lib/go/proto/broker"
)

func main() {
	log.Println("Publisher")
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	client := broker.NewBrokerClient(conn)

	baat, err := client.Kaho(context.Background(), &broker.KahoRequest{
		Topic:   "one",
		Message: "Hello?",
	})
	if err != nil {
		log.Fatal("failed to publish: ", err)
	}
	log.Println("Published: ", baat)

	baat, err = client.Kaho(context.Background(), &broker.KahoRequest{
		Topic:   "two",
		Message: "Hey?",
	})
	if err != nil {
		log.Fatal("failed to publish: ", err)
	}
	log.Println("Published: ", baat)
}
