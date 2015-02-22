package testme

import "testing"

type MyTest struct{}

func (m MyTest) TestShouldPass(e *Expect) {
	e.Expect(1).ToBe(1)
}

func (m MyTest) TestShouldFail(e *Expect) {
	e.Expect("foo").ToBe("bar")
}

func TestSetUp(t *testing.T) {
	Run(t, MyTest{})
}
