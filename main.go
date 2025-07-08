package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(slice []string) []string {
	var output []string
	for i := len(slice) - 1; i >= 0; i-- {
		output = append(output, slice[i])
	}
	return output
}

func fizzBuzz(i int) string {
	var output []string
	if i%3 == 0 {
		output = append(output, "Fizz")
	}
	if i%5 == 0 {
		output = append(output, "Buzz")
	}
	if i%7 == 0 {
		output = append(output, "Bang")
	}
	if i%11 == 0 {
		output = []string{"Bong"}
	}
	if i%17 == 0 {
		output = reverse(output)
	}
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

// This is our main function, this executes by default when we run the main package.
func main() {
	num, err := getNumber()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i <= num; i++ {
		fmt.Println(fizzBuzz(i))
	}

}
