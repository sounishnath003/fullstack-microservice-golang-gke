package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sounishnath003/go-auth-service/internal/core"
	"github.com/sounishnath003/go-auth-service/internal/server/handlers"
)

type Server struct {
	port int
}

func NewServer(co *core.Core) *Server {
	return &Server{
		port: co.PORT,
	}
}

func (s *Server) Start() error {
	e := echo.New()

	// Define middlewares.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.CSRF())
	e.Use(middleware.Gzip())

	// Define routes handlers
	e.GET("/", handlers.BasicHandler)

	e.Logger.Info("server has been started and running on", "port", s.port)
	return e.Start(fmt.Sprintf(":%d", s.port))
}
