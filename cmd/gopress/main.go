package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mateuszmierzwinski/gopress/internal/controllers"
	"github.com/mateuszmierzwinski/gopress/internal/controllers/healthcheck"
	"os"
)

const defaultBindAddress = "0.0.0.0:8080"

/*
getBindAddress returns binding address from ENV variables in case of those to be set (Docker).
Otherwise, defaultBindAddress const value is used and binds all interfaces with port 8080.
*/
func getBindAddress(provider func(string) string) string {
	addr := provider("BIND_IP")
	port := provider("PORT")

	if addr != "" && port != "" {
		return fmt.Sprintf("%s:%s", addr, port)
	}

	return defaultBindAddress
}

/*
initRouter initializes Gin-Gonic router and registers controllers
*/
func initRouter(apiControllers []controllers.APIController) (router *gin.Engine) {
	router = gin.New()
	for _, apiController := range apiControllers {
		apiController.Register(router)
	}
	return
}

func main() {
	// Add your controllers here
	apiControllers := []controllers.APIController{
		healthcheck.New(),
	}

	// General routing handling
	initRouter(apiControllers).Run(getBindAddress(os.Getenv))
}
