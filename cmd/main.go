package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alvinmatias69/shortener/internal/controller"
	"github.com/alvinmatias69/shortener/internal/handler"
	"github.com/alvinmatias69/shortener/internal/repository"
	"github.com/alvinmatias69/shortener/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var (
		connectionString = os.Getenv("DATABASE_URL")
		port             = os.Getenv("PORT")
	)

	pgpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Printf("error while creating db connection pool: %v\n", err)
		os.Exit(1)
	}

	repositoryInstance := repository.New(pgpool)
	controllerInstance := controller.New(repositoryInstance)
	handlerInstance := handler.New(controllerInstance)
	shortenerServer := server.New(handlerInstance)

	fmt.Printf("Starting shortener server in port: %v\n", port)
	shortenerServer.Start(port)
}
