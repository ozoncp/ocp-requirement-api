package main

import (
	"github.com/ozoncp/ocp-requirement-api/internal/api"
	requirementApiV1Description "github.com/ozoncp/ocp-requirement-api/pkg/ocp-requirement-api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const address string = ":9999"

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	requirementApiV1Description.RegisterOcpRequirementApiServer(server, api.NewRequirementApiV1())
	log.Printf("Starting serve on %s", address)
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
