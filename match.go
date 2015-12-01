package match

import (
	"reflect"
	"testing"
)

// Match creates a new Matcher for chaining conditions.
func Match(t *testing.T, value interface{}) *Matcher {
	return &Matcher{
		t:      t,
		value:  value,
		failed: false,
	}
}

// IsNil creates a new Matcher and asserts that value is nil.
func IsNil(t *testing.T, value interface{}) *Matcher {
	return Match(t, value).IsNil()
}

// IsNotNil creates a new Matcher and asserts that value is not nil.
func IsNotNil(t *testing.T, value interface{}) *Matcher {
	return Match(t, value).IsNotNil()
}

// Equals creates a new Matcher and asserts that value is equal to other.
func Equals(t *testing.T, value, other interface{}) *Matcher {
	return Match(t, value).Equals(other)
}

// NotEquals creates a new Matcher and asserts that value is not equal to other.
func NotEquals(t *testing.T, value, other interface{}) *Matcher {
	return Match(t, value).NotEquals(other)
}

// LessThan creates a new Matcher and asserts that value is less than other.
func LessThan(t *testing.T, value, other interface{}) *Matcher {
	return Match(t, value).LessThan(other)
}

// GreaterThan creates a new Matcher and asserts that value is greater than other.
func GreaterThan(t *testing.T, value, other interface{}) *Matcher {
	return Match(t, value).GreaterThan(other)
}

// Contains creates a new Matcher and asserts that value contains other.
func Contains(t *testing.T, value, other interface{}) *Matcher {
	return Match(t, value).Contains(other)
}

// Matches creates a new Matcher and asserts that value matches the pattern.
func Matches(t *testing.T, value interface{}, pattern string) *Matcher {
	return Match(t, value).Matches(pattern)
}

// KindOf creates a new Matcher and asserts that value is of the kind.
func KindOf(t *testing.T, value interface{}, kind reflect.Kind) *Matcher {
	return Match(t, value).KindOf(kind)
}
