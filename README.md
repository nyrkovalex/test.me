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

```

and run `go test`
