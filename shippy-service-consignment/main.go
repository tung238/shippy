// shippy-service-consignment/main.go
package main

import (
	"context"
	"log"
	"os"

	// Import the generated protobuf code
	"github.com/asim/go-micro/v3"
	pb "github.com/tung238/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/tung238/shippy/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	// Init will parse the command line flags
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("shippy.service.client", service.Client())
	h := &handler{repository, vesselClient}
	pb.RegisterShippingServiceHandler(service.Server(), h)

	// Run server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
