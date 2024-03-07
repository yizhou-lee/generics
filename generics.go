// Package generics provides functions for summing values in maps with different types.
// It includes functions for summing int64, float64, either int64 or float64 (SumIntsOrFloats),
// and a type defined as Number which can be either int64 or float64 (SumNumbers).
// The functions SumIntsOrFloats and SumNumbers use type parameters and type unions introduced in Go 1.18.
package generics

// Number is an interface that represents a type that can be either int64 or float64.
type Number interface {
	int64 | float64 // A type that is either int64 or float64
}

// SumInts calculates the sum of all int64 values in the given map.
// It iterates over the map and adds up all the values, then returns the total sum.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats calculates the sum of all float64 values in the given map.
// It iterates over the map and adds up all the values, then returns the total sum.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats calculates the sum of all int64 or float64 values in the given map.
// It iterates over the map and adds up all the values, then returns the total sum.
// The function uses type parameters, where K is a comparable type used as the key in the map,
// and V is either int64 or float64 used as the value in the map.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers calculates the sum of all Number values (either int64 or float64) in the given map.
// It iterates over the map and adds up all the values, then returns the total sum.
// The function uses type parameters, where K is a comparable type used as the key in the map,
// and V is a Number (either int64 or float64) used as the value in the map.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
