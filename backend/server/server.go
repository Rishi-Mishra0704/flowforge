package server

import (
	"net/http"

	"github.com/Rishi-Mishra0704/flowforge/backend/config"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	router *echo.Echo
	config config.Config
}

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewServer(config config.Config) (*Server, error) {	server := &Server{
		config: config,	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := echo.New()

	router.Use(middleware.CORS())

	router.Validator = &Validator{validator: validator.New()}

	// Middleware
	router.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	router.POST("/flowchart", server.GetFlowChartHandler)
	router.POST("/login", server.Login)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}
