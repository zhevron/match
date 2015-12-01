package match

import (
	"reflect"
	"testing"
)

func TestMatcherFatal(t *testing.T) {
	done := make(chan bool)
	go func() {
		test := &testing.T{}
		m := Matcher{test, 0}
		defer func() {
			recover()
			if !test.Failed() {
				t.Error("expected failure, but the test passed")
			}
			done <- true
		}()
		m.Equals(1).Fatal().Equals(0)
	}()
	<-done
}

func TestMatcherIsNil(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, nil}
	m.IsNil()
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, 0}
	m.IsNil()
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherIsNotNil(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0}
	m.IsNotNil()
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, nil}
	m.IsNotNil()
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherEquals(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0}
	m.Equals(0)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, nil}
	m.Equals(0)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherNotEquals(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, nil}
	m.NotEquals(0)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, 0}
	m.NotEquals(0)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherLessThan_Int(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, int(0)}
	m.LessThan(1)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, int(2)}
	m.LessThan(1)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherLessThan_Uint(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, uint(0)}
	m.LessThan(1)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, uint(2)}
	m.LessThan(1)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherLessThan_Float(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, float64(0.0)}
	m.LessThan(1.0)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, float64(2.0)}
	m.LessThan(1.0)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherLessThan_NonNumeric(t *testing.T) {
	done := make(chan bool)
	go func() {
		test := &testing.T{}
		m := Matcher{test, "0"}
		defer func() {
			recover()
			if !test.Failed() {
				t.Error("expected failure, but the test passed")
			}
			done <- true
		}()
		m.LessThan(1)
	}()
	<-done
}

func TestMatcherGreaterThan_Int(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, int(2)}
	m.GreaterThan(1)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, int(0)}
	m.GreaterThan(1)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherGreaterThan_Uint(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, uint(2)}
	m.GreaterThan(1)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, uint(0)}
	m.GreaterThan(1)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherGreaterThan_Float(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, float64(2.0)}
	m.GreaterThan(1.0)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m = Matcher{test, float64(0.0)}
	m.GreaterThan(1.0)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherGreaterThan_NonNumeric(t *testing.T) {
	done := make(chan bool)
	go func() {
		test := &testing.T{}
		m := Matcher{test, "2"}
		defer func() {
			recover()
			if !test.Failed() {
				t.Error("expected failure, but the test passed")
			}
			done <- true
		}()
		m.GreaterThan(1)
	}()
	<-done
}

func TestMatcherContains_Array(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, []int{0, 1, 2, 3}}
	m.Contains(1)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m.Contains(4)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherContains_Map(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, map[string]int{"a": 0, "b": 1}}
	m.Contains("b")
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m.Contains("c")
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherContains_String(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, "abc"}
	m.Contains("bc")
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m.Contains("cd")
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherContains_InvalidType(t *testing.T) {
	done := make(chan bool)
	go func() {
		test := &testing.T{}
		m := Matcher{test, 0}
		defer func() {
			recover()
			if !test.Failed() {
				t.Error("expected failure, but the test passed")
			}
			done <- true
		}()
		m.Contains(0)
	}()
	<-done
}

func TestMatcherMatches(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, "this is a test"}
	m.Matches("^this is a test$")
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m.Matches("^this is not a test$")
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}

func TestMatcherMatches_InvalidRegex(t *testing.T) {
	done := make(chan bool)
	go func() {
		test := &testing.T{}
		m := Matcher{test, 0}
		defer func() {
			recover()
			if !test.Failed() {
				t.Error("expected failure, but the test passed")
			}
			done <- true
		}()
		m.Matches("^0$")
	}()
	<-done
}

func TestMatcherMatches_NonString(t *testing.T) {
	done := make(chan bool)
	go func() {
		test := &testing.T{}
		m := Matcher{test, "this is a test"}
		defer func() {
			recover()
			if !test.Failed() {
				t.Error("expected failure, but the test passed")
			}
			done <- true
		}()
		m.Matches("(?x:^this is a test$)")
	}()
	<-done
}

func TestMatcherKindOf(t *testing.T) {
	test := &testing.T{}
	m := Matcher{test, 0}
	m.KindOf(reflect.Int)
	if test.Failed() {
		t.Error("expected pass, but the test failed")
	}
	m.KindOf(reflect.String)
	if !test.Failed() {
		t.Error("expected failure, but the test passed")
	}
}
