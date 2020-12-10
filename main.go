package main

import (
	"log"
	"net/http"

	"github.com/jahio/af/api"
	"github.com/jahio/af/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://dev:dev@localhost:5432/af?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err.Error())
	}

	models.Init(db)
	router := api.Init(db)

	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	server.ListenAndServe()

}
