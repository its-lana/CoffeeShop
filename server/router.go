package server

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/its-lana/coffee-shop/handlers"
	"github.com/its-lana/coffee-shop/middleware"
)

type RouterHandler struct {
	CustomerHandler *handlers.CustomerHandler
}

func NewRouter(opts RouterHandler, accessLogFile, errorLogFile *os.File) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(middleware.CorsMiddleware())

	v1 := router.Group("/api/v1")

	v1.GET("/", defaultHandler)

	gin.DefaultWriter = io.MultiWriter(os.Stdout, accessLogFile)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, errorLogFile)

	return router
}

func defaultHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Default API Coffee Shop"})
}
