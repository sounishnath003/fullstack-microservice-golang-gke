package core

import "database/sql"

type AuthServiceQueries struct {
	GetUserByID   *sql.Stmt `query:"getUserByID"`
	CreateNewUser *sql.Stmt `query:"createNewUser"`
}
