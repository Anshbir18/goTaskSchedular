package storage

import(
	"context"
	"database/sql"
	"github.com/Anshbir18/goscheduler/internal/models"
)


type MySQLTaskStore struct{
	db *sql.DB
}

func NewMySQLTaskStore(db *sql.DB) *MySQLTaskStore {
	return &MySQLTaskStore{db: db}
}

func (s *MySQLTaskStore) CreateTask(ctx context.Context, t *models.Task) (int64, error) {
	query := `
		INSERT INTO tasks
		(type, payload, status, execute_at, next_run_at, max_attempts)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	res,err := s.db.ExecContext(ctx, query,
		t.Type,
		t.Payload,
		t.Status,
		t.ExecuteAt,
		t.NextRunAt,
		t.MaxAttempts,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}