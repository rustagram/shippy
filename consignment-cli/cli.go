// shippy-cli-consignment/main.go

package main

import (
	"context"
	"encoding/json"
	pb "github.com/rustambek96/shippy/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	micro "github.com/micro/go-micro"
)

const (
	address			= "localhost:50051"
	defaultFilename	= "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	//Set up a connection to the server
	service := micro.NewService(micro.Name("shippy"))
	service.Init()

	client := pb.NewShippingService("shippy", service.Client())

	//Contact the server and print out its response
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}