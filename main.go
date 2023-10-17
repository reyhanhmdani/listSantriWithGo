package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"project1/database"
	"project1/routes"
	"project1/service"
)

func setupLogOutput() {
	f, _ := os.Create("gin-log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	ctx := context.Background()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	//ENV
	loadEnv()

	// INITAL DATABASE
	db, err := database.Db(ctx)
	if err != nil {
		return
	}

	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("Error running schema migration %v", err)
	}

	userRepo := database.NewUserRepository(db)
	santriRepo := database.NewSantriRepository(db)
	santriService := service.NewSantriService(santriRepo, userRepo)
	route := routes.NewRoute(santriService)
	routeInit := route.RouteInit()
	err = routeInit.Run(":7000")
	if err != nil {
		log.Info("ada yang salah di Route")
		log.Fatal(err)
	}

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Mengambil nilai variabel lingkungan
	dbHost := os.Getenv("DB_HOST")
	dbRootPassword := os.Getenv("DB_PASS")
	dbDatabase := os.Getenv("DB_NAME")

	// Contoh penggunaan nilai variabel lingkungan
	log.Printf("DB Host: %s", dbHost)
	log.Printf("DB Root Password: %s", dbRootPassword)
	log.Printf("DB Database: %s", dbDatabase)
}
