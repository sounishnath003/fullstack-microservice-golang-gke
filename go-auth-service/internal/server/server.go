package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sounishnath003/go-auth-service/internal/core"
	"github.com/sounishnath003/go-auth-service/internal/handlers"
)

type Server struct {
	co *core.Core
}

func NewServer(co *core.Core) *Server {
	return &Server{
		co: co,
	}
}

func (s *Server) Start() error {
	e := echo.New()

	// Declare the custom context in the route handler
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.HandlerContext{
				Context: c, Co: s.co,
			}
			return next(cc)
		}
	})
	// Define other middlewares.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002", "http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	// Define routes handlers
	e.GET("/", handlers.BasicHandler)
	e.GET("ping", handlers.PingHandler)

	// Group all handler with /api
	api := e.Group("/api")
	api.GET("/auth/users/:ID", handlers.GetUser)
	api.POST("/auth/login", handlers.LoginHandler)
	api.POST("/auth/signup", handlers.SignupHandler)

	e.Logger.Info("server has been started and running on", "port", s.co.PORT)
	return e.Start(fmt.Sprintf(":%d", s.co.PORT))
}
