package match

import (
	"reflect"
	"testing"
)

var t = &testing.T{}

func ExampleIsNil_simple() {
	var value string
	IsNil(t, value)
}

func ExampleIsNil_advanced() {
	var value string
	Match(t, value).IsNil()
}

func ExampleIsNotNil_simple() {
	value := 0
	IsNotNil(t, value)
}

func ExampleIsNotNil_advanced() {
	value := 0
	Match(t, value).IsNotNil()
}

func ExampleEquals_simple() {
	value := 0
	Equals(t, value, 0)
}

func ExampleEquals_advanced() {
	value := 0
	Match(t, value).Equals(1)
}

func ExampleNotEquals_simple() {
	value := 0
	NotEquals(t, value, 0)
}

func ExampleNotEquals_advanced() {
	value := 0
	Match(t, value).NotEquals(1)
}

func ExampleLessThan_simple() {
	value := 0
	LessThan(t, value, 1)
}

func ExampleLessThan_advanced() {
	value := 1
	Match(t, value).LessThan(1)
}

func ExampleGreaterThan_simple() {
	value := 1
	GreaterThan(t, value, 0)
}

func ExampleGreaterThan_advanced() {
	value := 1
	Match(t, value).GreaterThan(0)
}

func ExampleContains_simple() {
	value := "this is a string"
	Contains(t, value, "is a")
}

func ExampleContains_advanced() {
	value := "this is a string"
	Match(t, value).Contains("is a")
}

func ExampleContains_array() {
	value := []int{0, 1, 2, 3}
	Contains(t, value, 1)
}

func ExampleContains_map() {
	value := map[string]int{"a": 0, "b": 1, "c": 2}
	Contains(t, value, "b")
}

func ExampleMatches_simple() {
	value := "this is a string"
	Matches(t, value, "^this is a string$")
}

func ExampleMatches_advanced() {
	value := "this is a string"
	Match(t, value).Matches("^this is a string$")
}

func ExampleKindOf_simple() {
	value := 0
	KindOf(t, value, reflect.Int)
}

func ExampleKindOf_advanced() {
	value := 0
	Match(t, value).KindOf(reflect.Int)
}

func ExampleFatal() {
	value := 0
	Equals(t, value, 1).Fatal()
}
