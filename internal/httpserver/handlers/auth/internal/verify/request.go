package verify

type Request struct {
	VerificationCode string `json:"verificationCode"`
	Email            string `json:"email" validate:"email"`
}
