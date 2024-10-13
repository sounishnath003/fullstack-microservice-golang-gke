package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// GetUser helps to get the user bare minimum information.
// ID is the path param /api/users/:ID parameters captured.
func GetUser(c echo.Context) error {
	// Get the core
	IDs := c.Param("ID")
	if len(IDs) == 0 {
		return ErrorApiResponse(c, http.StatusBadRequest, errors.New("bad request, required user id"))
	}

	uid, err := strconv.Atoi(IDs)
	if err != nil {
		return ErrorApiResponse(c, http.StatusBadRequest, err)
	}

	var user User
	hctx := c.(*HandlerContext)

	hctx.GetCore().QueryStmts.GetUserByID.QueryRow(uid).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password)

	return c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, map[string]any{
		"user": user,
		"uid":  uid,
	}, nil))
}

// LoginHandler helps to create the authentication and authorization of a user.
// Sends JWT api token with short-lived parameter to login and authorize the resources.
func LoginHandler(c echo.Context) error {
	var loginUser LoginUserDto
	// Throws error.
	if err := c.Bind(&loginUser); err != nil {
		return ErrorApiResponse(c, http.StatusInternalServerError, err)
	}
	// Accuquire the handler context.
	hctx := c.(*HandlerContext)

	// Set custom claims.
	token, err := GenerateNewJWTClaimToken(loginUser.Username, loginUser.Password, hctx)
	if err != nil {
		return ErrorApiResponse(c, http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, echo.Map{
		"token":    token,
		"username": loginUser.Username,
	}, nil))

}

// SignupHandler helps to create a new user in auth service
//
// You can later use the /api/auth/login API endpoint to login into the systems.
func SignupHandler(c echo.Context) error {
	var newUser SignupUserDto

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, NewApiResponse(http.StatusBadRequest, nil, err))
	}
	// Generate a Hash Password
	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newUser.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewApiResponse(http.StatusInternalServerError, nil, err))
	}
	// Get the core context.
	hctx := c.(*HandlerContext)

	// Create a new entry of the user.
	_, err = hctx.GetCore().QueryStmts.CreateNewUser.Exec(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Email, string(hashPassword))
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewApiResponse(http.StatusBadRequest, nil, err))
	}
	// Get User ID and other info.
	err = hctx.GetCore().QueryStmts.GetUserByUsername.QueryRow(newUser.Username).Scan(&newUser.ID, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.Username, &newUser.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewApiResponse(http.StatusBadRequest, nil, err))
	}

	// Add user entry to role.
	_, err = hctx.GetCore().QueryStmts.AddUserToUserRole.Exec(newUser.ID, 1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewApiResponse(http.StatusBadRequest, nil, err))
	}

	return c.JSON(http.StatusCreated, NewApiResponse(http.StatusCreated, "user registered successfully",
		nil))
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

func ErrorApiResponse(c echo.Context, status int, err error) error {
	return c.JSON(status, NewApiResponse(status, nil, err))
}
