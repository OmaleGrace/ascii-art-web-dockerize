package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name   string
		method string
		target string
		code   int
	}{
		{
			name:   "valid request",
			method: http.MethodGet,
			target: "/",
			code:   http.StatusOK,
		},
		{
			name:   "wrong method",
			method: http.MethodPost,
			target: "/",
			code:   http.StatusMethodNotAllowed,
		},
		{
			name:   "wrong path",
			method: http.MethodGet,
			target: "/wrong",
			code:   http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.target, nil)
			rec := httptest.NewRecorder()

			HomeHandler(rec, req)

			if rec.Code != tt.code {
				t.Errorf("got %d want %d", rec.Code, tt.code)
			}
		})
	}
}

func TestArtHandler(t *testing.T) {
	tests := []struct {
		name   string
		method string
		text   string
		banner string
		code   int
	}{
		{
			name:   "valid request",
			method: http.MethodPost,
			text:   "A",
			banner: "standard",
			code:   http.StatusOK,
		},
		{
			name:   "wrong method",
			method: http.MethodGet,
			text:   "A",
			banner: "standard",
			code:   http.StatusMethodNotAllowed,
		},
		{
			name:   "empty text",
			method: http.MethodPost,
			text:   "",
			banner: "standard",
			code:   http.StatusBadRequest,
		},
		{
			name:   "missing banner",
			method: http.MethodPost,
			text:   "A",
			banner: "unknown",
			code:   http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			form := url.Values{}
			form.Set("text", tt.text)
			form.Set("banner", tt.banner)

			req := httptest.NewRequest(
				tt.method,
				"/ascii-art",
				strings.NewReader(form.Encode()),
			)

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			rec := httptest.NewRecorder()

			ArtHandler(rec, req)

			if rec.Code != tt.code {
				t.Errorf("got %d want %d", rec.Code, tt.code)
			}
		})
	}
}

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		name    string
		message string
		code    int
	}{
		{
			name:    "bad request",
			message: "Bad Request",
			code:    http.StatusBadRequest,
		},
		{
			name:    "not found",
			message: "Not Found",
			code:    http.StatusNotFound,
		},
		{
			name:    "method not allowed",
			message: "Wrong Method",
			code:    http.StatusMethodNotAllowed,
		},
		{
			name:    "internal server error",
			message: "Internal Server Error",
			code:    http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			_ = req

			rec := httptest.NewRecorder()

			ErrorHandler(rec, tt.message, tt.code)

			if rec.Code != tt.code {
				t.Errorf("got %d want %d", rec.Code, tt.code)
			}

			if !strings.Contains(rec.Body.String(), tt.message) {
				t.Errorf("response should contain %q", tt.message)
			}
		})
	}
}