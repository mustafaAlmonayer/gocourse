package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"slices"
)

func cryptoEx() {
	password := "password123"
	hash, salt, err := hashPass([]byte(password))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", salt)
	fmt.Printf("%x\n", hash)
	fmt.Println(verifyHash(salt, []byte(password), hash))
}

func hashPass(pass []byte) ([]byte, []byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, nil, err
	}
	saltedPassword := append(salt, pass...)
	hash := sha256.Sum256(saltedPassword)
	return hash[:], salt, nil
}

func verifyHash(salt []byte, pass []byte, hash []byte) bool {
	hashResult := sha256.Sum256(append(salt, pass...))
	if slices.Equal(hash, hashResult[:]) {
		return true
	}
	return false
}
