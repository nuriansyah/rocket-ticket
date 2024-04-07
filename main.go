package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nuriansyah/rocket-ticket/handlers"
	"github.com/nuriansyah/rocket-ticket/repository"
	"github.com/nuriansyah/rocket-ticket/services"
)

func main() {
	db, err := sql.Open("sqlite3", "./basis-app.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()
	repo := repository.NewRepository(db)

	vanueSvc := services.NewVenueService(repo)
	handlers := handlers.NewHandler(vanueSvc)
	http.HandleFunc("/venues", handlers.CreateVenue)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
