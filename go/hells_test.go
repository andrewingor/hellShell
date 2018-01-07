// Author: Andre Wingor (http://andr.ru)
// License: https://github.com/andrewingor/hellShell/blob/master/LICENSE

// Package main_test performs testing on main package
package main_test

import (
	"fmt"
	"testing"
)


// Race condition test
func TestGo(t *testing.T) {
	go fmt.Println("Hello")
}

// Example code test
func ExampleOutput() {
	fmt.Println("Hello")
	// Output: Hello
}
