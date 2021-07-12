package helper

import "testing"

func TestHello(t *testing.T) {
	result := HelloWorld("Hello")
	if result != "Hello" {
		t.Fail()
	}
}