package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	HelloWorld(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Fatalf("handler rturned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if b := w.Body.String(); b != "Hello World!" {
		t.Fatalf("body = %s, want no", b)
	}

}
