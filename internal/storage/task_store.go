package storage

import (
    "context"
    // "time"
    "github.com/Anshbir18/goscheduler/internal/models"
)

type TaskStore interface {
    CreateTask(ctx context.Context, t *models.Task) (int64, error)
    // GetTask(ctx context.Context, id int64) (*models.Task, error)
    // UpdateTaskStatus(ctx context.Context, id int64, status string) error
    // FetchDueTasks(ctx context.Context, now time.Time) ([]*models.Task, error)
    // UpdateAfterRun(ctx context.Context, t *models.Task) error
}
