package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/Anshbir18/goscheduler/internal/config"
	"github.com/Anshbir18/goscheduler/internal/rpc"
	"github.com/Anshbir18/goscheduler/internal/rpc/taskpb"
	"github.com/Anshbir18/goscheduler/internal/storage"
)

func main() {
	cfg := config.LoadConfig()

	storeWrapper, err := storage.NewStore(cfg.DSN())
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	taskStore := storage.NewMySQLTaskStore(storeWrapper.DB)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(
		grpcServer,
		rpc.NewTaskServer(taskStore),
	)

	log.Println("Scheduler gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
