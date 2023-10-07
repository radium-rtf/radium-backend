package hash

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Hasher interface {
	Hash(password string) (string, error)
	Equals(hashed string, password string) bool
}

type PasswordHasher struct {
	sha1 *SHA1Hasher
	cost int
}

func NewPasswordHasher(sha1salt string, cost int) Hasher {
	return &PasswordHasher{sha1: NewSHA1Hasher(sha1salt), cost: cost}
}

func (h *PasswordHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), err
}

func (h *PasswordHasher) Equals(hashed string, password string) bool {
	if strings.Contains(hashed, "$") {
		eq := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
		return eq == nil
	}
	hash, err := h.sha1.Hash(password)
	if err != nil {
		return false
	}
	return hashed == hash
}
