package centrifugo

import (
	"encoding/json"
	"log/slog"
	"strconv"

	"github.com/centrifugal/centrifuge-go"
	"github.com/golang-jwt/jwt/v5"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/auth"
)

type Centrifugo struct {
	tokenManager auth.TokenManager
	clients      map[string]*centrifuge.Client
}

func (c Centrifugo) GetConnectionToken(user string, exp int64) (string, error) {
	claims := jwt.MapClaims{"sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	return c.tokenManager.NewCustomJWT(claims)
}

func (c Centrifugo) GetClient(user string, exp int64) *centrifuge.Client {
	client, exists := c.clients[user]
	if exists {
		return client
	}

	token, _ := c.GetConnectionToken(user, exp)

	client = centrifuge.NewJsonClient(
		"ws://centrifugo:6969/connection/websocket",
		centrifuge.Config{
			Token: token,
			// GetToken: func(e centrifuge.ConnectionTokenEvent) (string, error) {
			// 	return c.getConnectionToken("1", 0), nil
			// },
		},
	)

	client.OnConnecting(func(e centrifuge.ConnectingEvent) {
		slog.Info("Connecting", "code", strconv.Itoa(int(e.Code)), "reason", e.Reason)
	})
	client.OnConnected(func(e centrifuge.ConnectedEvent) {
		slog.Info("Connected", "clientId", e.ClientID)
	})
	client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		slog.Info("Disconnected", "code", strconv.Itoa(int(e.Code)), "reason", e.Reason)
		go client.Connect() // keep alive on errors
	})
	client.OnPublication(func(e centrifuge.ServerPublicationEvent) {
		var msg map[string]interface{}
		_ = json.Unmarshal(e.Data, &msg)
		bt, _ := json.MarshalIndent(msg, "  ", " ")
		slog.Info("Publication", "data", string(bt))
	})

	c.clients[user] = client

	client.Connect()

	return client
}

func (c Centrifugo) GetSubscriptionToken(channel string, user string, exp int64) (string, error) {
	claims := jwt.MapClaims{"channel": channel, "sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	return c.tokenManager.NewCustomJWT(claims)
}

func (c Centrifugo) GetSubscription(channel string, user string, exp int64) (*centrifuge.Subscription, error) {
	client, exists := c.clients[user]
	if !exists {
		return nil, centrifuge.Error{Code: 1, Message: "client not found"}
	}

	sub, exists := client.GetSubscription(channel)
	// возможно саб не имеет обработки ивентов ниже
	if exists {
		return sub, nil
	}

	token, err := c.GetSubscriptionToken(channel, user, exp)
	if err != nil {
		return nil, err
	}

	sub, err = client.NewSubscription(channel, centrifuge.SubscriptionConfig{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
		slog.Info("Subscribed to channel", "channel", sub.Channel)
	})
	sub.OnUnsubscribed(func(e centrifuge.UnsubscribedEvent) {
		slog.Info("Unsubscribed from channel", "channel", sub.Channel)
	})
	sub.OnError(func(e centrifuge.SubscriptionErrorEvent) {
		slog.Error("Subscription error in channel", "channel", sub.Channel, "error", e.Error)
	})
	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		slog.Info("New publication in channel", "channel", sub.Channel, "data", string(e.Data))
	})

	return sub, nil
}

func New(tokenManager auth.TokenManager) Centrifugo {
	return Centrifugo{
		tokenManager: tokenManager,
		clients:      make(map[string]*centrifuge.Client),
	}
}
