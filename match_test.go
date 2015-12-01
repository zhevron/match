package match

import (
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	res := Match(t, nil)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestIsNil(t *testing.T) {
	res := IsNil(t, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestIsNotNil(t *testing.T) {
	res := IsNotNil(t, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestEquals(t *testing.T) {
	res := Equals(t, 0, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestNotEquals(t *testing.T) {
	res := NotEquals(t, 0, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestLessThan(t *testing.T) {
	res := LessThan(t, 0, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestGreaterThan(t *testing.T) {
	res := GreaterThan(t, 0, 0)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestContains(t *testing.T) {
	res := Contains(t, "", "")
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestMatches(t *testing.T) {
	res := Matches(t, "", ".*")
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}

func TestKindOf(t *testing.T) {
	res := KindOf(t, 0, reflect.Int32)
	if res == nil {
		t.Errorf("Expected a value, got nil")
	}
}
