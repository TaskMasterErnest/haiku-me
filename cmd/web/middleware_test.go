package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSecureHeaders(t *testing.T) {
	// initialize an httpResponseRecorder and a dummy http.Request
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	//a mock http handler that we can pass to the secureHeaders, containing a 200 response code and "OK" reponse body
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	//pass mock HTTP handler to the secureHeaders middleware
	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	//check that the middleware has correctly set the CSPs
	expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, rs.Header.Get("Content-Security-Policy"), expectedValue)

	// Check that the middleware has correctly set the Referrer-Policy
	expectedValue = "origin-when-cross-origin"
	assert.Equal(t, rs.Header.Get("Referrer-Policy"), expectedValue)

	// Check that the middleware has correctly set the X-Content-Type-Options
	expectedValue = "nosniff"
	assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), expectedValue)

	// Check that the middleware has correctly set the X-Frame-Options header
	expectedValue = "deny"
	assert.Equal(t, rs.Header.Get("X-Frame-Options"), expectedValue)

	// Check that the middleware has correctly set the X-XSS-Protection header
	expectedValue = "0"
	assert.Equal(t, rs.Header.Get("X-XSS-Protection"), expectedValue)

	// Check that the middleware has correctly called the next handler in line
	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")

}
