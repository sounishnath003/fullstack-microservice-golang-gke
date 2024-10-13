package core

import (
	"database/sql"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/knadh/goyesql/v2"
	_ "github.com/lib/pq"

	"github.com/sounishnath003/go-auth-service/internal/utils"
)

type Core struct {
	PORT       int
	DSN        string
	DB         *sql.DB
	QueryStmts *AuthServiceQueries
	Lo         *slog.Logger
	mu         sync.Mutex
}

func NewCore() *Core {
	dsn := utils.GetEnv("DSN", "postgres://root:password@127.0.0.1:5432/auth?sslmode=disable").(string)
	driver := utils.GetEnv("DRIVER", "postgres").(string)

	// Check the db open.
	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxIdleTime(100 * time.Second)

	// Check for ping.
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Parse prebuilt SQL queries using goyesql.
	queries := goyesql.MustParseFile("queries.sql")
	var queryStmts AuthServiceQueries
	goyesql.ScanToStruct(&queryStmts, queries, db)

	// Return the core.
	return &Core{
		PORT:       utils.GetEnv("PORT", 3000).(int),
		DSN:        dsn,
		DB:         db,
		QueryStmts: &queryStmts,
		Lo:         slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}
