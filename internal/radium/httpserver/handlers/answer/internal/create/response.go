package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/answer/verdict"
)

type Response struct {
	Verdict verdict.Type `json:"verdict"`
	PageId  uuid.UUID    `json:"pageId"`
}
