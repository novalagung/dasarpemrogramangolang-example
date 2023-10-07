package main

import (
	"clean-architecture-golang-example/config"
	"clean-architecture-golang-example/controllers"
	"clean-architecture-golang-example/repositories"
	"clean-architecture-golang-example/usecases"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		panic(fmt.Errorf("Error Environment : %s", err))
	}

	appConfig := struct {
		name       string
		host       string
		port       string
		versionapi string
	}{
		name:       os.Getenv("APP_NAME"),
		host:       os.Getenv("APP_HOST"),
		port:       os.Getenv("APP_PORT"),
		versionapi: "/api/" + os.Getenv("APP_VERSIONAPI"),
	}

	envDBConfig := config.DatabaseConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBname:   os.Getenv("DATABASE_Name"),
	}

	Logger := log.New(os.Stdout, appConfig.name+" || ", log.LstdFlags)
	validate := validator.New()

	db, err := config.NewDatabaseConfig(envDBConfig)
	if err != nil {
		panic(fmt.Errorf("Error Connection : %s", err))
	}

	serveMux := http.NewServeMux()

	productRepository := repositories.NewProductRepositories(db, true)
	productUsecase := usecases.NewProductUsecase(&productRepository, validate)
	controllers.NewProductController(appConfig.versionapi, Logger, serveMux, &productUsecase)

	s := &http.Server{
		Addr:         ":" + appConfig.port,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second, // max time for connection using TCP keep-alive
		ReadTimeout:  5 * time.Second,   // max time to read request from client
		WriteTimeout: 5 * time.Second,   // max time to write response to client
	}

	go func() {
		Logger.Printf("Starting server on port :%s\n", appConfig.port)
		err := s.ListenAndServe()
		if err != nil {
			Logger.Fatal(err)
		}
	}()

	// get signal when trap signal interupt and grafefully shutdown server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until signal is received
	sig := <-sigChan
	Logger.Println("Received terminate, Graceful shutdown: ", sig)

	// waiting max 30 second for current operation complete
	tContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tContext)
}
