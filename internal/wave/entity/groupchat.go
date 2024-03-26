package entity

import (
	"github.com/google/uuid"
	radium "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

type (
	GroupChat struct {
		bun.BaseModel `bun:"table:groups"`
		DBModel

		Name           string
		Members        []*radium.User    `bun:"m2m:group_chat_members,join:GroupChat=User"`
		Admins         []*radium.User    `bun:"m2m:group_chat_admins,join:GroupChat=User"`
		Settings       GroupChatSettings `bun:"rel:has-one"`
		PinnedMessages []*Message        `bun:"m2m:group_chat_pinned_messages,join:GroupChat=Message"`
		Messages       []*Message        `bun:"m2m:group_chat_messages,join:GroupChat=Message"`
	}

	GroupChatMember struct {
		bun.BaseModel `bun:"table:group_chat_members"`

		GroupChatId uuid.UUID    `bun:",pk"`
		GroupChat   *GroupChat   `bun:"rel:belongs-to,join:group_chat_id=id"`
		UserId      uuid.UUID    `bun:",pk"`
		User        *radium.User `bun:"rel:belongs-to,join:user_id=id"`
		Role        string
	}

	GroupChatAdmin struct {
		bun.BaseModel `bun:"table:group_chat_admins"`

		GroupChatId uuid.UUID    `bun:",pk"`
		GroupChat   *GroupChat   `bun:"rel:belongs-to,join:group_chat_id=id"`
		UserId      uuid.UUID    `bun:",pk"`
		User        *radium.User `bun:"rel:belongs-to,join:user_id=id"`
	}

	GroupChatSettings struct {
		bun.BaseModel `bun:"table:group_chat_settings"`
		DBModel
	}

	GroupChatMessage struct {
		bun.BaseModel `bun:"table:group_chat_messages"`

		GroupChatId uuid.UUID  `bun:",pk"`
		GroupChat   *GroupChat `bun:"rel:belongs-to,join:group_chat_id=id"`
		MessageId   uuid.UUID  `bun:",pk"`
		Message     *Message   `bun:"rel:belongs-to,join:message_id=id"`
	}

	GroupChatPinnedMessage struct {
		bun.BaseModel `bun:"table:group_chat_pinned_messages"`

		GroupChatId uuid.UUID  `bun:",pk"`
		GroupChat   *GroupChat `bun:"rel:belongs-to,join:group_chat_id=id"`
		MessageId   uuid.UUID  `bun:",pk"`
		Message     *Message   `bun:"rel:belongs-to,join:message_id=id"`
	}
)
