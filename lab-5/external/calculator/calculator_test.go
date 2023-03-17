package calculator

import "testing"

func Test_Sum_TwoNumbers_Correctly(t *testing.T) {
	expected := 9.0
	actual := Sum(4, 5)
	if expected != actual {
		t.Fatalf("expect: %f, actual: %f", expected, actual)
	}
}

func Test_Subtract_TwoNumbers_Correctly(t *testing.T) {
	expected := -7.0
	actual := Sub(3, 10)
	if expected != actual {
		t.Fatalf("expect: %f, actual: %f", expected, actual)
	}
}

func Test_Multiply_TwoNumbers_Correctly(t *testing.T) {
	expected := 64.0
	actual := Mult(4, 16)
	if expected != actual {
		t.Fatalf("expect: %f, actual: %f", expected, actual)
	}
}

func Test_Divide_TwoNumbers_Correctly(t *testing.T) {
	expected := 3.0
	actual, err := Div(27, 9)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expect: %f, actual: %f", expected, actual)
	}
}

func Test_Fail_Divide_By_Zero(t *testing.T) {
	_, err := Div(8, 0)
	if err == nil {
		t.Fatal("should throw an error when trying to divide by zero")
	}
}
