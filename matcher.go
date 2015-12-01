package match

import (
	"reflect"
	"regexp"
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
	if m.value != nil {
		m.t.Errorf("expected == nil, got %v", m.value)
	}
	return m
}

// IsNotNil asserts that the matched value is not nil.
func (m *Matcher) IsNotNil() *Matcher {
	if m.value == nil {
		m.t.Error("expected != nil, got nil")
	}
	return m
}

// Equals asserts that the matched value is equal to value.
func (m *Matcher) Equals(value interface{}) *Matcher {
	if !reflect.DeepEqual(m.value, value) {
		m.t.Errorf("expected == %v, got %v", value, m.value)
	}
	return m
}

// NotEquals asserts that the match value is not equal to value.
func (m *Matcher) NotEquals(value interface{}) *Matcher {
	if reflect.DeepEqual(m.value, value) {
		m.t.Errorf("expected != %v, got %v", value, m.value)
	}
	return m
}

// LessThan asserts that the matched value is less than value.
// Note: This function will only work on numeric values.
// If a non-numeric value is passed, a fatal error will be thrown.
func (m *Matcher) LessThan(value interface{}) *Matcher {
	ok := true
	rv := reflect.ValueOf(m.value)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ok = rv.Int() < reflect.ValueOf(value).Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ok = rv.Uint() < uint64(reflect.ValueOf(value).Int())
	case reflect.Float32, reflect.Float64:
		ok = rv.Float() < reflect.ValueOf(value).Float()
	default:
		m.t.Fatalf("expected numeric value, got %v", m.value)
	}
	if !ok {
		m.t.Errorf("expected < %v, got %v", value, m.value)
	}
	return m
}

// GreaterThan asserts that the matched value is greater then value.
// Note: This function will only work on numeric values.
// If a non-numeric value is passed, a fatal error will be thrown.
func (m *Matcher) GreaterThan(value interface{}) *Matcher {
	ok := true
	rv := reflect.ValueOf(m.value)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ok = rv.Int() > reflect.ValueOf(value).Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ok = rv.Uint() > uint64(reflect.ValueOf(value).Int())
	case reflect.Float32, reflect.Float64:
		ok = rv.Float() > reflect.ValueOf(value).Float()
	default:
		m.t.Fatalf("expected numeric value, got %v", m.value)
	}
	if !ok {
		m.t.Errorf("expected < %v, got %v", value, m.value)
	}
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
	rv := reflect.ValueOf(m.value)
	if rv.Kind() != reflect.String {
		m.t.Fatalf("expected string, got %s", rv.Kind().String())
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		m.t.Fatal(err)
	}
	if !re.MatchString(rv.String()) {
		m.t.Errorf("expected %#q to match pattern %#q", rv.String(), pattern)
	}
	return m
}

// KindOf asserts that the matched value of the kind.
func (m *Matcher) KindOf(kind reflect.Kind) *Matcher {
	rv := reflect.ValueOf(m.value)
	if rv.Kind() != kind {
		m.t.Errorf("expected kind %s, got %s", kind.String(), rv.Kind().String())
	}
	return m
}
