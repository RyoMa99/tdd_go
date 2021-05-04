package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat 5 count", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("repeat 9 count", func(t *testing.T) {
		repeated := Repeat("a", 9)
		expected := "aaaaaaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
