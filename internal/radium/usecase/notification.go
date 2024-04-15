package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
)

type NotificationUseCase struct {
	notification postgres.Notification
}

func NewNotificationUseCase(notification postgres.Notification) NotificationUseCase {
	return NotificationUseCase{notification: notification}
}

func (uc NotificationUseCase) Get(ctx context.Context, userId uuid.UUID) ([]entity.Notification, error) {
	return uc.notification.Get(ctx, userId)
}
