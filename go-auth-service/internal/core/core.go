package core

import (
	"log"
	"log/slog"
	"os"
	"sync"
)

type Core struct {
	PORT int
	DSN  string
	Lo   *slog.Logger
	mu   sync.Mutex
}

func NewCore() *Core {
	return &Core{
		PORT: getEnv("PORT", 3000).(int),
		DSN:  getEnv("DSN", "postgres://127.0.0.1/go-auth-service?sslmode=disable").(string),
		Lo:   slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func getEnv(key string, fallback any) any {
	if val, ok := os.LookupEnv(key); ok {
		log.Printf("key %s environment value found\n", key)
		return val
	}
	log.Printf("no environment value found. setting fallback value key=%s\n", key)
	return fallback
}
