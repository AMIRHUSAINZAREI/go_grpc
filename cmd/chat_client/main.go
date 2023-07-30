package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/AMIRHUSAINZAREI/go_grpc_sample/pkg"
	pb "github.com/AMIRHUSAINZAREI/go_grpc_sample/proto/chat"
)

func main() {
	// Load port number from .env file
	port, err := pkg.GetEnv("CHAT_GRPC_SERVER_PORT")
	if err != nil {
		log.Fatalf("Failed to load CHAT_GRPC_SERVER_PORT environment variable")
	}

	// Set up a connection to the server
	conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a new client using the connection
	client := pb.NewChatServiceClient(conn)

	// Start a stream to send messages to the server
	sendStream, err := client.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("Error creating send stream: %v", err)
	}

	// Start a goroutine to receive messages from the server
	go receiveMessages(sendStream)

	// Read messages from the user and send them to the server
	fmt.Println("(type exit to quit)")
	fmt.Println("Enter your name:")
	var sender string
	fmt.Scanln(&sender)

	for {
		content := readLineFromStdin()
		if content == "exit" {
			break
		}

		// Create a new ChatMessage and send it to the server
		message := &pb.ChatMessage{
			Sender:    sender,
			Content:   content,
			Timestamp: time.Now().Unix(),
		}

		if err := sendStream.Send(message); err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
	}

	// Close the send stream
	sendStream.CloseSend()
}

// Function to receive messages from the server and print them to the console
func receiveMessages(stream pb.ChatService_SendMessageClient) {
	for {
		message, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return
		}
		fmt.Printf("\n[%s] %s: %s\n", time.Unix(message.Timestamp, 0).Format("15:04:05"), message.Sender, message.Content)
	}
}

// Function to read lien from stdin
func readLineFromStdin() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}
