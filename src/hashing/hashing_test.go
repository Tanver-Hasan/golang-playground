package hashing

import (
	"crypto/sha512"
	"encoding/base64"
	"testing"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

func TestPbkdf2Hash(t *testing.T) {
	salt, hash := Pbkdf2Hash("VerySecurePassword")
	expectedHash := pbkdf2.Key([]byte("VerySecurePassword"), salt, 100000, 64, sha512.New)

	t.Run("Check hash", func(t *testing.T) {
		if base64.RawStdEncoding.EncodeToString((hash)) != base64.RawStdEncoding.EncodeToString(expectedHash) {
			t.Errorf(" Expected %q but got %q", expectedHash, base64.RawStdEncoding.EncodeToString(hash))
		}
	})

}

func TestBcryptHash(t *testing.T) {
	t.Run("Compare Bcrypt hash", func(t *testing.T) {
		hash := BcryptHash("VerySecurePassword")
		err := bcrypt.CompareHashAndPassword(hash, []byte("VerySecurePassword"))

		if err != nil {
			t.Errorf("Failed")
		}
	})
}

func TestArgonHash(t *testing.T) {
	saltArgon2, hashArgon2 := Argon2Hash("VerySecurePassword")
	expectedHashArgon2 := argon2.Key([]byte("VerySecurePassword"), saltArgon2, 3, 32*1024, 4, 64)

	t.Run("Compare argon2 hash", func(t *testing.T) {

		if base64.RawStdEncoding.EncodeToString((hashArgon2)) != base64.RawStdEncoding.EncodeToString(expectedHashArgon2) {
			t.Errorf(" Expected %q but got %q", expectedHashArgon2, base64.RawStdEncoding.EncodeToString(hashArgon2))
		}
	})
}

func BenchmarkPbkdf2Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pbkdf2Hash("VerySecurePassword")
	}
}

func BenchmarkBcryptHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BcryptHash("VerySecurePassword")
	}
}

func BenchmarkArgon2Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Argon2Hash("VerySecurePassword")
	}
}
