package main

import {
	"crypto/rand"
	"encoding/base64"
	"log"

	"bcrypt"
}

func hashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err:= bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func generateSessionToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}