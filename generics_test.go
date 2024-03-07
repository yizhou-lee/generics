package generics

import (
	"fmt"
	"testing"
)

func TestSumInts(t *testing.T) {
	m := map[string]int64{"one": 1, "two": 2, "three": 3}
	sum := SumInts(m)
	if sum != 6 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 6)
	}
}

func TestSumFloats(t *testing.T) {
	m := map[string]float64{"one": 1.1, "two": 2.2, "three": 3.3}
	sum := SumFloats(m)
	if sum != 6.6 {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", sum, 6.6)
	}
}

func TestSumIntsOrFloats(t *testing.T) {
	m := map[string]int64{"one": 1, "two": 2, "three": 3}
	sum := SumIntsOrFloats(m)
	if sum != 6 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 6)
	}
}

func TestSumNumbers(t *testing.T) {
	m := map[string]int64{"one": 1, "two": 2, "three": 3}
	sum := SumNumbers(m)
	if sum != 6 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 6)
	}
}

func ExampleSumIntsOrFloats() {
	m := map[string]int64{"one": 1, "two": 2, "three": 3}
	sum := SumIntsOrFloats(m)
	fmt.Println(sum)
	// Output: 6
}

func ExampleSumNumbers() {
	m := map[string]int64{"one": 1, "two": 2, "three": 3}
	sum := SumNumbers(m)
	fmt.Println(sum)
	// Output: 6
}
