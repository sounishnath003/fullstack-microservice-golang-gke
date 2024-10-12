package core

import (
	"database/sql"
	"log/slog"
	"os"
	"sync"

	"github.com/sounishnath003/go-auth-service/internal/utils"
)

type Core struct {
	PORT int
	DSN  string
	DB   *sql.DB
	Lo   *slog.Logger
	mu   sync.Mutex
}

func NewCore() *Core {
	dsn := utils.GetEnv("DSN", "postgres://root:password@127.0.0.1:5432/auth?sslmode=disable").(string)
	driver := utils.GetEnv("DRIVER", "postgres").(string)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	return &Core{
		PORT: utils.GetEnv("PORT", 3000).(int),
		DSN:  dsn,
		DB:   db,
		Lo:   slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}
