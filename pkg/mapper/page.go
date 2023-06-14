package mapper

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Page struct {
	section Section
}

func (p Page) Page(page *entity.Page, studentId string, verdicts map[string]map[uuid.UUID]entity.Verdict) entity.PageDto {
	return entity.PageDto{
		Id:       page.Id,
		Slug:     page.Slug,
		Name:     page.Name,
		Sections: p.section.Sections(page.Sections, studentId, verdicts),
	}
}
