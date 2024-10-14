package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"Email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// Validate the JWT token with the auth-service.
func (jc *JwtCustomClaims) Validate(hctx *HandlerContext, token *jwt.Token) (bool, error) {
	authUrl := fmt.Sprintf("%s/api/auth/verify/%s", hctx.GetCore().AuthServiceEndpoint, token.Raw)

	resp, err := http.Post(authUrl, "application/json", nil)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusAccepted {
		return false, errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, "Unauthorized"))
	}

	return true, nil
}
