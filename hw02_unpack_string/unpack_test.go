package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		// my
		{input: `\1`, expected: `1`},
		// {input: "\\", expected: "\\"},
		{input: "abc🐋5", expected: "abc🐋🐋🐋🐋🐋"},
		// {input: `\`, expected: `\`},
		{input: `\4\5`, expected: `45`},
		{input: `\5`, expected: `5`},
		{input: `a\\t`, expected: `a\t`},
		{input: `\\n5`, expected: `\nnnnn`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", `qw\ne`, "a\\🐋", `\`, `\\\`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

// func BenchmarkUnpack(b *testing.B) {
// 	// Unpack("a4bc2d5e")
// 	// Unpack("abcdef")
// 	Unpack("a4b10c2d5e111")
// }
