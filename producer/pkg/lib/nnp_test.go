package lib_test

import (
	"net/http"
	"testing"
)

func TestProducer(t *testing.T) {
	t.Run("pushes notification", func(t *testing.T) {

	})
}

func pushNotificationHandlerFor(t testing.TB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
