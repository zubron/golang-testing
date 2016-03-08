package operations

import (
	"math/rand"
	"testing"
)

var sumTests = []struct {
	values   []int
	expected int
}{
	{[]int{}, 0},
	{[]int{1, 2, 3}, 6},
	{[]int{1, 2, 3, 4}, 10},
	{[]int{10, 17, 12, 2}, 41},
	{[]int{-1, 3, -4, 5}, 3},
}

func TestSum(t *testing.T) {
	for _, testCase := range sumTests {
		actual := Sum(testCase.values)
		if actual != testCase.expected {
			t.Errorf("Test failed. Expected: %d, Actual %d", testCase.expected, actual)
		}
	}
}

func benchmarkSum(values []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum(values)
	}
}

func BenchmarkSum1(b *testing.B) {
	benchmarkSum([]int{5}, b)
}

func BenchmarkSum2(b *testing.B) {
	benchmarkSum([]int{5, 4}, b)
}

func BenchmarkSum5(b *testing.B) {
	benchmarkSum([]int{5, 4, 3, 2, 1}, b)
}

func BenchmarkSum10(b *testing.B) {
	benchmarkSum([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, b)
}

func BenchmarkSum20(b *testing.B) {
	benchmarkSum([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, b)
}

func BenchmarkSum100(b *testing.B) {
	values := make([]int, 100)
	for i := 0; i < 100; i++ {
		values = append(values, rand.Int())
	}
	b.ResetTimer()
	benchmarkSum(values, b)
}
