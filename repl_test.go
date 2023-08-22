package main

import (
	"testing"
)

func TestExtractParams(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  map  ",
			expected: []string{"map"},
		},
		{
			input:    "  explore  world  ",
			expected: []string{"explore", "world"},
		},
		{
			input:    "  eXpLoRe  MiDdLE EARTH  ",
			expected: []string{"explore", "middle", "earth"},
		},
		{
			input:    "explore eterna-city-area",
			expected: []string{"explore", "eterna-city-area"},
		},
	}

	for _, c := range cases {
		actual := extractParams(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
