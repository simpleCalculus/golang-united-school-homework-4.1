package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	// "Hello world!"
	have := ReverseString("Hello world!")
	want := "!dlrow olleH"
	if have != want {
		t.Error("func ReverseString don't work")
	}
}
