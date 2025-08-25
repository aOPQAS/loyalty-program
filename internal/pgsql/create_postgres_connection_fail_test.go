package pgsql

import (
	"testing"

	"microservice/config"

	"github.com/stretchr/testify/assert"
)

func TestCreatePostgresConnection_Fail(t *testing.T) {
	cfg := config.PostgresConfig{
		Host:     "localhost",
		Port:     "9999",
		User:     "wrong",
		Password: "wrong",
		Database: "wrong",
		SSLMode:  "disable",
	}

	conn, err := CreatePostgresConnection(cfg)
	assert.Nil(t, conn)
	assert.Error(t, err)
}
