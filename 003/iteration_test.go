package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, expected, repeated)
	})
	t.Run("repeat 'a' 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertCorrectMessage(t, expected, repeated)
	})
	t.Run("repeat 'a' -1 times", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := ""

		assertCorrectMessage(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 6)
	fmt.Println(repeat)
	// Output: aaaaaa
}

func assertCorrectMessage(t testing.TB, expected, repeated string) {
	t.Helper()
	if strings.Compare(expected, repeated) != 0 {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
