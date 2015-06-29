package testme

import (
	"reflect"
	"testing"
)

// E function wraps standard testing type T with expect assertions
func E(t *testing.T) Expecter {
	return &tester{t}
}

// Expecter interface serves as a starting point for expect assertions
type Expecter interface {
	// Expect accpets actual value to be compared or tested against
	// expected value passed to Expecation methods
	Expect(actual interface{}) Expectation
}

// Expectation interface provides assertion methods
type Expectation interface {
	// ToBe performs simple comparison of an actual value passed to Expecter and
	// expected value provided
	ToBe(expected interface{})

	// NotToBe performs negated comparison similar to ToBe
	NotToBe(expected interface{})

	// ToPanic checks if a function wrapped with Expect panics with an argument
	// provided
	ToPanic(expected interface{})
}

type tester struct {
	t *testing.T
}

func (t *tester) Expect(actual interface{}) Expectation {
	return &expectation{t.t, actual}
}

type expectation struct {
	t      *testing.T
	actual interface{}
}

func (e *expectation) fail(condition string, expected interface{}) {
	e.t.Errorf("Expected %v "+condition+" %v", e.actual, expected)
}

func (e *expectation) Expect(actual interface{}) *expectation {
	e.actual = actual
	return e
}

func (e *expectation) ToBe(expected interface{}) {
	if expected != e.actual {
		e.fail("to be", expected)
	}
}

func (e *expectation) NotToBe(expected interface{}) {
	if expected == e.actual {
		e.fail("not to be", expected)
	}
}

func (e *expectation) ToPanic(expected interface{}) {
	defer func() {
		expect := &tester{e.t}
		err := recover()
		expect.Expect(err).ToBe(expected)
	}()
	reflect.ValueOf(e.actual).Call([]reflect.Value{})
	e.fail("to panic with", expected)
}
