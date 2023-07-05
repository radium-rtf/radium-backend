package mapper

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Course struct {
	m Module
}

func (c Course) ToDto(course *entity.Course) *entity.CourseDto {
	links := make([]entity.LinkDto, 0, len(course.Links))
	for _, link := range course.Links {
		links = append(links, entity.LinkDto{Name: link.Name, Link: link.Link})
	}

	authors := make([]entity.UserDto, 0, len(course.Authors))
	for _, author := range course.Authors {
		authors = append(authors,
			entity.UserDto{Id: author.Id, Name: author.Name, Avatar: author.Avatar, Email: author.Email})
	}

	return &entity.CourseDto{
		Id:               course.Id,
		Name:             course.Name,
		Slug:             course.Slug,
		ShortDescription: course.ShortDescription,
		Description:      course.Description,
		Logo:             course.Logo,
		Banner:           course.Banner,
		Authors:          authors,
		Links:            links,
		Modules:          c.m.ModulesToDto(course.Modules),
	}
}
