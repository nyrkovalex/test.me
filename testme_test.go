package testme

import "testing"

func TestShouldPass(t *testing.T) {
	E(t).Expect(1).ToBe(1)
}

func TestShouldNotFail(t *testing.T) {
	E(t).Expect(1).NotToBe(2)
}

func TestShouldPanic(t *testing.T) {
	E(t).Expect(func() {
		panic("foo")
	}).ToPanic("foo")
}
