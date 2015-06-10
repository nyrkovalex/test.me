package testme

import "testing"

type MyTest struct{}

func (m MyTest) TestShouldPass(expect Expect) {
	expect(1).ToBe(1)
}

func (m MyTest) TestShouldNotFail(expect Expect) {
	expect(1).NotToBe(2)
}

func (m MyTest) TestShouldPanic(expect Expect) {
	expect(func() {
		panic("foo")
	}).ToPanic("foo")
}

func TestSetUp(t *testing.T) {
	Run(t, MyTest{})
}
