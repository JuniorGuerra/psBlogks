package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(name string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(name), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store: ", string(hash))
	return base64.StdEncoding.EncodeToString(hash)
}
