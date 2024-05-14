package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/centrifugal/centrifuge-go"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/lib/centrifugo"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type MessageUseCase struct {
	message    postgres2.Message
	centrifugo centrifugo.Centrifugo
}

func (uc MessageUseCase) GetMessage(ctx context.Context) (*entity.Message, error) {
	message, err := uc.message.Get(ctx)
	return message, err
}

func (uc MessageUseCase) GetMessagesFrom(ctx context.Context, chatId uuid.UUID) ([]*entity.Message, error) {
	message, err := uc.message.Get(ctx)
	return []*entity.Message{message}, err
}

func (uc MessageUseCase) SendMessage(ctx context.Context, chatId uuid.UUID, content model.Content) (*model.Message, error) {
	var err error
	client := uc.centrifugo.Client
	sub, exists := client.GetSubscription(chatId.String())
	if !exists {
		sub, err = client.NewSubscription(chatId.String(), centrifuge.SubscriptionConfig{
			GetToken: func(e centrifuge.SubscriptionTokenEvent) (string, error) {
				token := uc.centrifugo.GetSubscriptionToken(e.Channel, "1", time.Now().Unix()+10)
				return token, nil
			},
		})
		if err != nil {
			return nil, err
		}
	}

	sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
		fmt.Println("Subscribed to channel", sub.Channel)
	})
	sub.OnError(func(e centrifuge.SubscriptionErrorEvent) {
		fmt.Println("Subscription error:", sub.Channel, e.Error)
	})
	sub.OnUnsubscribed(func(e centrifuge.UnsubscribedEvent) {
		fmt.Println("Unsubscribed from channel", sub.Channel, e.Reason)
	})
	sub.OnPublication(func(e centrifuge.PublicationEvent) {
		fmt.Println("New publication in channel", sub.Channel, string(e.Data))
	})
	defer sub.Unsubscribe()
	sub.Subscribe()

	json_data := []byte(`{"value":"` + content.Text + `"}`)
	_, err = sub.Publish(ctx, json_data)
	if err != nil {
		return nil, err
	}
	message := model.Message{
		ChatId:  chatId,
		Content: content,
	}
	return &message, nil
}

func NewMessageUseCase(messageRepo postgres2.Message, centrifugo centrifugo.Centrifugo) MessageUseCase {
	return MessageUseCase{message: messageRepo, centrifugo: centrifugo}
}
