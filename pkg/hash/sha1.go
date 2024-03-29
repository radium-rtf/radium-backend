package hash

import (
	"crypto/sha1"
	"fmt"
)

type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{salt: salt}
}

func (h *SHA1Hasher) Hash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}

func (h *SHA1Hasher) Equals(hashed string, password string) bool {
	hash, err := h.Hash(password)
	return hashed == hash && err == nil
}
