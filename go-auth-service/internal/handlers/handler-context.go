package handlers

import (
	"github.com/labstack/echo"
	"github.com/sounishnath003/go-auth-service/internal/core"
)

// Declare custom context.
type HandlerContext struct {
	echo.Context
	Co *core.Core
}

func (hc *HandlerContext) GetCore() *core.Core {
	return hc.Co
}
