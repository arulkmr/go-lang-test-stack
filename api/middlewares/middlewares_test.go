package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetMiddlewareJSON(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	middleware := SetMiddlewareJSON(handler)
	middleware.ServeHTTP(res, req)

	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	assert.Equal(t, "Hello World", res.Body.String())
}
