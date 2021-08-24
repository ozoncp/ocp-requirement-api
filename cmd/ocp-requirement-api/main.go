package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-requirement-api/internal/api"
	"github.com/ozoncp/ocp-requirement-api/internal/config"
	"github.com/ozoncp/ocp-requirement-api/internal/flusher"
	"github.com/ozoncp/ocp-requirement-api/internal/kafka"
	"github.com/ozoncp/ocp-requirement-api/internal/metrics"
	repo "github.com/ozoncp/ocp-requirement-api/internal/repository"
	"github.com/ozoncp/ocp-requirement-api/internal/tracing"
	requirementApiV1Description "github.com/ozoncp/ocp-requirement-api/pkg/ocp-requirement-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

const batchSize = 10

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

func runMonitoring(cfg config.Config) {
	mux := http.NewServeMux()
	mux.Handle(cfg.Prometheus.Endpoint, promhttp.Handler())

	srv := &http.Server{
		Addr:    cfg.Prometheus.Address,
		Handler: mux,
	}

	log.Printf("Starting serve metrics on http %s", cfg.Prometheus.Address)
	if err := srv.ListenAndServe(); err != nil {
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
	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	requirementApiV1Description.RegisterOcpRequirementApiServer(
		server, api.NewRequirementApiV1(
			repository,
			flusher.NewFlusher(batchSize, repository),
			producer,
		),
	)
	closer := tracing.InitTracer(cfg)

	metrics.InitMetrics()
	go runREST(cfg)
	go runMonitoring(cfg)

	log.Printf("Starting serve grpc on %s", cfg.Grpc.Address)
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}

	if err := closer.Close(); err != nil {
		log.Fatal(err)
	}
}
