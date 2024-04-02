package internal

import (
	"log"
	"net"

	"github.com/kinjalrk2k/Kaho-aur-Suno/lib/go/proto/broker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer() {
	var address = "0.0.0.0" + ":" + "8000"

	log.Println("Starting TCP Listener")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	broker.RegisterBrokerServer(server, &BrokerService{
		Topics: make(map[string][]chan *broker.Baat),
	})

	log.Println("Starting GRPC Server at " + address)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve grpc server: ", err)
	}
}
