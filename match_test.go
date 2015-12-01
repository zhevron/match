package match

import (
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	test := &testing.T{}
	res := Match(test, nil)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestIsNil(t *testing.T) {
	test := &testing.T{}
	res := IsNil(test, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestIsNotNil(t *testing.T) {
	test := &testing.T{}
	res := IsNotNil(test, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestEquals(t *testing.T) {
	test := &testing.T{}
	res := Equals(test, 0, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestNotEquals(t *testing.T) {
	test := &testing.T{}
	res := NotEquals(test, 0, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestLessThan(t *testing.T) {
	test := &testing.T{}
	res := LessThan(test, 0, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestGreaterThan(t *testing.T) {
	test := &testing.T{}
	res := GreaterThan(test, 0, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestContains(t *testing.T) {
	test := &testing.T{}
	res := Contains(test, "", "")
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestMatches(t *testing.T) {
	test := &testing.T{}
	res := Matches(test, "", ".*")
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestKindOf(t *testing.T) {
	test := &testing.T{}
	res := KindOf(test, 0, reflect.Int32)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}
