package main

import (
	"bytes"
	"encoding/json"
	"go-http-demo/handlers"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	handlers.HelloHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	expected := "Hello, World! 👋\n"
	if string(body) != expected {
		t.Errorf("Expected %q but got %q", expected, string(body))
	}
}

func TestGreetHandler_WithoutBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/greet", nil)
	rec := httptest.NewRecorder()

	handlers.GreetHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	expected := "Invalid JSON\n"
	if string(body) != expected {
		t.Errorf("Expected %q but got %q", expected, string(body))
	}
}

func TestGreetHandler_WithBody(t *testing.T) {
	greetReq := handlers.GreetRequest{Name: "Shaun"}
	greetMarshall, _ := json.Marshal(greetReq)

	req := httptest.NewRequest(http.MethodPost, "/greet", bytes.NewReader(greetMarshall))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handlers.GreetHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}

	var got handlers.GreetResponse
	greetRes := handlers.GreetResponse{Message: "Hello, Shaun!"}
	err := json.NewDecoder(res.Body).Decode(&got)
	if err != nil {
		t.Fatal("failed to decode response: ", err)
	}

	if got != greetRes {
		t.Errorf("Expected %+v, got %+v", greetRes, got)
	}
}
