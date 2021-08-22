package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-requirement-api/internal/api"
	"github.com/ozoncp/ocp-requirement-api/internal/config"
	repo "github.com/ozoncp/ocp-requirement-api/internal/repository"
	requirementApiV1Description "github.com/ozoncp/ocp-requirement-api/pkg/ocp-requirement-api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

func initRepo(cfg config.Config) repo.RequirementsRepo {

	db, err := sqlx.Connect("pgx", cfg.Database.Dsn)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}

	log.Printf("Connection to database has been established.")

	return repo.NewRequirementsRepo(*db)
}

func runREST(cfg config.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := requirementApiV1Description.RegisterOcpRequirementApiHandlerFromEndpoint(ctx, mux, cfg.Grpc.Address, opts)
	if err != nil {
		panic(err)
	}

	log.Printf("Starting serve http on %s", cfg.Rest.Address)

	if err := http.ListenAndServe(cfg.Rest.Address, mux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg, err := config.ReadConfigYAML(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.Listen("tcp", cfg.Grpc.Address)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	repository := initRepo(cfg)
	requirementApiV1Description.RegisterOcpRequirementApiServer(server, api.NewRequirementApiV1(repository))
	go runREST(cfg)

	log.Printf("Starting serve grpc on %s", cfg.Grpc.Address)
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
