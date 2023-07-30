package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	"github.com/AMIRHUSAINZAREI/go_grpc_sample/pkg"
	pb "github.com/AMIRHUSAINZAREI/go_grpc_sample/proto/chat"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	clients map[pb.ChatService_SendMessageServer]bool
	mu      sync.Mutex
}

func (s *chatServer) SendMessage(stream pb.ChatService_SendMessageServer) error {
	// Register the client
	s.mu.Lock()
	s.clients[stream] = true
	s.mu.Unlock()

	// De-register the client when this function ends
	defer func() {
		s.mu.Lock()
		delete(s.clients, stream)
		s.mu.Unlock()
	}()

	for {
		fmt.Println()
		message, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return err
		}

		// Broadcast the received message to all connected clients
		s.mu.Lock()
		for client := range s.clients {
			if client != stream {
				if err := client.Send(message); err != nil {
					log.Printf("Error sending message to client: %v", err)
				}
			}
		}
		s.mu.Unlock()
	}
}

func main() {
	// Load port number form .env file
	port, err := pkg.GetEnv("CHAT_GRPC_SERVER_PORT")
	if err != nil {
		log.Fatalf("Failed to load CHAT_GRPC_SERVER_PORT environment variable")
	}

	// Create a tcp connection
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	// Register ChatServer
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, &chatServer{
		clients: make(map[pb.ChatService_SendMessageServer]bool),
	})

	// Start listening
	fmt.Println("Server started. Listening on port", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
