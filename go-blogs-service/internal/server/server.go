package server

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sounishnath003/go-blogs-service/internal/core"
	"github.com/sounishnath003/go-blogs-service/internal/handlers"
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
	e := echo.New() // Use echo/v4 instead of echo

	// Declare the custom context in the route handler
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.HandlerContext{
				Context: c,
				Co:      s.co,
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
	e.Use(middleware.Logger())

	// Define routes handlers
	e.GET("/", handlers.BasicHandler)
	e.GET("ping", handlers.PingHandler)

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handlers.JwtCustomClaims)
		},
		SigningKey: []byte(s.co.JWTSecret),
	}
	// Group all handler with /api
	api := e.Group("/api")
	api.Use(echojwt.WithConfig(config))
	api.GET("/blogs/recommendations", handlers.BlogsRecommendationHandler)
	api.POST("/blogs/create", handlers.CreateNewBlogpostHandler)

	e.Logger.Info("server has been started and running on", "port", s.co.PORT)
	return e.Start(fmt.Sprintf(":%d", s.co.PORT))
}
