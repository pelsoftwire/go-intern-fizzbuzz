package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// rule stuff
type Rule struct {
	mod       int
	word      string
	wipe      bool
	reverse   bool
	priority  int
	useBefore bool
	before    byte
}

type Rules []Rule

func (rule *Rule) apply(i int, slice []string) []string {
	output := slice
	if i%rule.mod == 0 {
		// reverse if applicable -> wipe if applicable -> calculate before/put at the end
		if rule.reverse {
			output = reverse(output)
		}
		if rule.wipe {
			output = []string{}
		}

		insertAt := -1
		if rule.useBefore {
			insertAt = findFirstBeginningWith(output, rule.before)
		}
		if insertAt == -1 && rule.word != "" {
			output = append(output, rule.word)
		} else if rule.word != "" { // TODO: add boolean for whether or not to use "word"?
			output = slices.Insert(output, insertAt, rule.word)
		}
	}
	return output
}

// TODO: consider rewriting applyAll() so that applied rules are removed from some list to avoid iterating over them again
func (rules *Rules) applyAll(i int, slice []string) []string {
	currentPriority := 0
	appliedARule := true
	output := slice

	for appliedARule {
		appliedARule = false
		for _, rule := range *rules {
			if rule.priority == currentPriority {
				appliedARule = true
				output = rule.apply(i, output)
			}
		}
		currentPriority = currentPriority + 1
	}

	return output
}

func findFirstBeginningWith(slice []string, letter byte) int {
	for i := 0; i < len(slice); i++ {
		if slice[i][0] == letter {
			return i
		}
	}
	return -1
}

func reverse(slice []string) []string {
	var output []string
	for i := len(slice) - 1; i >= 0; i-- {
		output = append(output, slice[i])
	}
	return output
}

func fizzBuzz(baseRules Rules, i int) string {
	output := baseRules.applyAll(i, []string{})

	if len(output) == 0 {
		return strconv.Itoa(i)
	} else {
		return strings.Join(output, "")
	}
}

func getNumber() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a maximum number >> ")
	scanner.Scan()
	return strconv.Atoi(scanner.Text())
}

func getBaseRules() Rules {
	// TODO: refactor to make rules more amenable to extension
	fizzRule := Rule{3, "Fizz", false, false, 0, false, ' '}
	buzzRule := Rule{5, "Buzz", false, false, 1, false, ' '}
	bangRule := Rule{7, "Bang", false, false, 2, false, ' '}
	bongRule := Rule{11, "Bong", true, false, 3, false, ' '}
	revRule := Rule{17, "", false, true, 4, false, ' '}

	return Rules{fizzRule, buzzRule, bangRule, bongRule, revRule}
}

// This is our main function, this executes by default when we run the main package.
func main() {
	num, err := getNumber()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i <= num; i++ {
		fmt.Println(fizzBuzz(getBaseRules(), i))
	}

}
