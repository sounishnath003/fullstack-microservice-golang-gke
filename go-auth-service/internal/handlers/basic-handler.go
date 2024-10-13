package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/sounishnath003/go-auth-service/internal/utils"
)

func BasicHandler(c echo.Context) error {
	hostname, _ := os.Hostname()

	dataResp := map[string]any{
		"version":      utils.GetEnv("K_VERSION", "v0.0.1").(string),
		"serviceName":  "go-auth-service",
		"releasedDate": time.Now().Format(time.DateOnly),
		"hostname":     hostname,
		"apiUrls": []string{
			"/api/auth/signup",
			"/api/auth/login",
			"/api/auth/password-reset",
		},
	}

	return c.JSON(http.StatusOK, dataResp)
}

func PingHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, NewApiResponse(
		http.StatusOK,
		map[string]string{"message": "pong", "timestamp": time.Now().Format(time.RFC3339), "version": utils.GetEnv("K_VERSION", "v0.0.1").(string),
			"serviceName": "go-auth-service"},
		nil,
	))
}
