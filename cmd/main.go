package main

import (
	"fmt"
	"os"

	"microservice/config"
	"microservice/internal/deps"
	"microservice/internal/pgsql"
	"microservice/internal/server"
	"microservice/pkg/log"
	"microservice/pkg/telebon"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	log.Info("setting default timezone to UTC")
	if err := os.Setenv("TZ", "UTC"); err != nil {
		log.Fatal("failed to set UTC timezone", zap.Error(err))
	}

	log.Info("loading .env file")
	if err := godotenv.Load(); err != nil {
		log.Error("failed to load .env file", zap.Error(err))
	} else {
		fmt.Println(".env loaded successfully")
	}

	fmt.Println("ACCESS_TOKEN =", os.Getenv("ACCESS_TOKEN"))

	log.Info("loading config")
	cfg := config.NewConfig()
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("failed to parse .env to config")
	}

	log.SetLogEncoding(cfg.Logger.Encoding)
	log.SetLogLevel(cfg.Logger.Level)

	log.Info("creating pgsql connection")
	conn, err := pgsql.CreatePostgresConnection(cfg.Postgres)
	if err != nil {
		log.Fatal("failed to make pg connection", zap.Error(err))
	}

	log.Info("loading pgsql client")
	pg := pgsql.NewClient(conn)

	log.Info("Initializing telebon client")
	client := telebon.New(os.Getenv("ACCESS_TOKEN"))
	data, err := client.GetSubproducts()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data.([]byte)))

	log.Info("starting server")
	s := server.New(&deps.Deps{
		PG:      pg,
		Telebon: client,
	})

	if err := s.App.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}
