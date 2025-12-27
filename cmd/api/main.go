package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/Anshbir18/goscheduler/internal/api"
	"github.com/Anshbir18/goscheduler/internal/rpc/taskpb"
)

func main() {
	conn, err := grpc.Dial(
		"127.0.0.1:50051",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("gRPC connection failed: %v", err)
	}

	client := taskpb.NewTaskServiceClient(conn)

	handler := api.NewTaskHandler(client)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.CreateTask).Methods(http.MethodPost)

	log.Println("API server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
