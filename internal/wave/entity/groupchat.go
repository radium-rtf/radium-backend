package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	GroupChat struct {
		bun.BaseModel `bun:"table:groups"`
		DBModel

		Name           string
		Members        []*GroupChatMember        `bun:"rel:has-many"`
		Admins         []*GroupChatMember        `bun:"rel:has-many"`
		Settings       GroupChatSettings         `bun:"rel:has-one"`
		PinnedMessages []*GroupChatPinnedMessage `bun:"rel:has-many"`
		Messages       []*GroupChatMessage       `bun:"rel:has-many"`
	}

	GroupChatMember struct {
		bun.BaseModel `bun:"table:group_members"`

		GroupChatId uuid.UUID  `bun:",pk"`
		GroupChat   *GroupChat `bun:"rel:belongs-to,join:group_id=id"`
		UserId      uuid.UUID  `bun:",pk"`
		Role        string
	}

	GroupChatSettings struct {
		bun.BaseModel `bun:"table:group_settings"`
		DBModel
	}

	GroupChatMessage struct {
		bun.BaseModel `bun:"table:group_messages"`

		GroupChatId uuid.UUID  `bun:",pk"`
		GroupChat   *GroupChat `bun:"rel:belongs-to,join:group_id=id"`
		MessageId   uuid.UUID  `bun:",pk"`
		Message     *Message   `bun:"rel:belongs-to,join:message_id=id"`
	}

	GroupChatPinnedMessage struct {
		bun.BaseModel `bun:"table:group_pinned_messages"`

		GroupChatId uuid.UUID  `bun:",pk"`
		GroupChat   *GroupChat `bun:"rel:belongs-to,join:group_id=id"`
		MessageId   uuid.UUID  `bun:",pk"`
		Message     *Message   `bun:"rel:belongs-to,join:message_id=id"`
	}
)
