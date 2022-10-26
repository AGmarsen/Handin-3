package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	gRPC "github.com/AGmarsen/Handin-3/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Same principle as in client. Flags allows for user specific arguments/values
var clientsName = flag.String("name", "default", "Senders name")
var serverPort = flag.String("server", "5400", "Tcp server")

var server gRPC.TemplateClient  //the server
var ServerConn *grpc.ClientConn //the server connection

var clock = int32(0)
var id = ""
var mutex *sync.Mutex = &sync.Mutex{}

func main() {
	//parse flag/arguments
	flag.Parse()

	fmt.Println("--- CLIENT APP ---")

	//connect to server and close the connection when program closes
	fmt.Println("--- join Server ---")
	ConnectToServer()
	defer ServerConn.Close()

	joinServer()
	go subscribe()

	//start the biding
	parseInput()
}

// connect to server
func ConnectToServer() {

	//dial options
	//the server is not using TLS, so we use insecure credentials
	//(should be fine for local testing but not in the real world)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	//use context for timeout on the connection
	timeContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() //cancel the connection when we are done

	//dial the server to get a connection to it
	log.Printf("client %s: Attempts to dial on port %s\n", *clientsName, *serverPort)
	conn, err := grpc.DialContext(timeContext, fmt.Sprintf(":%s", *serverPort), opts...)
	if err != nil {
		log.Fatalf("Fail to Dial : %v", err)
		return
	}

	// makes a client from the server connection and saves the connection
	// and prints rather or not the connection was is READY
	server = gRPC.NewTemplateClient(conn)
	ServerConn = conn
	log.Println("the connection is: ", conn.GetState().String())
}

func parseInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the chat!")
	fmt.Println("--------------------")

	//Infinite loop to listen for clients input.
	for {
		fmt.Print("-> ")

		//Read input into var input and any errors into err
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) //Trim input

		if !conReady(server) {
			log.Printf("Client %s: something was wrong with the connection to the server :(", *clientsName)
			continue
		}

		//Convert string to int64, return error if the int is larger than 32bit or not a number
		if err == nil {
			send(input)
		} else {
			log.Fatal(err)
		}
	}
}

func joinServer() {
	response, err := server.Join(context.Background(), &gRPC.Empty{})
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	if response.Id == "" {
		log.Fatal(response.Content)
	} else {
		id = response.Id
		print(response)
	}
}

//https://github.com/itisbsg/grpc-push-notif/blob/master/client/client.go
func subscribe() {
	mutex.Lock()
	clock++
	mutex.Unlock()
	stream, err := server.Subscribe(context.Background()) //open stream
	if err != nil {
		log.Fatalf("%v", err)
	}
	stream.Send(&gRPC.Lamport{Id: id, Clock: clock})
	for {
		rec, er := stream.Recv()
		if er == io.EOF {
			break
		}
		if er != nil {
			log.Fatalf("%v", er)
		}
		print(rec)
	}
}

func send(message string) {
	mutex.Lock()
	defer mutex.Unlock()
	_, err := server.Send(context.Background(), &gRPC.Lamport{Id: id, Clock: int32(clock) + 1, Content: message})
	if err != nil {
		log.Printf("%v", err)
	}
}

// Function which returns a true boolean if the connection to the server is ready, and false if it's not.
func conReady(s gRPC.TemplateClient) bool {
	return ServerConn.GetState().String() == "READY"
}

func print(msg *gRPC.Lamport) {
	mutex.Lock()
	defer mutex.Unlock()
	clock = max (msg.Clock, clock) + 1
	log.Printf("%s (%s, %d)\n", msg.Content, id, clock)
}

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}