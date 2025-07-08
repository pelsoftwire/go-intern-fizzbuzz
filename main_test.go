package main

import "testing"

type fizzBuzzTest struct {
	input          int
	expectedOutput string
}

var fizzBuzzTests = []fizzBuzzTest{
	{1, "1"},
	{3, "Fizz"},
	{5, "Buzz"},
	{15, "FizzBuzz"},
	{7, "Bang"},
	{21, "FizzBang"},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range fizzBuzzTests {
		result := fizzBuzz(test.input)
		if result != test.expectedOutput {
			t.Errorf("FizzBuzz(%d): Expected '%s', got '%s'", test.input, test.expectedOutput, result)
		}
	}
}
