package answer

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChoice(t *testing.T) {
	defaultSection := &entity.Section{
		Answer:   "4",
		Variants: []string{"3", "5", "не хочу отвечать на этот вопрос", "4"},
		Content:  "2 + 2 = ?",
	}
	tests := []struct {
		name        string
		expected    verdict.Type
		expectedErr error
		answer      *entity.Answer
		section     *entity.Section
	}{
		{
			name:     "WA",
			expected: verdict.WA,
			section:  defaultSection,
			answer: &entity.Answer{
				Answer: "2",
			},
		},
		{
			name:     "OK",
			expected: verdict.OK,
			section:  defaultSection,
			answer: &entity.Answer{
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
