package main

import (
	"github.com/sounishnath003/go-auth-service/internal/core"
)

func main() {
	co := core.NewCore()
	co.Lo.Info("co", co)
}
