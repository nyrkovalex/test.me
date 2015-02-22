package testme

import (
	"reflect"
	"testing"
)

type Expect struct {
	t    *testing.T
	name string
}

func (t *Expect) Expect(actual interface{}) *Expectation {
	return &Expectation{t.t, t.name, actual}
}

type Expectation struct {
	t      *testing.T
	name   string
	actual interface{}
}

func (e *Expectation) ToBe(expected interface{}) {
	if expected != e.actual {
		e.t.Logf("%s: expected %s to be %s", e.name, expected, e.actual)
		e.t.Fail()
	}
}

func Run(t *testing.T, suite interface{}) {
	suiteType := reflect.TypeOf(suite)
	for i := 0; i < suiteType.NumMethod(); i++ {
		method := suiteType.Method(i)
		expect := &Expect{t, method.Name}
		method.Func.Call([]reflect.Value{
			reflect.ValueOf(suite),
			reflect.ValueOf(expect),
		})
	}
}
