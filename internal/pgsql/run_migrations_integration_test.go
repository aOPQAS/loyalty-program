package pgsql

import (
	"fmt"
	"microservice/config"
	"testing"
)

func TestRunMigrations_Integration(t *testing.T) {
	t.Skip("Integration test: requires real Postgres and migrations folder")

	cfg := config.PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		Database: "testdb",
		SSLMode:  "disable",
	}

	conn, err := CreatePostgresConnection(cfg)
	if err != nil {
		t.Fatal(err)
	}

	client := NewClient(conn)
	if err := client.RunMigrations(); err != nil {
		t.Fatal(err)
	}
	fmt.Println("Migrations applied successfully")
}
