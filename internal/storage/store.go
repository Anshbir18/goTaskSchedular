package storage

import (
    "database/sql"
)

type Store struct {
	DB *sql.DB
}

func NewStore(dsn string) (*Store, error){
    db, err := sql.Open("mysql", dsn)
    if err!= nil{
        return nil, err
    }

    return &Store{DB: db}, nil
}
