package main

import "testing"

func TestCleanInput(t *testing.T) {

    cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
			// test size
			if len(actual) != len(c.expected) { t.Errorf("Incorrect length") }
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// test each word accuracy
			if word != expectedWord { t.Errorf("%s does not match %s", word, expectedWord) }
		}
	}
}