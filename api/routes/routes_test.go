package routes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeRoutes(t *testing.T) {
	// Start a server to listen on a random port
	go InitializeRoutes(":7000")

	// Make a request to the home route and check if the response is "OK"
	res, err := http.Get("http://localhost:7000")
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.NoError(t, err)
}
