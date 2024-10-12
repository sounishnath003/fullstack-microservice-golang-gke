package core

import (
	"log/slog"
	"os"
	"sync"

	"github.com/sounishnath003/go-auth-service/internal/utils"
)

type Core struct {
	PORT int
	DSN  string
	Lo   *slog.Logger
	mu   sync.Mutex
}

func NewCore() *Core {
	return &Core{
		PORT: utils.GetEnv("PORT", 3000).(int),
		DSN:  utils.GetEnv("DSN", "postgres://root:password@127.0.0.1:5432/auth?sslmode=disable").(string),
		Lo:   slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}
