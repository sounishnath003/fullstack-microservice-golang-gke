package core

import (
	"database/sql"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/knadh/goyesql/v2"
	_ "github.com/lib/pq"

	"github.com/sounishnath003/go-blogs-service/internal/utils"
)

type Core struct {
	PORT                int
	DSN                 string
	AuthServiceEndpoint string
	QueryStmts          *BlogsServiceQueries
	Lo                  *slog.Logger
	mu                  sync.Mutex
	JWTSecret           string
}

func NewCore() *Core {
	// Define the logger.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	dsn := utils.GetEnv("DSN", "postgres://root:password@127.0.0.1:5432/blogs?sslmode=disable").(string)
	driver := utils.GetEnv("DRIVER", "postgres").(string)

	// Check the db open.
	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	// Set the connections defaults.
	db.SetMaxIdleConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(100 * time.Second)

	// Check for ping.
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Parse prebuilt SQL queries using goyesql.
	queries := goyesql.MustParseFile("queries.sql")
	var queryStmts BlogsServiceQueries
	// Prepares a given set of Queries and assigns the resulting *sql.Stmt statements to the fields of a given struct.
	err = goyesql.ScanToStruct(&queryStmts, queries, db)
	if err != nil {
		panic(err)
	}

	// Return the core.
	return &Core{
		PORT:                utils.GetEnv("PORT", 3001).(int),
		DSN:                 dsn,
		JWTSecret:           utils.GetEnv("JWT_SECRET", "my5u43Rs3CR3T0k3N$!(1).*").(string),
		AuthServiceEndpoint: utils.GetEnv("AUTH_SERVICE_ENDPOINT", "http://127.0.0.1:3000").(string),
		QueryStmts:          &queryStmts,
		Lo:                  logger,
	}
}
