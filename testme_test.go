package testme

import "testing"

type MyTest struct{}

func (m MyTest) TestShouldPass(e *Expect) {
	e.Expect(1).ToBe(1)
}

func (m MyTest) TestShouldFail(e *Expect) {
	e.Expect("foo").ToBe("bar")
}

func (m MyTest) TestShouldPanic(e *Expect) {
	e.Expect(func() {
		panic("foo")
	}).ToPanic("foo")
}

func (m MyTest) TestShouldLogTheMessage(e *Expect) {
	e.Log("test %d", 1)
}

func TestSetUp(t *testing.T) {
	Run(t, MyTest{})
}
