package models

import (
	"log"
	"gorm.io/gorm"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *gorm.DB

// Init - initialize the models system
func Init(appDB *gorm.DB) {
	db = appDB
}

// ResetDBTest - resets the db for testing
func ResetDBTest() {
	m, err := migrate.New("file://models/migrations", "postgres://dev:dev@localhost:5432/af?sslmode=disable")
	if err != nil {
		log.Fatalf("Cannot start migrations: %+v", err)
	}

	if err := m.Drop(); err != nil {
		log.Fatalf("Cannot drop all migrations: %+v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("Cannot migrate up: %+v", err)
	}
}
