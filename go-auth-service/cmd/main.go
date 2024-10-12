package main

import (
	"github.com/sounishnath003/go-auth-service/internal/core"
	"github.com/sounishnath003/go-auth-service/internal/server"
)

func main() {
	co := core.NewCore()
	co.Lo.Info("print.core.struct", "co", co.DSN)

	server := server.NewServer(co)
	err := server.Start()
	if err != nil {
		co.Lo.Error("server could not able to start", "err", err)
	}
}
