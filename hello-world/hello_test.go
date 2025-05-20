package helloworld

import "testing"

func TestHello(T *testing.T) {
	T.Run("saying hello to people", func(t *testing.T) {
		name := "fulanito"
		got := Hello(name, "")
		want := EnglishHelloPrefix + name
		assertCorrectMessage(t, got, want)
	})
	T.Run("saying 'Hello World!' when empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := EnglishHelloPrefix + "World"
		assertCorrectMessage(t, got, want)
	})
	T.Run("in spanish", func(t *testing.T) {
		name := "Zutano"
		got := Hello(name, "spanish")
		want := SpanishHelloPrefix + name
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("the value %q is not equal to expected %q", got, want)
	}
}
