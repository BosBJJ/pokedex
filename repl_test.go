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
		input: "green red blue",
		expected: []string{"green", "red", "blue"},
	},
	{
		input: " onepokemon two pokemon",
		expected: []string{"onepokemon", "two", "pokemon"},
	},
}
for _, c := range cases{
	actual := cleanInput(c.input)

	if len(actual) != len(c.expected){
		t.Errorf("Expected length %d, got %d", len(c.expected), len(actual))
		continue
	}
	
	for i := range actual{
		word := actual[i]
		expectedWord := c.expected[i]

		
		if word != expectedWord{
			t.Errorf("expected %v, got %v", expectedWord, word)
		}
	}
}
}