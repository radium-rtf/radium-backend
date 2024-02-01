package answer

import (
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/answer/verdict"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChoice(t *testing.T) {
	defaultSection := &entity2.Section{
		Answer:   "4",
		Variants: []string{"3", "5", "не хочу отвечать на этот вопрос", "4"},
		Content:  "2 + 2 = ?",
		Type:     entity2.ChoiceType,
	}
	tests := []struct {
		name        string
		expected    verdict.Type
		expectedErr error
		answer      *entity2.Answer
		section     *entity2.Section
	}{
		{
			name:     "WA",
			expected: verdict.WA,
			section:  defaultSection,
			answer: &entity2.Answer{
				Type:   entity2.ChoiceType,
				Answer: "2",
			},
		},
		{
			name:     "OK",
			expected: verdict.OK,
			section:  defaultSection,
			answer: &entity2.Answer{
				Type:   entity2.ChoiceType,
				Answer: "4",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			checker := Checker{}
			ver, err := checker.Check(tt.section, tt.answer)

			require.Equal(t, tt.expectedErr, err)
			require.Equal(t, tt.expected, ver.Verdict)
		})
	}
}
