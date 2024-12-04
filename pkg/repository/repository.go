package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kahuri1/part_of_the_authentication_service/pkg/model"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func NewPostgresDB(cfg model.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLmode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if err := goose.Up(db.DB, "migrations"); err != nil {
		logrus.Fatal(err)
	}
	return db, nil
}
