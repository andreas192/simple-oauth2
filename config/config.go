package config

import (
	"io/ioutil"
	"log"
)

type Config struct {
	PrivateKey []byte
	PublicKey  []byte
}

var (
	conf *Config
)

func Init() {
	conf = &Config{
		PrivateKey: loadPrivateKeyPEM(),
		PublicKey:  loadPublicKeyPEM(),
	}

}

func Get() *Config {
	return conf
}

func loadFile(filepath string) []byte {
	// Read the file
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func loadPrivateKeyPEM() []byte {
	privateKey := loadFile("private-key.pem")

	return privateKey
}

func loadPublicKeyPEM() []byte {
	publicKey := loadFile("public-key.pem")

	return publicKey
}
