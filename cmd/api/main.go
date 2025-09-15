package main

import (
	"log"
	"task-manager/config"
	"task-manager/router" // Import router
)

func main() {
	// Koneksi database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Setup router dari package router dan inject dependency 'db'
	r := router.SetupRouter(db)

	// Jalankan server
	log.Println("Server started at :8080")
	r.Run(":8080")
}