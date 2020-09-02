package main

import (
	"encoding/base64"
	"fmt"
	"testproject/src/hashing"
)

func main() {

	password := "VerySecurPassword"

	pbkdf2Salt, pbkdf2Hash := hashing.Pbkdf2Hash(password)
	fmt.Println("$pbkdf2-sha512$i=100000,l=64$" + base64.RawStdEncoding.EncodeToString(pbkdf2Salt) + "$" + base64.RawStdEncoding.EncodeToString(pbkdf2Hash))

	hashBcrypt := hashing.BcryptHash(password)
	fmt.Println("$2b$10$" + base64.StdEncoding.EncodeToString(hashBcrypt))

	saltArgon2, hashArgon2 := hashing.Argon2Hash(password)
	fmt.Println("$argon2id$v=19$m=65536,t=2,p=1$" + base64.RawStdEncoding.EncodeToString(saltArgon2) + "$" + base64.RawStdEncoding.EncodeToString(hashArgon2))

	saltHmac, hashHmac := hashing.HmacHash(password, "sldkjf")
	fmt.Println(base64.StdEncoding.EncodeToString(saltHmac))
	fmt.Println(hashHmac)
}
