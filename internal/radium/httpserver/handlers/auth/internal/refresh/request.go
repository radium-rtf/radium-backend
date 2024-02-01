package refresh

import "github.com/google/uuid"

type Refresh struct {
	RefreshToken uuid.UUID `json:"refreshToken"`
}
