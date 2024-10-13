package core

import "database/sql"

type AuthServiceQueries struct {
	GetUserByID       *sql.Stmt `query:"getUserByID"`
	CreateNewUser     *sql.Stmt `query:"createNewUser"`
	GetUserByUsername *sql.Stmt `query:"getUserByUsername"`
	GetUserForJWT     *sql.Stmt `query:"getUserForJWT"`
	AddUserToUserRole *sql.Stmt `query:"addUserToUserRole"`
}
