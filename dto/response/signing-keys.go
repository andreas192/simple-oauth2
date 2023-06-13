package dto

type SigningKeys struct {
	Kty string `json:"kty"`
	E   int    `json:"e"`
	N   string `json:"n"`
}
