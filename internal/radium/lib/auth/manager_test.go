package auth

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func getAllRoleVariants() []*model.Roles {
	var variants []*model.Roles
	variants = append(variants, nil)

	var roles [3]bool
	for i := 0; i < (1 << 3); i++ {
		roles[0] = i%2 != 0
		roles[1] = (i>>1)%2 != 0
		roles[2] = (i>>2)%2 != 0

		role := &model.Roles{
			IsTeacher:  roles[0],
			IsCoauthor: roles[1],
			IsAuthor:   roles[2],
		}

		variants = append(variants, role)
	}

	return variants
}

type TestCase struct {
	token []string
	isErr bool
}

func newTestCases(token string) []TestCase {
	return []TestCase{
		{[]string{token}, true},
		{[]string{"Bearer1", token}, true},
		{[]string{"Bearer", token}, false},
		{[]string{"Bearer", ""}, true},
	}
}

func TestManager(t *testing.T) {
	manager := NewManager("arawdawdwea")
	userId := uuid.New()
	user := &model.User{
		Id:     userId,
		Email:  "radium@radium-rtf.ru",
		Avatar: "",
	}

	for _, role := range getAllRoleVariants() {
		user.Roles = role

		token, err := manager.NewJWT(user, time.Now().Add(time.Hour*24))
		require.Nil(t, err)

		testCases := newTestCases(token)

		for _, tt := range testCases {
			if user.Roles == nil {
				user.Roles = new(model.Roles)
			}

			isAuthor, err := manager.ExtractIsAuthor(tt.token)
			require.Equal(t, tt.isErr, err != nil)
			if err == nil {
				require.Equal(t, user.Roles.IsAuthor, isAuthor)
			}

			isCoauthor, err := manager.ExtractIsCoauthor(tt.token)
			require.IsType(t, tt.isErr, err != nil)
			if err == nil {
				require.Equal(t, user.Roles.IsCoauthor, isCoauthor)
			}

			isTeacher, err := manager.ExtractIsTeacher(tt.token)
			require.IsType(t, tt.isErr, err != nil)
			if err == nil {
				require.Equal(t, user.Roles.IsTeacher, isTeacher)
			}

			id, err := manager.ExtractUserId(tt.token)
			require.IsType(t, tt.isErr, err != nil)
			if err == nil {
				require.Equal(t, user.Id, id)
			}

			canEditCourse, err := manager.ExtractCanEditCourse(tt.token)
			require.IsType(t, tt.isErr, err != nil)

			// TODO: должно быть в геттере, убрать все остальные user.Roles.IsCoauthor || user.Roles.IsAuthor
			if err == nil {
				require.Equal(t, user.Roles.IsCoauthor || user.Roles.IsAuthor, canEditCourse)
			}
		}
	}
}
