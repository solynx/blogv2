package app

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)
const rsaPrivateKeyLocation = "private-key.pem"

type ServerConfig struct {
	App  *fiber.App
	Port string
	RouterPrivateKey *rsa.PrivateKey
}

func (server *ServerConfig) Start() {
	server.App.Listen(server.Port)
}

func (server *ServerConfig) Setup() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	server.App = fiber.New()
	server.Port = ":" + port

	//load public key
	// pubKey, err := os.ReadFile(rsaPrivateKeyLocation)
	// if err != nil {
	// 	log.Fatal("File location not found")
	// }
	// pemPubKey, _ := pem.Decode(pubKey)
	// pub, err := x509.ParsePKCS1PublicKey(pemPubKey.Bytes)
	// if err != nil {
	// 	log.Fatal("Parse private key failed")
	// }
	// server.RouterPublicKey = pub

	//load private key
	pemData, err := os.ReadFile(rsaPrivateKeyLocation)
	if err != nil {
		log.Fatal("File location not found")
	}
  
	pemBlock, _ := pem.Decode(pemData)
	priv, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		log.Fatal("Parse private key failed", err)
	}
	server.RouterPrivateKey = priv.(*rsa.PrivateKey)


}
