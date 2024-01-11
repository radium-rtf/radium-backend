package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrl(t *testing.T) {
	type url struct {
		Url string `validate:"url"`
	}

	tests := []struct {
		url  string
		name string
		out  error
	}{
		{"http://localhost/", "http://localhost/", nil},
		{"http://localhost", "http://localhost", nil},
		{"", "not required", nil},
		{"gskpfeds", "invalid", ValidationError{}},
		{"https://localhost", "https://localhost", nil},
		{"https://localhost", "https://localhost/", nil},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := Struct(url{tt.url})

			assert.IsType(t, tt.out, err)
		})
	}
}

func TestPassword(t *testing.T) {
	type password struct {
		Password string `validate:"password"`
	}

	tests := []struct {
		password string
		name     string
		isErr    bool
	}{
		{"garad", "min length", true},
		{"1abcde", "min length", false},
		{"ewqdawara", "digits is required", true},
		{"141241243", "non-digits is required", true},
		{"dwadwda3141", "valid", false},
		{"AFASDRAA3141", "valid", false},
		{"4ХАХАХИХА1", "valid", false},
		{"__!;%:\"@1", "valid", false},
		{"__!;%:\"@1а", "valid", false},
		{"ХА-ХА_HA-HA_123_aa_фф", "valid", false},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := Struct(password{tt.password})

			assert.Equal(t, tt.isErr, err != nil, fmt.Sprintf("password=%s", tt.password))
		})
	}
}
