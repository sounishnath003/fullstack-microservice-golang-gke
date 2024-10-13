package handlers

import (
	"net/http"
	"sync"

	"github.com/labstack/echo"
)

var (
	users map[int]*UserSignupDto = map[int]*UserSignupDto{}
	lock  sync.Mutex             = sync.Mutex{}
)

func GetAllUsers(c echo.Context) error {
	cc := c.(*HandlerContext)
	results, err := cc.Co.DB.Query("SELECT Username, Password FROM users ORDER BY CreatedAt DESC LIMIT 10;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewApiResponse(http.StatusInternalServerError, nil, err))
	}
	// Store sql.rows into variables and return as response
	var users []UserSignupDto
	// iterate over the array
	for results.Next() {
		var user UserSignupDto
		err := results.Scan(&user.Username, &user.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, NewApiResponse(http.StatusInternalServerError, nil, err))
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, users, nil))
}

func GetUser(c echo.Echo) error {
	return nil
}

func CreateUser(c echo.Context) error {

	return c.JSON(http.StatusCreated, nil)
}
