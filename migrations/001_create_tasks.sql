CREATE TABLE tasks (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,

    type VARCHAR(255) NOT NULL,
    payload JSON NULL,

    status ENUM('pending', 'scheduled', 'running', 'succeeded', 'failed') 
        NOT NULL DEFAULT 'pending',

    execute_at DATETIME NOT NULL,
    next_run_at DATETIME DEFAULT NULL,

    attempts INT NOT NULL DEFAULT 0,
    max_attempts INT NOT NULL DEFAULT 3,

    last_error TEXT NULL,
    result JSON NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_status_next_run (status, next_run_at),
    INDEX idx_execute_at (execute_at)
);
