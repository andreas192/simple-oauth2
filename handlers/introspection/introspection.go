package introspection

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-oauth2/config"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Do(w http.ResponseWriter, req *http.Request) {

	configs := config.Get()

	// Check authorization header for access token
	authorizationHeader := req.Header.Get("Authorization")
	if authorizationHeader == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessToken := extractAccessToken(authorizationHeader)
	if accessToken == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Perform token introspection
	introspectionResponse, err := introspectToken(accessToken, configs.PublicKey)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response body
	responseBytes, err := json.Marshal(introspectionResponse)
	if err != nil {
		log.Println("Failed to marshal introspection response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(responseBytes)
	if err != nil {
		log.Println("Failed to write response:", err)
	}
}

// Extracts the access token from the authorization header.
func extractAccessToken(authorizationHeader string) string {
	parts := strings.SplitN(authorizationHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}
	return parts[1]
}

// Introspects the access token.
func introspectToken(accessToken string, publicKey []byte) (tokenDecoded json.Token, err error) {
	standardClaims := new(jwt.StandardClaims)
	tokenDecoded, err = jwt.ParseWithClaims(accessToken, standardClaims, func(token *jwt.Token) (pvk interface{}, err error) {
		return jwt.ParseRSAPublicKeyFromPEM(publicKey)
	})
	if err != nil {
		log.Println("Failed to decode token:", err)
		return
	}

	return
}
