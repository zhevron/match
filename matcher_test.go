package match

import (
	"reflect"
	"testing"
)

func TestMatcherFatal(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0, false}
	m.Equals(0).Fatal().Equals(1)
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m.Equals(1).Fatal().Equals(0)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherIsNil(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, nil, false}
	m.IsNil()
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m = Matcher{test, 0, false}
	m.IsNil()
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherIsNotNil(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0, false}
	m.IsNotNil()
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m = Matcher{test, nil, false}
	m.IsNotNil()
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherEquals(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0, false}
	m.Equals(0)
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m = Matcher{test, nil, false}
	m.Equals(0)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherNotEquals(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, nil, false}
	m.NotEquals(0)
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m = Matcher{test, 0, false}
	m.NotEquals(0)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherLessThan(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0, false}
	m.LessThan(1)
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m = Matcher{test, 2, false}
	m.LessThan(1)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherLessThan_NonNumeric(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, "0", false}
	m.LessThan(1)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherGreaterThan(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 2, false}
	m.GreaterThan(1)
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m = Matcher{test, 0, false}
	m.GreaterThan(1)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherGreaterThan_NonNumeric(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, "2", false}
	m.GreaterThan(1)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherContains_Array(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, []int{0, 1, 2, 3}, false}
	m.Contains(1)
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m.Contains(4)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherContains_Map(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, map[string]int{"a": 0, "b": 1}, false}
	m.Contains("b")
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m.Contains("c")
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherContains_String(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, "abc", false}
	m.Contains("bc")
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m.Contains("cd")
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherMatches(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, "this is a test", false}
	m.Matches("^this is a test$")
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m.Matches("^this is not a test$")
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherMatches_InvalidRegex(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0, false}
	m.Matches("^0$")
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherMatches_NonString(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, "this is a test", false}
	m.Matches("(?x:^this is a test$)")
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}

func TestMatcherKindOf(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0, false}
	m.KindOf(reflect.Int)
	if test.Failed() {
		t.Error("Expected pass, but the test failed")
	}
	m.KindOf(reflect.String)
	if !test.Failed() {
		t.Error("Expected failure, but the test passed")
	}
}
