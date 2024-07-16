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
	AuthHandler     *handlers.AuthHandler
	MerchantHandler *handlers.MerchantHandler
}

func NewRouter(opts RouterHandler, accessLogFile, errorLogFile *os.File) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(middleware.CorsMiddleware())

	v1 := router.Group("/api/v1")
	v1Auth := router.Group("/v1/auth", middleware.AuthorizeHandler())
	v1.GET("/", defaultHandler)

	const login = "/login"
	v1.POST(login+"/customer", opts.AuthHandler.LoginCustomer)
	v1.POST(login+"/merchant", opts.AuthHandler.LoginMerchant)

	const customer = "/customer"
	v1.POST(customer+"/register", opts.CustomerHandler.RegisterCustomer)
	v1Auth.GET(customer, opts.CustomerHandler.GetAllCustomers)

	const merchant = "/merchant"
	v1.POST(merchant+"/register", opts.MerchantHandler.RegisterMerchant)
	v1Auth.GET(merchant, opts.MerchantHandler.GetAllMerchants)

	gin.DefaultWriter = io.MultiWriter(os.Stdout, accessLogFile)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, errorLogFile)

	return router
}

func defaultHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Default API Coffee Shop"})
}
