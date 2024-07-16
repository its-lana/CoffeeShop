package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/its-lana/coffee-shop/config"
	"github.com/its-lana/coffee-shop/handlers"
	"github.com/its-lana/coffee-shop/logger"
	"github.com/its-lana/coffee-shop/repository"
	"github.com/its-lana/coffee-shop/server"
	"github.com/its-lana/coffee-shop/usecase"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logger.Log.Info("recovering from panic")
		}
	}()

	if err := godotenv.Load(); err != nil {
		logger.Log.Errorf("unable to load env")
	}

	accessLogFile, accessLogErr := os.OpenFile("access.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if accessLogErr != nil {
		log.Fatal("error creating access log file, " + accessLogErr.Error())
	}
	errorLogFile, errorLogErr := os.OpenFile("error.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if errorLogErr != nil {
		log.Fatal("error creating error log file, " + errorLogErr.Error())
	}
	customLog := logger.NewLogger()
	logger.SetLogger(customLog)
	db, err := config.NewPG(context.Background(), accessLogFile)
	if err != nil {
		log.Fatal("unable to connect to database")
	}

	db.MigratingDatabase()

	// construct route
	custRepo := repository.NewCustomerRepository(db)
	custUC := usecase.NewCustomerUseCase(custRepo)
	custH := handlers.NewCustomerHandler(custUC)

	authUC := usecase.NewAuthCustomerUsecase(custRepo)
	authH := handlers.NewAuthHandler(authUC)

	opts := server.RouterHandler{
		CustomerHandler: custH,
		AuthHandler:     authH,
	}

	r := server.NewRouter(opts, accessLogFile, errorLogFile)

	appPort := os.Getenv("APP_PORT")

	srv := http.Server{
		Addr:    appPort,
		Handler: r,
	}
	fmt.Println("running on port ", appPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Info("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	logger.Log.Info("Server exiting")
}
