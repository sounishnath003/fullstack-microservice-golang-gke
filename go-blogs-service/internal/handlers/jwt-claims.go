package handlers

import (
	"bytes"
	"encoding/json"
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
//
// Returns UserID (>0) and error (int, error). if error then ID = -1
func (jc *JwtCustomClaims) Validate(hctx *HandlerContext, token *jwt.Token) (int, error) {
	// If token is not valid.
	if !token.Valid {
		return -1, errors.New("invalid token")
	}
	// Creating the dynamic auth endpoint url for the user verification.
	authUrl := fmt.Sprintf("%s/api/auth/verify/%s", hctx.GetCore().AuthServiceEndpoint, token.Raw)
	// Payload for Request body.
	payload := map[string]any{
		"email":    jc.Email,
		"username": jc.Username,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return -1, err
	}
	// Grab the resp.
	resp, err := http.Post(authUrl, "application/json", bytes.NewReader(payloadBytes))
	// Throws err.
	if err != nil {
		return -1, err
	}
	// Check for the 202 response only.
	if resp.StatusCode != http.StatusAccepted {
		return -1, errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, "Unauthorized"))
	}

	// Parse the resp.Body into struct.
	var parsed VerifyUserResp
	json.NewDecoder(resp.Body).Decode(&parsed)
	// Log the output.
	hctx.GetCore().Lo.Info("user.verification", "parsed", parsed, "userid", parsed.Data.User.ID, "resp.statuscode", resp.StatusCode)
	// Return true or false.
	return parsed.Data.User.ID, nil
}

type VerifyUserResp struct {
	Data       Data  `json:"data"`
	StatusCode int64 `json:"statusCode"`
}

type Data struct {
	User  User `json:"user"`
	Valid bool `json:"valid"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
