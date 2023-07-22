package signin

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
