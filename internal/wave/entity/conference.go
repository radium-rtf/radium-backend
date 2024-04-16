package entity

import (
	"github.com/google/uuid"
	radium "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/uptrace/bun"
)

type (
	Conference struct {
		bun.BaseModel `bun:"table:wave.conferences"`
		DBModel
		Name string

		Comments    ConferenceCommentSection `bun:"rel:has-one"`
		Settings    ConferenceSettings       `bun:"rel:has-one"`
		Subscribers []*radium.User           `bun:"m2m:wave.conference_subscriber,join:Conference=User"`
		OwnerId     uuid.UUID                `bun:",notnull"`
		Owner       *radium.User             `bun:"rel:belongs-to,join:owner_id=id"`
	}

	ConferenceSettings struct {
		bun.BaseModel `bun:"table:wave.conference_settings"`
		DBModel

		ConferenceId uuid.UUID `bun:",pk"`

		Roles []*ConferenceRole `bun:"rel:has-many"`
	}

	ConferenceSubscribers struct {
		bun.BaseModel `bun:"table:wave.conference_subscriber"`
		DBModel

		ConferenceId uuid.UUID   `bun:",pk"`
		Conference   *Conference `bun:"rel:belongs-to,join:conference_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`

		Roles []*ConferenceRole `bun:"m2m:wave.conference_subscriber_role,join:Subscriber=Role"`
	}

	ConferenceSubscriberRoles struct {
		bun.BaseModel `bun:"table:wave.conference_subscriber_role"`

		SubscriberId uuid.UUID              `bun:",pk"`
		Subscriber   *ConferenceSubscribers `bun:"rel:belongs-to,join=subscriber_id=id"`

		RoleId uuid.UUID       `bun:",pk"`
		Role   *ConferenceRole `bun:"rel:belongs-to,join:role_id=id"`
	}

	ConferenceRole struct {
		bun.BaseModel `bun:"table:wave.conference_role"`
		DBModel

		Name string `bun:",notnull,unique"`

		Permissions ConferenceRolePermissions `bun:"rel:has-one"`
	}

	ConferenceRolePermissions struct {
		bun.BaseModel `bun:"table:wave.group_role_permissions"`

		ConferenceRoleId uuid.UUID `bun:",pk"`

		// Absolute
		Administrator bool
		// Moderation
		DeleteMessages bool
		AddUsers       bool
		KickUsers      bool
		MuteUsers      bool
		// General
		Write           bool
		CreateThreads   bool
		AttachFiles     bool
		MentionEveryone bool
		StartCalls      bool
		Speak           bool
		EnableCamera    bool
		ScreenShare     bool
	}

	ConferenceCommentSection struct {
		bun.BaseModel `bun:"table:wave.conference_comment_section"`

		ConferenceId uuid.UUID   `bun:",pk"`
		Conference   *Conference `bun:"rel:belongs-to,join:conference_id=id"`

		// Settings ConferenceCommentsSettings `bun:"rel:has-one"`

		Comments []*Message `bun:"m2m:wave.conference_comment,join:Conference=Message"`
	}

	ConferenceComments struct {
		bun.BaseModel `bun:"table:wave.conference_comment"`

		ConferenceId uuid.UUID   `bun:",pk"`
		Conference   *Conference `bun:"rel:belongs-to,join:conference_id=id"`

		MessageId uuid.UUID `bun:",pk"`
		Message   *Message  `bun:"rel:belongs-to,join:message_id=id"`
	}
)
