package operations

import (
	"testing"
)

func TestAddZero(t *testing.T) {
	expected := 0
	actual := Add(0, 0)
	if expected != actual {
		t.Errorf("Test failed. Expected: %d, Actual %d", expected, actual)
	}
}

func TestAddPositive(t *testing.T) {
	expected := 17
	actual := Add(9, 8)
	if expected != actual {
		t.Errorf("Test failed. Expected: %d, Actual %d", expected, actual)
	}
}

func TestAddNegative(t *testing.T) {
	expected := -12
	actual := Add(-7, -5)
	if expected != actual {
		t.Errorf("Test failed. Expected: %d, Actual %d", expected, actual)
	}
}

func TestAddPositiveNegative(t *testing.T) {
	expected := 4
	actual := Add(7, -3)
	if expected != actual {
		t.Errorf("Test failed. Expected: %d, Actual %d", expected, actual)
	}
}

var addTests = []struct {
	a, b     int // input
	expected int // expected result
}{
	{1, 3, 4},
	{0, 5, 5},
	{-3, 3, 0},
	{127, -31, 96},
	{-13, -45, -58},
}

func TestAdd(t *testing.T) {
	for _, testCase := range addTests {
		actual := Add(testCase.a, testCase.b)
		if actual != testCase.expected {
			t.Errorf("Test failed. Expected: %d, Actual %d", testCase.expected, actual)
		}
	}
}
