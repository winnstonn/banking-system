package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/banking-system/internal/config"
	_ "github.com/lib/pq"
)

type IDatabase interface {
	Login() string
	Close()
}

type Repo struct {
	dbConn *sql.DB
}

func NewRepository(dbConfig *config.DatabaseConfig) IDatabase {
	if dbConfig == nil {
		panic(errors.New("failed to obtain DB config"))
	}
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return &Repo{
		dbConn: db,
	}
}

func (r *Repo) Close() {
	err := r.dbConn.Close()
	if err != nil {
		log.Println("Failed to close DB connection", err)
		return
	}

	log.Println("shutting DB down")
}
