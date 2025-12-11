package main

import (
	"log"

	"github.com/codesayhi/golang-clean/internal/bootstrap"
	"github.com/codesayhi/golang-clean/internal/config"
	dbconnect "github.com/codesayhi/golang-clean/internal/infrastructure/db/gormrepo/connect"
)

func main() {
	// Load config
	loadConfig := config.Load()
	url := loadConfig.DBUrl
	// Kết nối database
	db, err := dbconnect.NewGormDB(url)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Println("Connected to database", db)
	//

	app := bootstrap.NewApplication(db)

	// 3. Create Gin server
	server := bootstrap.NewServer(app)

	// 4. Run server
	log.Println("🚀 Server is running at http://localhost:8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("❌ Server error: %v", err)
	}
}
