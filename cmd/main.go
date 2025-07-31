package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"vk-server-task/internal/app"
	"vk-server-task/internal/handler"
	"vk-server-task/internal/service"
	"vk-server-task/internal/storage"
	"vk-server-task/pkg/postgres"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	rootCtx := context.Background()
	db, err := postgres.NewPostgresDB(rootCtx, postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("SSLMODE"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	storage := storage.New(db)
	service := service.New(storage)
	handler := handler.New(service)

	srv := app.NewServer(os.Getenv("HTTP_PORT"), handler.InitRoutes())
	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Println("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("App shutting down")
	if err := srv.Shutdown(rootCtx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
