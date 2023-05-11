package test

import (
	"go-di/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
