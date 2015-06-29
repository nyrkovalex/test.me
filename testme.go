// Package testme provides simple expect-style assertioin by wrapping standard
// *testing.T type with E(*t) function
package testme

import (
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const libraryName = "testme.go"

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

func firstExternalFileLine() (file string, line int) {
	pc := make([]uintptr, 5)
	written := runtime.Callers(1, pc)
	for i := 0; i < written; i++ {
		current := pc[i] - 1
		fn := runtime.FuncForPC(current)
		file, line = fn.FileLine(current)
		if strings.LastIndex(file, libraryName) == -1 {
			return
		}
	}
	return
}

func fileLine() (file string, line int) {
	file, line = firstExternalFileLine()
	splitted := strings.Split(file, string(os.PathSeparator))
	file = splitted[len(splitted)-1]
	return
}

func (e *expectation) logError(condition string, expected interface{}, file string, line int) {
	e.t.Errorf("\n> %s:%d: Expected %v "+condition+" %v", file, line, e.actual, expected)
}

func (e *expectation) fail(condition string, expected interface{}) {
	file, line := fileLine()
	e.logError(condition, expected, file, line)
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
	file, line := fileLine()
	defer func() {
		err := recover()
		if err != e.actual {
			e.logError("to panic with", expected, file, line)
		}
	}()
	reflect.ValueOf(e.actual).Call([]reflect.Value{})
	e.fail("to panic with", expected)
}
