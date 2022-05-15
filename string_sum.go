package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var operator = "+"
	rplstr := strings.Replace(input, "-+", "-", -1)
	rplstr = strings.Replace(rplstr, "+-", "-", -1)

	trstr := strings.Trim(rplstr, " ")

	if trstr == "" {
		return "", errorEmptyInput
	}

	val1 := strings.Split(strings.Trim(trstr, " "), "+")

	if len(val1) == 1 {
		val1 = strings.Split(strings.Trim(trstr, " "), "-")

		if len(val1) == 1 {
			return "", fmt.Errorf("not enough operators")
		}

		operator = "-"
	}

	if len(val1) > 2 {
		if len(val1) == 3 && val1[0] == "" {
			val1[0] = "-" + val1[1]
			val1[1] = val1[2]
		} else {
			return "", fmt.Errorf("want expecting two operands, but received more or less")
		}
	}

	left, err1 := strconv.Atoi(strings.Trim(val1[0], " "))

	if err1 != nil {
		return "", err.(*strconv.NumError).Err
	}

	right, err2 := strconv.Atoi(strings.Trim(val1[1], " "))

	if err2 != nil {
		return "", err.(*strconv.NumError)
	}

	if err == nil {
		if operator == "+" {
			output = strconv.Itoa(left + right)
		} else {
			output = strconv.Itoa(left - right)
		}

	}

	return output, err
}
