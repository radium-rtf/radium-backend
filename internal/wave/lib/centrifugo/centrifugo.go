package centrifugo

import (
	"encoding/json"
	"log"

	"github.com/centrifugal/centrifuge-go"
	"github.com/golang-jwt/jwt"
	"github.com/radium-rtf/radium-backend/config"
)

func getConnectionToken(user string, exp int64, token string) string {
	claims := jwt.MapClaims{"sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(token))
	if err != nil {
		panic(err)
	}
	return t
}

type Centrifugo struct {
	Client *centrifuge.Client
	token  string
}

func (c Centrifugo) GetSubscriptionToken(channel string, user string, exp int64) string {
	claims := jwt.MapClaims{"channel": channel, "sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(c.token))
	if err != nil {
		panic(err)
	}
	return t
}

func (c Centrifugo) Connect() error {
	return c.Client.Connect()
}

func (c Centrifugo) Close() {
	c.Client.Close()
}

func New(cfg config.Centrifugo) Centrifugo {
	config := centrifuge.Config{
		Token: getConnectionToken("1", 0, cfg.Token),
		// GetToken: func(e centrifuge.ConnectionTokenEvent) (string, error) {
		// 	return getConnectionToken("1", 0, cfg.Token), nil
		// },
	}
	client := centrifuge.NewJsonClient(
		"ws://centrifugo:6969/connection/websocket",
		config,
	)

	client.OnConnecting(func(e centrifuge.ConnectingEvent) {
		log.Printf("Connecting - %d (%s)", e.Code, e.Reason)
	})
	client.OnConnected(func(e centrifuge.ConnectedEvent) {
		log.Printf("Connected with ID %s", e.ClientID)
	})
	client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		log.Printf("Disconnected: %d (%s)", e.Code, e.Reason)
		go client.Connect() // keep alive on errors
	})
	client.OnPublication(func(e centrifuge.ServerPublicationEvent) {
		var msg map[string]interface{}
		_ = json.Unmarshal(e.Data, &msg)
		bt, _ := json.MarshalIndent(msg, "  ", " ")
		log.Printf("Publication: (%s)\n", string(bt))
	})

	return Centrifugo{
		Client: client,
		token:  cfg.Token,
	}
}
