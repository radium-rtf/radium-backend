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
	}

	ConferenceSubscribers struct {
		bun.BaseModel `bun:"table:wave.conference_subscriber"`

		ConferenceId uuid.UUID   `bun:",pk"`
		Conference   *Conference `bun:"rel:belongs-to,join:conference_id=id"`

		UserId uuid.UUID    `bun:",pk"`
		User   *radium.User `bun:"rel:belongs-to,join:user_id=id"`

		RoleCode string
		Role     *ConferenceRole `bun:"rel:belongs-to,join:role_code=code"`
	}

	ConferenceRole struct {
		bun.BaseModel `bun:"table:wave.conference_role"`
		DBModel

		Code string `bun:",unique"`
		Name string
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
