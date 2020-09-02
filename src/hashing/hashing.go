package hashing

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

// Pbkdf2Hash : pbkdf2 password hash
func Pbkdf2Hash(password string) ([]byte, []byte) {

	salt := GenerateSalt()
	hash := pbkdf2.Key([]byte(password), salt, 100000, 64, sha512.New)

	return salt, hash
}

// BcryptHash :  Bcrypt password hash
func BcryptHash(password string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	return hash
}

// Argon2Hash : argon2 algorithm
func Argon2Hash(password string) ([]byte, []byte) {
	salt := GenerateSalt()

	hash := argon2.Key([]byte(password), salt, 3, 32*1024, 4, 64)
	return salt, hash
}

// HmacHash - Hmac algorithm
func HmacHash(password string, secret string) ([]byte, string) {

	salt := GenerateSalt()
	hash := hmac.New(sha512.New, []byte(secret))
	hash.Write([]byte(password))
	hashString := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return salt, hashString
}

// GenerateSalt - GenerateSalt
func GenerateSalt() []byte {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		fmt.Println("Error", err)
	}
	return salt
}
