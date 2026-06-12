package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerPing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	HandlerPing(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	if string(body) != "pong" {
		t.Errorf("expected body 'pong', got '%s'", string(body))
	}
}

func TestHandlerHello(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		expectedStatus int
		expectedMsg    string
	}{
		{"valid name", "?name=Batman", http.StatusOK, "Batman"},
		{"missing name", "", http.StatusBadRequest, "name is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/hello"+tt.query, nil)
			w := httptest.NewRecorder()

			HandlerHello(w, req)

			res := w.Result()
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			body, _ := io.ReadAll(res.Body)
			if tt.expectedStatus == http.StatusOK {
				var result map[string]string
				json.Unmarshal(body, &result)
				if result["message"] != "Hello, "+tt.expectedMsg {
					t.Errorf("unexpected message: %s", result["message"])
				}
			}
		})
	}
}

func TestHandlerHelloWithServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HandlerHello)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello?name=Batman")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}
