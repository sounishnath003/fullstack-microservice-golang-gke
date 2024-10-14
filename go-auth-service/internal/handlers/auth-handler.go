package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// GetUserByID helps to get the user bare minimum information.
// ID is the path param /api/users/:ID parameters captured.
func GetUserByID(c echo.Context) error {
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

	return c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, echo.Map{
		"user": user,
		"uid":  uid,
	}, nil))
}

// GetUserByUsername helps to get the user information by the username
func GetUserByUsername(c echo.Context) error {
	// Get username from path param.
	username := c.Param("Username")
	// Check for real username.
	if len(username) < 4 {
		return ErrorApiResponse(c, http.StatusBadRequest, errors.New("bad request, required username"))
	}
	// Save user info.
	var user User
	// Accuquire the context.
	hctx := c.(*HandlerContext)

	err := hctx.GetCore().QueryStmts.GetUserByUsername.QueryRow(username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password)
	// Throws err.
	if err != nil {
		return ErrorApiResponse(c, http.StatusBadRequest, errors.New("Username or blogs found."))
	}

	return c.JSON(http.StatusOK, NewApiResponse(http.StatusOK, echo.Map{
		"username": username,
		"user":     user,
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

// Helps to verify the JWT token and User existence else throws err.
// Reads the JWT token from the Param
func VerifyJwtTokenHandler(c echo.Context) error {
	jwtToken := c.Param("JwtToken")
	if len(jwtToken) < 30 {
		return ErrorApiResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
	}

	var verifyUser VerifyUserDto
	if err := c.Bind(&verifyUser); err != nil {
		return ErrorApiResponse(c, http.StatusBadRequest, err)
	}

	if len(verifyUser.Username) == 0 || len(verifyUser.Email) == 0 {
		return ErrorApiResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
	}

	var user User
	hctx := c.(*HandlerContext)

	hctx.GetCore().QueryStmts.GetUserByUsername.QueryRow(verifyUser.Username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password)

	// TODO: Proper JWT Verification and User existence checks for context security
	return c.JSON(http.StatusAccepted, NewApiResponse(http.StatusAccepted, echo.Map{
		"valid": true,
		"user":  user,
	}, nil))
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}
