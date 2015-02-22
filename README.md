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
	Run(t, MyTest{}) // Hook up to the go testing library
}

func (m MyTest) TestShouldPass(e *Expect) {
    e.Expect(1).ToBe(1)
}

func (m MyTest) TestShouldFail(e *Expect) {
    e.Expect("foo").ToBe("bar")
}

```

and run `go test`
