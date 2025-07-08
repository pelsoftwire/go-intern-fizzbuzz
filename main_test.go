package main

import "testing"

type fizzBuzzTest struct {
	input          int
	expectedOutput string
}

var fizzBuzzTests = []fizzBuzzTest{
	{1, "1"},          // works when no rules
	{3, "Fizz"},       // Fizz works alone
	{5, "Buzz"},       // Buzz works alone
	{15, "FizzBuzz"},  // Fizz, Buzz work together
	{7, "Bang"},       // Bang works alone
	{21, "FizzBang"},  // Bang works with other rules
	{11, "Bong"},      // Bong works alone
	{33, "Bong"},      // Bong overrides other rules
	{17, "17"},        // 17 alone does nothing
	{255, "BuzzFizz"}, // 17 rule reverses
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range fizzBuzzTests {
		result := fizzBuzz(test.input)
		if result != test.expectedOutput {
			t.Errorf("FizzBuzz(%d): Expected '%s', got '%s'", test.input, test.expectedOutput, result)
		}
	}
}
