package model

type (
	CentrifugoToken struct {
		Token string `json:"token"`
	}

	CentrifugoEvent struct {
		Event   string   `json:"event"`
		Message *Message `json:"message"`
	}
)

func NewCentrifugoToken(token string) CentrifugoToken {
	return CentrifugoToken{
		Token: token,
	}
}

func NewCentrifugoEvent(event string, message *Message) CentrifugoEvent {
	return CentrifugoEvent{
		Event:   event,
		Message: message,
	}
}
