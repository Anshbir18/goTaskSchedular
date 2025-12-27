package models

import "time"

type Task struct {
    ID          int64      `json:"id"`
    Type        string     `json:"type"`
    Payload     string     `json:"payload"`
    Status      string     `json:"status"`
    ExecuteAt   time.Time  `json:"execute_at"`
    NextRunAt   *time.Time `json:"next_run_at"`
    Attempts    int        `json:"attempts"`
    MaxAttempts int        `json:"max_attempts"`
    LastError   *string    `json:"last_error"`
    Result      *string    `json:"result"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}
