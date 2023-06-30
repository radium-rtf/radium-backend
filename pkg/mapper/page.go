package mapper

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Page struct {
	section Section
}

func (p Page) Page(page *entity.Page, answers map[uuid.UUID]*entity.Answer) *entity.PageDto {
	sectionsDto := p.section.Sections(page.Sections, answers)
	return &entity.PageDto{
		Id:       page.Id,
		Slug:     page.Slug,
		Name:     page.Name,
		Sections: sectionsDto,
	}
}
