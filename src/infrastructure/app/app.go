package app

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-api/src/infrastructure/repository"
	"go-api/src/infrastructure/service"
	"go-api/src/pkg/database/postgres"
	"go-api/src/transport/handler"
	"log"

	"github.com/spf13/viper"
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

	if err = initConfig(); err != nil {
		log.Fatalf(err.Error())
	}

	if err = godotenv.Load(); err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("DB_PASSWORD: %v\n", os.Getenv("DB_PASSWORD"))

	db, err := postgres.NewPostgresDB(&postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)
	newHandler := handler.NewHandler(newService)

	server := new(Server)

	go func() {
		if err = server.Run(viper.GetString("port"), newHandler.InitRoutes()); err != nil {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server shutting down")

	err = server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
}
