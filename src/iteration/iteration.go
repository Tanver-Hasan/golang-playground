package iteration

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

// Repeat - Repeate characters
func Repeat(character string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

