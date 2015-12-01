// Package match implements several Matchers for use in unit testing.
//
// Chaining
//
// The matchers can be used individually or chained together to make complex
// statements.
//     func TestComplexStatement(t *testing.T) {
//         value := DoSomethingThatReturns()
//         match.IsNotNil(t, value).KindOf(reflect.String)
//     }
//
// Fatal Errors
//
// In order to handle fatal errors where further tests should not be
// executed, there is a .Fatal function that will cause the test to fail
// immediately.
// If .Fatal is used in a chain, it will only apply to the previous statements
// and not the subsequent ones.
//     func TestFatalError(t *testing.T) {
//         value := DoSomethingThatReturns()
//         match.IsNotNil(t, value).Fatal().KindOf(reflect.String)
//         DoSomethingWith(value)
//     }
// In the above example, the test will fail immediately if value is nil, but
// will continue if the value is not a string.
package match
