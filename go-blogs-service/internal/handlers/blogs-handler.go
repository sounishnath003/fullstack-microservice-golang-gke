package handlers

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Blogs struct {
	ID        string `json:"id"`
	UserID    string `json:"userID"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func BlogsRecommendationHandler(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	// If token is not valid.
	if !token.Valid {
		return ErrorApiResponse(c, http.StatusBadRequest, errors.New("invalid token"))
	}
	// Claims the token.
	claims := token.Claims.(*JwtCustomClaims)
	hctx := c.(*HandlerContext)
	if _, err := claims.Validate(hctx, token); err != nil {
		return ErrorApiResponse(c, http.StatusUnauthorized, err)
	}

	var blogs []Blogs

	resultRows, err := hctx.GetCore().QueryStmts.GetLatestRecommendedBlogs.Query()
	if err != nil {
		return ErrorApiResponse(c, http.StatusInternalServerError, err)
	}

	for resultRows.Next() {
		var blog Blogs
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
	if _, err := claims.Validate(hctx, token); err != nil {
		return ErrorApiResponse(c, http.StatusUnauthorized, err)
	}

	var newBlog CreateBlogDto
	// Throws error.
	if err := c.Bind(&newBlog); err != nil {
		return ErrorApiResponse(c, http.StatusBadRequest, err)
	}

	_, err := hctx.GetCore().QueryStmts.CreateNewBlogpost.Exec(1, newBlog.Title, newBlog.Subtitle, newBlog.Content)
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
