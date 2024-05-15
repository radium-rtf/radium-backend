package model

type (
	CentrifugoToken struct {
		Token string `json:"token"`
	}
)

func NewCentrifugoToken(token string) CentrifugoToken {
	return CentrifugoToken{
		Token: token,
	}
}
