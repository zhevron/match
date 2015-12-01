package match

import (
	"reflect"
	"testing"
)

// Matcher is a helper object for chaining matcher conditions.
type Matcher struct {
	t     *testing.T
	value interface{}
}

// Fatal causes the unit test to immediately fail is one of the previous
// conditions has fail.
// Note: This will only affect previous conditions, not subsequent ones. You
// can build complex conditions this way.
// Ex: match.IsNotNil(t, value).Fatal().Equals(other)
func (m *Matcher) Fatal() *Matcher {
	if m.t.Failed() {
		m.t.FailNow()
	}
	return m
}

// IsNil asserts that the matched value is nil.
func (m *Matcher) IsNil() *Matcher {
	// TODO: Implement Matcher.IsNil
	return m
}

// IsNotNil asserts that the matched value is not nil.
func (m *Matcher) IsNotNil() *Matcher {
	// TODO: Implement Matcher.IsNotNil
	return m
}

// Equals asserts that the matched value is equal to value.
func (m *Matcher) Equals(value interface{}) *Matcher {
	// TODO: Implement Matcher.Equals
	return m
}

// NotEquals asserts that the match value is not equal to value.
func (m *Matcher) NotEquals(value interface{}) *Matcher {
	// TODO: Implement Matcher.NotEquals
	return m
}

// LessThan asserts that the matched value is less than value.
// Note: This function will only work on numeric values.
// If a non-numeric value is passed, a fatal error will be thrown.
func (m *Matcher) LessThan(value interface{}) *Matcher {
	// TODO: Implement Matcher.LessThan
	return m
}

// GreaterThan asserts that the matched value is greater then value.
// Note: This function will only work on numeric values.
// If a non-numeric value is passed, a fatal error will be thrown.
func (m *Matcher) GreaterThan(value interface{}) *Matcher {
	// TODO: Implement Matcher.GreaterThan
	return m
}

// Contains asserts that the matched value contains value.
// If used on an array, asserts that the array contains value.
// If used on a map, asserts that the map contains a key with value.
// If used on a string, asserts that the string contains value.
func (m *Matcher) Contains(value interface{}) *Matcher {
	// TODO: Implement Matcher.Contains
	return m
}

// Matches asserts that the matched value matches the regex in pattern.
// Note: If the regex cannot be compiled, a fatal error will be thrown.
// If used on a non-string value, a fatal error will be thrown.
func (m *Matcher) Matches(pattern string) *Matcher {
	// TODO: Implement Matcher.Matches
	return m
}

// KindOf asserts that the matched value of the kind.
func (m *Matcher) KindOf(t reflect.Kind) *Matcher {
	// TODO: Implement Matcher.KindOf
	return m
}
