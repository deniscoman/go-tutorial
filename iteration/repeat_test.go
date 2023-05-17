package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeted := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeted != expected {
		t.Errorf("expected %q, but got %q", expected, repeted)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 10)
	fmt.Println(repeat)
	// Output: aaaaaaaaaa
}
