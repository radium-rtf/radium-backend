package answer

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChoice(t *testing.T) {
	defaultSection := &entity.ChoiceSection{
		Answer:   "4",
		Variants: []string{"3", "5", "не хочу отвечать на этот вопрос", "4"},
		Question: "2 + 2 = ?",
	}
	tests := []struct {
		name        string
		expected    verdict.Type
		expectedErr error
		answer      *entity.ChoiceSectionAnswer
		section     *entity.ChoiceSection
	}{
		{
			name:     "WA",
			expected: verdict.WA,
			section:  defaultSection,
			answer: &entity.ChoiceSectionAnswer{
				Answer: "2",
			},
		},
		{
			name:     "OK",
			expected: verdict.OK,
			section:  defaultSection,
			answer: &entity.ChoiceSectionAnswer{
				Answer: "4",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			checker := Checker{}
			section := &entity.Section{ChoiceSection: tt.section}
			answer := &entity.Answer{Choice: tt.answer}
			ver, err := checker.Check(section, answer)

			require.Equal(t, tt.expectedErr, err)
			require.Equal(t, tt.expected, ver.Verdict)
		})
	}
}
