package hurdle_test

import (
	"hurdle-server/hurdle"
	"testing"
)

func TestHurdle(t *testing.T) {
	tests := []struct {
		answer   string
		guess    string
		expected string
	}{
		{
			"WORLD",
			"HELLO",
			"00021",
		},
		{
			"WORLD",
			"WORLD",
			"22222",
		},
		{
			"WORLD",
			"DLROW",
			"11211",
		},
		{
			"SPACE",
			"APPLE",
			"12002",
		},
	}

	for _, test := range tests {
		actual := hurdle.Hurdle(test.answer, test.guess)
		if actual != test.expected {
			t.Errorf("hurdle(%s, %s): expected %s, actual %s", test.guess, test.answer, test.expected, actual)
		}
	}
}
