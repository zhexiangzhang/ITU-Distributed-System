package main

import (
	"context"
	"flag"
	proto "github.com/zhexiangzhang/ITU-Distributed-System-2023/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type Server struct {
	proto.UnimplementedTimeAskServer // Necessary
	name                             string
	port                             int
}

func startServer(server *Server) {
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterTimeAskServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func (c *Server) AskForTime(ctx context.Context, in *proto.AskForTimeMessage) (*proto.TimeMessage, error) {
	log.Printf("Client with ID %d asked for the time\n", in.ClientId)
	return &proto.TimeMessage{Time: time.Now().String()}, nil
}

var port = flag.Int("port", 0, "server port number")

func main() {
	// get the port number from the command line
	flag.Parse()

	// create a server struct
	server := &Server{
		name: "Server001",
		port: *port,
	}

	// start the server
	go startServer(server)

	// Keep the server running until it is manually quit
	for {

	}
}
