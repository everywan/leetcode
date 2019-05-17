package utils

import (
	"fmt"

	"github.com/smartystreets/assertions"
)

// ShouldEqual is 比较相等
func ShouldEqual(actual interface{}, expected ...interface{}) {
	msg := assertions.ShouldEqual(actual, expected...)
	if msg == "" {
		msg = "success"
	}
	fmt.Printf("%s\n-----------\n", msg)
}

// ShouldResemble is 数组比较
func ShouldResemble(actual interface{}, expected ...interface{}) {
	msg := assertions.ShouldResemble(actual, expected...)
	if msg == "" {
		msg = "success"
	}
	fmt.Printf("%s\n-----------\n", msg)
}
