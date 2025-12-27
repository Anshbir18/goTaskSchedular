package storage

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Store struct {
	DB *sql.DB
}

func NewStore(dsn string) (*Store, error){
    println("USING DSN:", dsn)
    db, err := sql.Open("mysql", dsn)
    if err!= nil{
        return nil, err
    }

    return &Store{DB: db}, nil
}
