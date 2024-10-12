package handlers

import (
	"errors"
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/sounishnath003/go-auth-service/internal/utils"
)

var (
	users map[int]*UserSignupDto = map[int]*UserSignupDto{}
	lock  sync.Mutex             = sync.Mutex{}
)

func GetUser(c echo.Context) error {
	userId := c.Param("id")
	if len(userId) == 0 {
		return c.JSON(http.StatusBadRequest, utils.NewApiResponse(http.StatusBadRequest, nil, errors.New("user id is required")))
	}

	return c.JSON(http.StatusNotFound,
		utils.NewApiResponse(http.StatusNotFound, nil, errors.New("user not found")),
	)
}

func CreateUser(c echo.Context) error {

	return c.JSON(http.StatusCreated, nil)
}
