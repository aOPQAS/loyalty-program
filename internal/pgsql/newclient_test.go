package pgsql

import (
	"testing"

	"github.com/gocraft/dbr/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	conn := &dbr.Connection{}
	client := NewClient(conn)
	assert.NotNil(t, client)
	assert.Equal(t, conn, client.conn)
}
