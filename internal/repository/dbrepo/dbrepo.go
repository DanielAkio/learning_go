package dbrepo

import (
	"database/sql"

	"github.com/DanielAkio/learning_go/internal/config"
	"github.com/DanielAkio/learning_go/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
