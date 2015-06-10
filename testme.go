package testme

import (
	"reflect"
	"testing"
)

type Expect func(actual interface{}) *Expectation

type tester struct {
	t    *testing.T
	name string
}

func (e *tester) Expect(actual interface{}) *Expectation {
	return &Expectation{e.t, e.name, actual}
}

type Expectation struct {
	t      *testing.T
	name   string
	actual interface{}
}

func (e *Expectation) ToBe(expected interface{}) {
	if expected != e.actual {
		e.t.Logf("%s: expected %v to be %v", e.name, e.actual, expected)
		e.t.Fail()
	}
}

func (e *Expectation) NotToBe(expected interface{}) {
	if expected == e.actual {
		e.t.Logf("%s: expected %v not to be %v", e.name, e.actual, expected)
		e.t.Fail()
	}
}

func (e *Expectation) ToPanic(expected interface{}) {
	defer func() {
		expect := &tester{e.t, e.name}
		err := recover()
		expect.Expect(err).ToBe(expected)
	}()
	reflect.ValueOf(e.actual).Call([]reflect.Value{})
	e.t.Logf("%s: expected %v to panic with %v", e.name, e.actual, expected)
	e.t.Fail()
}

func Run(t *testing.T, suite interface{}) {
	suiteType := reflect.TypeOf(suite)
	for i := 0; i < suiteType.NumMethod(); i++ {
		method := suiteType.Method(i)
		injectable := &tester{t, method.Name}
		method.Func.Call([]reflect.Value{
			reflect.ValueOf(suite),
			reflect.ValueOf(injectable.Expect),
		})
	}
}
