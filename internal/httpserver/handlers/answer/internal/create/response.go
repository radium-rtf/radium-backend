package create

import "github.com/google/uuid"

type Response struct {
	Verdict string    `json:"verdict"`
	PageId  uuid.UUID `json:"pageId"`
}
