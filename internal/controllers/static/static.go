package static

import (
	"github.com/gin-gonic/gin"
	"github.com/mateuszmierzwinski/gopress/internal/controllers"
	"net/http"
)

type staticController struct {
}

/*
Register function registers endpoints of this controller in Gin-gonic router
*/
func (s *staticController) Register(e *gin.Engine) {
	// handling static mapping
	e.Static("/web", "web")

	// root redirects to /web
	e.GET("/", s.rootRedirect)
}

/*
rootRedirect redirects root path "/" to "/web" subpath
*/
func (s *staticController) rootRedirect(context *gin.Context) {
	context.Redirect(http.StatusPermanentRedirect, "/web")
}

/*
New function initializes staticController struct and keeps it aligned with controllers.APIController interface
*/
func New() controllers.APIController {
	return &staticController{}
}
