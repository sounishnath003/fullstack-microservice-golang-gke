package main

import (
	"github.com/sounishnath003/go-blogs-service/internal/core"
	"github.com/sounishnath003/go-blogs-service/internal/server"
)

func main() {
	co := core.NewCore()
	co.Lo.Info("core.struct", "co", co)

	server := server.NewServer(co)
	err := server.Start()
	if err != nil {
		co.Lo.Error("server could not able to start", "err", err)
	}
}
