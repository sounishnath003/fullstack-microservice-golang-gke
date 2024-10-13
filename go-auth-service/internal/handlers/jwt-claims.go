package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JwtClaims struct {
	Username string `json:"username"`
	Email    string `json:"Email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type JwtUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

// GenerateNewJWTClaimToken helps to generate a new JWT token for the User
func GenerateNewJWTClaimToken(username, password string, hctx *HandlerContext) (string, error) {
	var user JwtUser

	// Scan the objects attribs.
	// Check if the user. present in DB.
	err := hctx.GetCore().QueryStmts.GetUserForJWT.QueryRow(username).Scan(&user.FirstName, &user.LastName, &user.UserName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return "", err
	}

	// Check if the password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	claims := JwtClaims{
		Username: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(hctx.GetCore().JWTSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}
