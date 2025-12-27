package rpc

import (
    "context"
    "github.com/Anshbir18/goscheduler/internal/rpc/taskpb"
)

type TaskServer struct {
    taskpb.UnimplementedTaskServiceServer
}

func NewTaskServer() *TaskServer {
    return &TaskServer{}
}

func (s *TaskServer) CreateTask(
    ctx context.Context,
    req *taskpb.CreateTaskRequest,
) (*taskpb.Task, error) {

    // DB logic will come tomorrow
    return &taskpb.Task{
        Id:     1,
        Payload: req.Payload,
        Status: "pending",
    }, nil
}

func (s *TaskServer) GetTask(
    ctx context.Context,
    req *taskpb.GetTaskRequest,
) (*taskpb.GetTaskResponse, error) {

    return &taskpb.GetTaskResponse{}, nil
}

func (s *TaskServer) FetchDueTasks(
    ctx context.Context,
    req *taskpb.FetchDueTasksRequest,
) (*taskpb.FetchDueTasksResponse, error) {

    return &taskpb.FetchDueTasksResponse{}, nil
}

func (s *TaskServer) UpdateTaskStatus(
    ctx context.Context,
    req *taskpb.UpdateTaskStatusRequest,
) (*taskpb.Task, error) {

    return &taskpb.Task{}, nil
}
