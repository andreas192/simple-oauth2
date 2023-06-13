package signingKeys

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"log"
	"net/http"
	"simple-oauth2/config"
	dtoResponse "simple-oauth2/dto/response"
)

func List(w http.ResponseWriter, req *http.Request) {

	configs := config.Get()

	// Decode the PEM-encoded key
	block, _ := pem.Decode(configs.PublicKey)
	if block == nil {
		log.Printf("Failed to decode PEM block")
		http.Error(w, "Unable to decode PEM!", http.StatusInternalServerError)
		return
	}

	// Parse the DER-encoded public key
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Printf("Failed to parse public key: %v", err)
		http.Error(w, "Failed to parse public key!", http.StatusInternalServerError)
		return
	}

	jwkResponse := dtoResponse.SigningKeys{
		Kty: "RSA",
		E:   pubKey.(*rsa.PublicKey).E,
		N:   pubKey.(*rsa.PublicKey).N.String(),
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the response as JSON
	json.NewEncoder(w).Encode(jwkResponse)
}
