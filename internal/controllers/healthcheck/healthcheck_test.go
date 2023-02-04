package healthcheck

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mateuszmierzwinski/gopress/internal/controllers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckRegisterMethod(t *testing.T) {
	// with (initiation)
	hc := New()
	writer := httptest.NewRecorder()
	_, r := gin.CreateTestContext(writer)
	hasRoute := false

	// do (execution)
	hc.Register(r)

	server := httptest.NewServer(r)
	defer server.Close()

	client := server.Client()
	result, err := client.Get(fmt.Sprintf("%s/health", server.URL))

	// then (validation)
	assert.Implements(t, (*controllers.APIController)(nil), hc, "healthCheck controller needs to implement APIController interface")

	routes := r.Routes()
	for _, route := range routes {
		if route.Method == http.MethodGet &&
			route.Path == "/health" {
			hasRoute = true
		}
	}

	assert.True(t, hasRoute, "Controller should register /health path")
	assert.NoErrorf(t, err, "HTTP Client returned error %x while calling endpoint /health", err)
	assert.Equalf(t, http.StatusOK, result.StatusCode, "HTTP Client returned status %d but expected 200 (OK) from /health endpoint")
}
