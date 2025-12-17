package main

import (
	"fmt"
	"log"
	"net/http"
	"session-18/database"
	"session-18/handler"
	"session-18/repository"
	"session-18/router"
	"session-18/service"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	r := router.NewRouter(handler)

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error server")
	}
}
