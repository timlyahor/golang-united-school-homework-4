package string_sum

import (
	"testing"
)

func TestTask1SimpleSum(t *testing.T) {
	sourceData := "3+5"
	res, _ := StringSum(sourceData)

	if res != "8" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "8")
	}
}

func TestTask1SimpleSum2(t *testing.T) {
	sourceData := "5+2"
	res, _ := StringSum(sourceData)

	if res != "7" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "7")
	}
}

func TestTask1SimpleSum3(t *testing.T) {
	sourceData := "-5+2"
	res, _ := StringSum(sourceData)

	if res != "-3" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "-3")
	}
}

func TestTask1SimpleSumWithSpaces(t *testing.T) {
	sourceData := "5 + 3"
	res, _ := StringSum(sourceData)

	if res != "8" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "8")
	}
}

func TestTask1SimpleSumWithDoubleOperators(t *testing.T) {
	sourceData := "5+-3"
	res, _ := StringSum(sourceData)

	if res != "2" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "2")
	}
}

func TestTask1SimpleSumEmptyInput(t *testing.T) {
	sourceData := ""
	_, err := StringSum(sourceData)

	if err == nil {
		t.Errorf("Result was incorect, expected exception")
	}
}

func TestTask1SimpleSumWithDoubleOperators2(t *testing.T) {
	sourceData := "5-+3"
	res, _ := StringSum(sourceData)

	if res != "2" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "2")
	}
}

func TestTask1SimpleSumDoubleOperand(t *testing.T) {
	sourceData := "4 ++ 3"
	_, err := StringSum(sourceData)

	if err == nil {
		t.Errorf("Result was incorect, expected exception")
	}
}

func TestTask1SimpleSumWithDoubleOperators3(t *testing.T) {
	sourceData := "-5-+3"
	res, _ := StringSum(sourceData)

	if res != "-8" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "-8")
	}
}

func TestTask1SimpleSumWithDoubleOperators4(t *testing.T) {
	sourceData := "-5+3"
	res, _ := StringSum(sourceData)

	if res != "-2" {
		t.Errorf("Result was incorect, got: %s, want: %s.", res, "-2")
	}
}

func TestTask1SimpleSumDoubleOperand2(t *testing.T) {
	sourceData := "4--3"
	_, err := StringSum(sourceData)

	if err == nil {
		t.Errorf("Result was incorect, expected exception")
	}
}

//func TestTask1SimpleSumFirstParamIncorrect(t *testing.T) {
//	sourceData := "a + 3"
//	_, err := StringSum(sourceData)
//
//	if !errors.Is(err, err.(*strconv.NumError)) {
//		t.Errorf("Result was incorect, expected exception")
//	}
//}

func TestTask1SimpleSum3Operands(t *testing.T) {
	sourceData := "2+3+5"
	_, err := StringSum(sourceData)

	if err == nil {
		t.Errorf("Result was incorect, expected exception")
	}
}
