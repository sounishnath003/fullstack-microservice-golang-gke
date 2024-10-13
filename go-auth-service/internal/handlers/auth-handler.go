package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo"
)

var (
	users map[int]*UserSignupDto = map[int]*UserSignupDto{}
	lock  sync.Mutex             = sync.Mutex{}
)

func GetUser(c echo.Context) error {
	// Get the core
	IDs := c.Param("ID")
	if len(IDs) == 0 {
		return c.JSON(http.StatusBadRequest, NewApiResponse(http.StatusBadRequest, nil, errors.New("bad request, required user id")))
	}

	uid, err := strconv.Atoi(IDs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewApiResponse(
			http.StatusBadRequest, nil, err,
		))
	}

	var user User
	hctx := c.(*HandlerContext)
	hctx.GetCore().QueryStmts.GetUserByID.QueryRow(uid).Scan(&user.Username, &user.Password)

	return c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, map[string]any{
		"user": user,
		"uid":  uid,
	}, nil))

}

func CreateUser(c echo.Context) error {

	return c.JSON(http.StatusCreated, nil)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password,-"`
}
