package main

import (
	"net/http"

	"simple-oauth2/config"
	"simple-oauth2/introspection"
	"simple-oauth2/jwtGenerator"
	signingKeys "simple-oauth2/signing-keys"
)

func main() {

	config.Init()

	http.HandleFunc("/token", jwtGenerator.Create)
	http.HandleFunc("/signing-keys", signingKeys.List)
	http.HandleFunc("/introspection", introspection.Do)

	http.ListenAndServe(":8090", nil)
}
