package responses

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{"message": "Hello, World!"}
	statusCode := http.StatusOK

	// Test 1
	JSON(w, statusCode, data)
	assert.Equal(t, statusCode, w.Code)

	var result map[string]string
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, data, result)

	// Test 2
	w = httptest.NewRecorder()
	data = nil
	statusCode = http.StatusBadRequest

	JSON(w, statusCode, data)
	assert.Equal(t, statusCode, w.Code)
	assert.NotEqual(t, "", w.Body.String())
}

func TestERROR(t *testing.T) {
	// create a request to pass to the unknown handler
	_, err := http.NewRequest("GET", "/unknown", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	ERROR(rr, http.StatusBadRequest, err)

	// check the status code of the response
	assert.Equal(t, http.StatusBadRequest, rr.Code)

}
