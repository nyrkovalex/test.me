# test.me
Simple test assertion wrapper for Go tests

# Usage

Usage is very simple, just do


```go
package main

import (
	. "github.com/nyrkovalex/testme"
	"testing"
)

type MyTest struct{}

func TestSetUp(t *testing.T) {
	Run(t, MyTest{})
}

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

```

and run `go test`
