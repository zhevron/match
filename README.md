match - Matcher library for testing
===================================

[![wercker status](https://app.wercker.com/status/d82e11215fa0ebf705f77a6b864de6f6/s/master "wercker status")](https://app.wercker.com/project/bykey/d82e11215fa0ebf705f77a6b864de6f6)
[![Coverage Status](https://coveralls.io/repos/zhevron/match/badge.svg?branch=master&service=github)](https://coveralls.io/github/zhevron/match?branch=master)
[![GoDoc](https://godoc.org/github.com/zhevron/match?status.svg)](https://godoc.org/github.com/zhevron/match)

**match** is a matcher library for use in unit tests written in [Google Go](https://golang.org).  
For full package documentation, see the GoDoc link above.

## Matchers

### Nil/Not Nil
```go
match.IsNil(t, currentValue)
match.IsNotNil(t, currentValue)
```

### Equals
```go
match.Equals(t, currentValue, shouldBeValue)
match.NotEquals(t, currentValue, shouldNotBeValue)
```

### Less/Greater Than
```go
match.LessThan(t, currentValue, shouldBeLessThan)
match.GreaterThan(t, currentValue, shouldBeGreaterThan)
```

### Contains
```go
match.Contains(t, currentValue, sliceItem)
match.Contains(t, currentValue, mapKey)
match.Contains(t, currentValue, str)
```

### Matches (Regex)
```go
match.Matches(t, currentValue, pattern)
```

### KindOf
```go
match.KindOf(t, currentValue, reflect.String)
```

## Chaining

All matchers can be chained together to build more complex tests.
This can be done using the `match.Match` function with the initial value.

```go
match.Match(t, currentValue).
  Equals(shouldBeValue).
  KindOf(match.String).
  Fatal()
```

## Errors

All errors will be reported via `t.Errorf`.
In order to mark a failure as fatal, chain the call with `.Fatal()`

```go
match.Equals(t, currentValue, shouldBeValue).Fatal()
```

## License

**match** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
