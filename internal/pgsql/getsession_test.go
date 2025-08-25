package pgsql

import (
	"testing"

	"github.com/gocraft/dbr/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetSession(t *testing.T) {
	conn := &dbr.Connection{}
	client := NewClient(conn)
	session := client.GetSession()
	assert.NotNil(t, session)
}
