CREATE TABLE task_attempts (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    task_id BIGINT NOT NULL,
    attempt_number INT NOT NULL,
    started_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    finished_at DATETIME NULL,
    error TEXT NULL,

    FOREIGN KEY (task_id) REFERENCES tasks(id)
);
