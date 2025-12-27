package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Anshbir18/goscheduler/internal/rpc/taskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskHandler struct {
	client taskpb.TaskServiceClient
}

func NewTaskHandler(c taskpb.TaskServiceClient) *TaskHandler {
	return &TaskHandler{client: c}
}

type createTaskRequest struct {
	Type        string    `json:"type"`
	Payload     string    `json:"payload"`
	ExecuteAt   time.Time `json:"execute_at"`
	MaxAttempts int32     `json:"max_attempts"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.client.CreateTask(
		r.Context(),
		&taskpb.CreateTaskRequest{
			Type:        req.Type,
			Payload:     req.Payload,
			ExecuteAt:   timestamppb.New(req.ExecuteAt),
			MaxAttempts: req.MaxAttempts,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
