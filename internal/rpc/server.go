package rpc

import (
	"context"
	"fmt"

	"github.com/Anshbir18/goscheduler/internal/models"
	"github.com/Anshbir18/goscheduler/internal/rpc/taskpb"
	"github.com/Anshbir18/goscheduler/internal/storage"
)

type TaskServer struct {
	taskpb.UnimplementedTaskServiceServer
	store storage.TaskStore
}

func NewTaskServer(store storage.TaskStore) *TaskServer {
	return &TaskServer{store: store}
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

func (s *TaskServer) CreateTask(
	ctx context.Context,
	req *taskpb.CreateTaskRequest,
) (*taskpb.Task, error) {

	task := &models.Task{
		Type:        req.Type,
		Payload:     fmt.Sprintf(`"%s"`, req.Payload),
		Status:      "pending",
		ExecuteAt:   req.ExecuteAt.AsTime(),
		MaxAttempts: int(req.MaxAttempts),
	}

	id, err := s.store.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	task.ID = id

	return &taskpb.Task{
		Id:          task.ID,
		Name:        task.Type,
		Payload:     task.Payload,
		Status:      task.Status,
		MaxAttempts: int32(task.MaxAttempts),
	}, nil
}
