package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashString(to_hash string) (string, error) {
	to_hash_byte_array := []byte(to_hash)
	hashed_byte_array, err := Hash(to_hash_byte_array)

	return string(hashed_byte_array), err
}

func Hash(to_hash []byte) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword(to_hash, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hashed, nil
}
