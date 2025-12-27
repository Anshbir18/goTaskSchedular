package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "github.com/Anshbir18/goscheduler/internal/rpc"
    "github.com/Anshbir18/goscheduler/internal/rpc/taskpb"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    taskpb.RegisterTaskServiceServer(
        grpcServer,
        rpc.NewTaskServer(),
    )

    log.Println("Scheduler gRPC server running on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
