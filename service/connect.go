package service

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "rpc-util/service/user/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var UserConnection *grpc.ClientConn

func ConnectUserService() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	UserConnection = conn
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := client.GetName(context.Background(), &pb.NameRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}
