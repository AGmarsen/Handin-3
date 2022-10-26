package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	// this has to be the same as the go.mod module,
	// followed by the path to the folder the proto file is in.
	gRPC "github.com/AGmarsen/Handin-3/proto"

	"google.golang.org/grpc"
)

type Server struct {
	gRPC.UnimplementedTemplateServer        // You need this line if you have a server
	name                             string // Not required but useful if you want to name your server
	port                             string // Not required but useful if your server needs to know what port it's listening to
	clock                            int32
	mutex                            sync.Mutex // used to lock the server to avoid race conditions.
	subStrm                          map[string]gRPC.Template_SubscribeServer
}

// flags are used to get arguments from the terminal. Flags take a value, a default value and a description of the flag.
// to use a flag then just add it as an argument when running the program.
var serverName = flag.String("name", "default", "Senders name") // set with "-name <name>" in terminal
var port = flag.String("port", "5400", "Server port")           // set with "-port <port>" in terminal
var idArray = []string{"A", "B", "C", "D", "E", "F", "G"}

func main() {

	// setLog() //uncomment this line to log to a log.txt file instead of the console

	// This parses the flags and sets the correct/given corresponding values.
	flag.Parse()
	fmt.Println(".:server is starting:.")

	// starts a goroutine executing the launchServer method.
	go launchServer()

	// This makes sure that the main method is "kept alive"/keeps running
	for {
		time.Sleep(time.Second * 5)
	}
}

func launchServer() {
	log.Printf("Server %s: Attempts to create listener on port %s\n", *serverName, *port)

	// Create listener tcp on given port or default port 5400
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Printf("Server %s: Failed to listen on port %s: %v\n", *serverName, *port, err) //If it fails to listen on the port, run launchServer method again with the next value/port in ports array
		return
	}

	// makes gRPC server using the options
	// you can add options here if you want or remove the options part entirely
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// makes a new server instance using the name and port from the flags.
	server := &Server{
		name:    *serverName,
		port:    *port,
		clock:   0,
		subStrm: make(map[string]gRPC.Template_SubscribeServer),
	}

	gRPC.RegisterTemplateServer(grpcServer, server) //Registers the server to the gRPC server.

	log.Printf("Server %s: Listening on port %s\n", *serverName, *port)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	// code here is unreachable because grpcServer.Serve occupies the current thread.
}

func (s *Server) Join(ctx context.Context, joinrequest *gRPC.Empty) (*gRPC.Lamport, error) {
	s.mutex.Lock()
	s.clock++
	giveId := ""
	
	for _, i := range idArray { //for all id's A, B, C ... 
		_, exists := s.subStrm[i]  //is it already in subStrm?
		if !exists {
			giveId = i
			break
		}
	}

	if giveId == "" {
		defer s.mutex.Unlock()
		log.Printf("User denied due to maximum capacity reached (S, %d)\n", s.clock)
		return &gRPC.Lamport{Id: "", Clock: s.clock, Content: "Too many users are already connected, try another time"}, nil
	} else {
		defer s.mutex.Unlock()
		log.Printf("New user accepted (S, %d)\n", s.clock)
		ack := &gRPC.Lamport{Id: giveId, Clock: s.clock, Content: fmt.Sprintf("Mr. %s has joined", giveId)}
		s.NotifyAll(ack.Content)
		return ack, nil
	}
}

func (s *Server) Send(ctx context.Context, message *gRPC.Lamport) (*gRPC.Empty, error) {
	s.mutex.TryLock()
	defer s.mutex.Unlock()
	s.clock = max(s.clock, message.Clock)

	s.NotifyAll(message.Content)
	return &gRPC.Empty{}, nil
}

func (s *Server) NotifyAll(message string) {
	s.Print(&gRPC.Lamport{Id: "S", Clock: s.clock, Content: message})
	for i, stream := range s.subStrm {
		s.clock++
		log.Printf("Sent to %s (S, %d)\n", i, s.clock)
		stream.Send(&gRPC.Lamport{Id: "S", Clock: s.clock, Content: message})
	}
}

// https://github.com/itisbsg/grpc-push-notif/blob/master/client/client.go
func (s *Server) Subscribe(stream gRPC.Template_SubscribeServer) error {
	var clientId string
	for {
		client, err := stream.Recv()
		
		if err == io.EOF {
			return nil
		}
		
		if err != nil {
			s.mutex.Lock()
			delete(s.subStrm, clientId)
			s.NotifyAll("Mr. " + clientId + " has disconnected")
			s.mutex.Unlock()
			return err
		}
		clientId = client.Id
		s.mutex.Lock()
		s.clock = max(s.clock, client.Clock) + 1
		log.Printf("Presence of Mr. %s received (S, %d)\n", client.Id, s.clock)
		
		s.subStrm[client.Id] = stream
		s.mutex.Unlock()
	}
}

func (s *Server) Print(msg *gRPC.Lamport) {
	s.clock = int32(msg.Clock) + 1
	log.Printf("%s (S, %d)\n", msg.Content, s.clock)
}

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
