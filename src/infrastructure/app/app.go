package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-api/src/infrastructure/repository"
	"go-api/src/infrastructure/service"
	"go-api/src/pkg/database/postgres"
	"go-api/src/transport/handler"
	"log"

	"github.com/spf13/viper"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func initConfig() error {
	viper.AddConfigPath("src/configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func Run() {
	err := os.Chdir("./")
	if err != nil {
		log.Fatal(err)
	}

	logsFolder := "logs"
	_, err = os.Stat(logsFolder)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(logsFolder, 0755)
		if errDir != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}
	}

	logFile, err := os.OpenFile("logs/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	if err := initConfig(); err != nil {
		log.Fatalf("Caught error while initializing config: %v", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Caught error while loading .env file: %v", err.Error())
	}
	log.Println("DB_PASSWORD: %v", os.Getenv("DB_PASSWORD"))

	db, err := postgres.NewPostgresDB(&postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("Caught error while creating database: %v", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error while running http server: %v", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server shutting down")

	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("Error while shutting down http server: %v", err.Error())
	}
}
