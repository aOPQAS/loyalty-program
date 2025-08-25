package pgsql

import (
	"database/sql"
	"testing"

	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestRunMigrations_MockConn(t *testing.T) {
	db, _ := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	conn := &dbr.Connection{DB: db}
	client := NewClient(conn)

	err := client.RunMigrations()
	assert.Error(t, err)
}
