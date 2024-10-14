package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sounishnath003/go-blogs-service/internal/core"
)

// Declare custom context.
type HandlerContext struct {
	echo.Context
	Co *core.Core
}

func (hc *HandlerContext) GetCore() *core.Core {
	return hc.Co
}
