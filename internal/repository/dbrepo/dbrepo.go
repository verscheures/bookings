package dbrepo

import (
	"database/sql"

	"github.com/verscheures/bookings/internal/config"
	"github.com/verscheures/bookings/internal/repository"
)

type postgressDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgreRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo{
	return &postgressDBRepo{
		App: a,
		DB: conn,
	}
}
