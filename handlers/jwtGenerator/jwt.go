package jwtGenerator

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"simple-oauth2/config"
	dtoRequest "simple-oauth2/dto/request"
	dtoResponse "simple-oauth2/dto/response"

	"github.com/golang-jwt/jwt/v4"
)

const (
	// Token expiration time
	tokenExpiration = 1800 // 30 mins

)

func Create(w http.ResponseWriter, req *http.Request) {
	var payload dtoRequest.Creds

	configs := config.Get()

	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Unable to read body!", http.StatusBadRequest)
		return
	}

	// Create a new JWT token
	// Generate the private key for signing

	// Parse the private key
	key, err := jwt.ParseRSAPrivateKeyFromPEM(configs.PrivateKey)
	if err != nil {
		http.Error(w, "Internal server error 2", http.StatusInternalServerError)
		return
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "client",
		"iss": "your-issuer",
		"exp": time.Now().Add(time.Second * time.Duration(tokenExpiration)).Unix(),
	})

	// Sign the token with the private key
	tokenString, err := token.SignedString(key)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Generate the token response
	response := dtoResponse.JWT{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   tokenExpiration,
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the response as JSON
	json.NewEncoder(w).Encode(response)
}
