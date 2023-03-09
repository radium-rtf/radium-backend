package otp

import "github.com/xlzd/gotp"

type Generator interface {
	RandomSecret(length int) string
}

type OTPGenerator struct{}

func NewOTPGenerator() *OTPGenerator {
	return &OTPGenerator{}
}

func (g *OTPGenerator) RandomSecret(length int) string {
	return gotp.RandomSecret(length)
}
