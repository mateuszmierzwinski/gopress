package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/mateuszmierzwinski/gopress/internal/controllers"
	"net/http"
)

type healthCheck struct {
}

/*
Register method is used to bind URLs to methods within healthCheck structure
*/
func (h *healthCheck) Register(e *gin.Engine) {
	e.GET("/health", h.getHealthEndpoint)
}

/*
getHealthEndpoint defines handling of healthcheck endpoint on code level
*/
func (h *healthCheck) getHealthEndpoint(context *gin.Context) {
	context.String(http.StatusOK, "text/plain", "Ok")
}

/*
New function guarantees alignment of this internal module with APIController interface
*/
func New() controllers.APIController {
	return &healthCheck{}
}
