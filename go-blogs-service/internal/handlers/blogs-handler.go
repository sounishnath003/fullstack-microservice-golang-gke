package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Blog struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userID"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func BlogsRecommendationHandler(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	// Claims the token.
	claims := token.Claims.(*JwtCustomClaims)
	hctx := c.(*HandlerContext)
	if _, err := claims.Validate(hctx, token); err != nil {
		return ErrorApiResponse(c, http.StatusUnauthorized, err)
	}

	var blogs []Blog

	resultRows, err := hctx.GetCore().QueryStmts.GetLatestRecommendedBlogs.Query()
	if err != nil {
		return ErrorApiResponse(c, http.StatusInternalServerError, errors.New("Unauthorized"))
	}

	for resultRows.Next() {
		var blog Blog
		err := resultRows.Scan(
			&blog.ID,
			&blog.UserID,
			&blog.Title,
			&blog.Subtitle,
			&blog.Content,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		)
		if err != nil {
			return ErrorApiResponse(c, http.StatusInternalServerError, err)
		}
		blogs = append(blogs, blog)
	}

	if len(blogs) == 0 {
		return ErrorApiResponse(c, http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, NewApiResponse(
		http.StatusOK,
		blogs,
		nil,
	))
}

// CreateNewBlogpostHandler helps to create a blog post on behalf of the user
func CreateNewBlogpostHandler(c echo.Context) error {
	// Accuquire context.
	hctx := c.(*HandlerContext)
	// grab the user
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JwtCustomClaims)
	userID, err := claims.Validate(hctx, token)
	if err != nil {
		return ErrorApiResponse(c, http.StatusUnauthorized, err)
	}

	var newBlog CreateBlogDto
	// Throws error.
	if err := c.Bind(&newBlog); err != nil {
		return ErrorApiResponse(c, http.StatusBadRequest, err)
	}

	_, err = hctx.GetCore().QueryStmts.CreateNewBlogpost.Exec(userID, newBlog.Title, newBlog.Subtitle, newBlog.Content)
	if err != nil {
		return ErrorApiResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, NewApiResponse(
		http.StatusAccepted,
		echo.Map{
			"message": "blog has been created",
		},
		nil,
	))
}

func GetBlogsByUsernameHandler(c echo.Context) error {
	// Accuquire context.
	hctx := c.(*HandlerContext)
	// Grab the user
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JwtCustomClaims)

	_, err := claims.Validate(hctx, token)
	if err != nil {
		return ErrorApiResponse(c, http.StatusUnauthorized, err)
	}

	// Get the username.
	username := c.Param("Username")
	// Throws err
	if len(username) == 0 {
		return ErrorApiResponse(c, http.StatusBadRequest, errors.New("Username not found"))
	}
	user, err := GetUserInfoByUsername(hctx, username)
	if err != nil {
		return ErrorApiResponse(c, http.StatusBadRequest, err)
	}

	resultRows, err := hctx.GetCore().QueryStmts.GetBlogsByUsername.Query(user.ID)
	if err != nil {
		return ErrorApiResponse(c, http.StatusBadRequest, err)
	}

	var blogs []Blog

	for resultRows.Next() {
		var blog Blog
		err := resultRows.Scan(
			&blog.ID,
			&blog.UserID,
			&blog.Title,
			&blog.Subtitle,
			&blog.Content,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		)
		if err != nil {
			return ErrorApiResponse(c, http.StatusInternalServerError, err)
		}
		blogs = append(blogs, blog)
	}

	if len(blogs) == 0 {
		return ErrorApiResponse(c, http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, NewApiResponse(
		http.StatusOK,
		echo.Map{
			"username": username,
			"userID":   user.ID,
			"blogs":    blogs,
		},
		nil,
	))
}

// GetUserInfoByUsername helps to get the user info from the auth service.
// This is used to get the user info from the auth service.
func GetUserInfoByUsername(hctx *HandlerContext, username string) (User, error) {
	authUrl := fmt.Sprintf("%s/api/auth/users/%s", hctx.GetCore().AuthServiceEndpoint, username)
	resp, err := http.Get(authUrl)
	if err != nil {
		return User{}, err
	}

	// Throws error if the statuscode != 200 (OK)
	if resp.StatusCode == http.StatusBadRequest {
		return User{}, errors.New("Invalid username")
	}
	
	var userInfo VerifyUserResp
	json.NewDecoder(resp.Body).Decode(&userInfo)
	if userInfo.Data.User.ID == 0 {
		return User{}, errors.New("Unauthorized")
	}

	hctx.GetCore().Lo.Info("userinfo", "userinfo.resp", userInfo)

	return userInfo.Data.User, nil
}
