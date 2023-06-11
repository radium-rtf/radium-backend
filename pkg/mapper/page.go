package mapper

import "github.com/radium-rtf/radium-backend/internal/entity"

type Page struct {
	section Section
}

func (p Page) Page(page *entity.Page) entity.PageDto {
	return entity.PageDto{
		Id:       page.Id,
		Slug:     page.Slug,
		Name:     page.Name,
		Sections: p.section.Sections(page.Sections),
	}
}
