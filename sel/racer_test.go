package sel

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speed of servers, returning the fastest one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)
		defer func() {
			slowServer.Close()
			fastServer.Close()
		}()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("did not expected an error but got one, %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("return an error if server does not send a response withing 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(3 * time.Millisecond)
		fastServer := makeDelayedServer(3 * time.Millisecond)
		defer func() {
			slowServer.Close()
			fastServer.Close()
		}()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 1*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if delay > 0 {
			time.Sleep(delay)
		}
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
