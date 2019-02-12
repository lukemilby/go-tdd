package helloworld

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string){
		t.Helper()
		if got != want{
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello WORLD' when a + is supplied", func(t *testing.T){
		got := Hello("+", "English")
		want := "Hello WORLD"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Luke", "French" )
		want := "Bonjour Luke"
		assertCorrectMessage(t, got, want)
	})
}