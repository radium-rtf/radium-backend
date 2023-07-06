package mapper

import (
	"github.com/dranikpg/dto-mapper"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/otp"
)

type Group struct {
	otp.OTPGenerator
}

func (g Group) PostToGroup(post entity.GroupPost) *entity.Group {
	courses := make([]*entity.Course, 0, len(post.StudentsIds))
	for _, id := range post.StudentsIds {
		courses = append(courses, &entity.Course{DBModel: entity.DBModel{Id: id}})
	}

	students := make([]*entity.User, 0, len(post.StudentsIds))
	for _, id := range post.StudentsIds {
		students = append(students, &entity.User{DBModel: entity.DBModel{Id: id}})
	}

	return &entity.Group{
		DBModel:    entity.DBModel{Id: uuid.New()},
		Name:       post.Name,
		InviteCode: g.RandomSecret(10),
		Students:   students,
		Courses:    courses,
	}
}

func (g Group) ToDto(group *entity.Group) *entity.GroupDto {
	groupDto := entity.GroupDto{}
	dto.Map(&groupDto, group)
	return &groupDto
}

func (g Group) ToGroupsDto(groups []*entity.Group) []*entity.GroupDto {
	groupsDto := make([]*entity.GroupDto, 0, len(groups))
	for _, group := range groups {
		groupsDto = append(groupsDto, g.ToDto(group))
	}
	return groupsDto
}
