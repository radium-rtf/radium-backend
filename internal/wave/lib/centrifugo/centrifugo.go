package centrifugo

import (
	"encoding/json"
	"log"

	"github.com/centrifugal/centrifuge-go"
	"github.com/golang-jwt/jwt"
	"github.com/radium-rtf/radium-backend/config"
)

type Centrifugo struct {
	clients map[string]*centrifuge.Client
	token   string
}

func (c Centrifugo) getConnectionToken(user string, exp int64) string {
	claims := jwt.MapClaims{"sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(c.token))
	if err != nil {
		panic(err)
	}
	return t
}

func (c Centrifugo) GetClient(user string, exp int64) *centrifuge.Client {
	client, exists := c.clients[user]
	if exists {
		return client
	}

	client = centrifuge.NewJsonClient(
		"ws://centrifugo:6969/connection/websocket",
		centrifuge.Config{
			Token: c.getConnectionToken(user, exp),
			// GetToken: func(e centrifuge.ConnectionTokenEvent) (string, error) {
			// 	return c.getConnectionToken("1", 0), nil
			// },
		},
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

	c.clients[user] = client

	return client
}

func (c Centrifugo) getSubscriptionToken(channel string, user string, exp int64) string {
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

func (c Centrifugo) GetSubscription(channel string, user string, exp int64) (*centrifuge.Subscription, error) {
	client, exists := c.clients[user]
	if !exists {
		panic("client not found")
	}

	sub, exists := client.GetSubscription(channel)
	if exists {
		return sub, nil
	}

	sub, err := client.NewSubscription(channel, centrifuge.SubscriptionConfig{
		GetToken: func(e centrifuge.SubscriptionTokenEvent) (string, error) {
			return c.getSubscriptionToken(e.Channel, user, exp), nil
		},
	})
	if err != nil {
		return nil, err
	}

	sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
		log.Printf("Subscribed to channel %s", sub.Channel)
	})
	sub.OnUnsubscribed(func(e centrifuge.UnsubscribedEvent) {
		log.Printf("Unsubscribed from channel %s", sub.Channel)
	})
	sub.OnError(func(e centrifuge.SubscriptionErrorEvent) {
		log.Printf("Subscription error: %s", e.Error)
	})
	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		log.Println("New publication in channel", sub.Channel, string(e.Data))
	})

	return sub, nil
}

func New(cfg config.Centrifugo) Centrifugo {
	return Centrifugo{
		token:   cfg.Token,
		clients: make(map[string]*centrifuge.Client),
	}
}
