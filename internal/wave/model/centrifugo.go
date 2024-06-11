package model

import "github.com/google/uuid"

type (
	CentrifugoToken struct {
		Token string `json:"token"`
	}

	CentrifugoEvent struct {
		Event    string    `json:"event"`
		ChatId   uuid.UUID `json:"chat_id"`
		ChatType string    `json:"chat_type,omitempty"`
		Message  *Message  `json:"message,omitempty"`
		Dialogue *Dialogue `json:"dialogue,omitempty"`
		Group    *Group    `json:"group,omitempty"`
	}
)

func NewCentrifugoToken(token string) CentrifugoToken {
	return CentrifugoToken{
		Token: token,
	}
}
