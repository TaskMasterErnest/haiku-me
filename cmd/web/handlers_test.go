package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/TaskMasterErnest/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")

}

func TestSnippetView(t *testing.T) {
	// Create a new instance of our application struct which uses the mocked dependencies.
	app := newTestApplication(t)

	// Establish a new test server for running end-to-end tests.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Set up some table-driven tests to check the responses sent by our
	// application for different URLs.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond...",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			assert.Equal(t, code, tt.wantCode)

			if tt.wantBody != "" {
				assert.StringContains(t, body, tt.wantBody)
			}
		})
	}
}

func TestUserSignup(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName     = "Bob"
		validPassword = "validPa$$word"
		validEmail    = "bob@example.com"
		formTag       = "<form action='/user/signup' method='POST' novalidate>"
	)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantFormTag  string
	}{
		{
			name:         "Valid submission",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusSeeOther,
		},
		{
			name:         "Invalid CSRF Token",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    "wrongToken",
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "Empty name",
			userName:     "",
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty email",
			userName:     validName,
			userEmail:    "",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty password",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: "",
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Invalid email",
			userName:     validName,
			userEmail:    "bob@example.",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Short password",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: "pa$$",
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Duplicate email",
			userName:     validName,
			userEmail:    "dupe@example.com",
			userPassword: validPassword,
			csrfToken:    validCSRFToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)

			assert.Equal(t, code, tt.wantCode)

			if tt.wantFormTag != "" {
				assert.StringContains(t, body, tt.wantFormTag)
			}
		})
	}
}

// TestPing tests ping handler for the correct response status code, 200 and
// the correct response body, "OK".
// func TestPing(t *testing.T) {
// 	t.Parallel()

// 	app := newTestApplication(t)
// 	ts := newTestServer(t, app.routes())
// 	defer ts.Close()

// 	code, _, body := ts.get(t, "/healthcheck")

// 	if code != http.StatusOK {
// 		t.Errorf("want %d; got %d", http.StatusOK, code)
// 	}

// 	if string(body) != "OK" {
// 		t.Errorf("want body to equal %q", "OK")
// 	}
// }

// func TestShowSnippet(t *testing.T) {
// 	t.Parallel()
// 	// Create a new instance of our application struct which uses the mocked
// 	// dependencies.
// 	app := newTestApplication(t)

// 	// Establish a new test server for running end-to-end tests.
// 	ts := newTestServer(t, app.routes())
// 	defer ts.Close()

// 	// Set up some table-driven tests to check the responses sent by our
// 	// application for different URLS
// 	tests := []struct {
// 		name     string
// 		urlPath  string
// 		wantCode int
// 		wantBody []byte
// 	}{
// 		{"Valid ID", "/snippet/1", http.StatusOK, []byte("An old silent pond...")},
// 		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
// 		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
// 		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
// 		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
// 		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
// 		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
// 	}

// 	for _, tt := range tests {
// 		// rebind tt into this lexical scope to avoid concurrency bug from running
// 		// sub-tests
// 		// tt := tt

// 		t.Run(tt.name, func(t *testing.T) {
// 			// t.Parallel() // can't run in parallel with current testServer implementation

// 			code, _, body := ts.get(t, tt.urlPath)
// 			t.Logf("testing %q for want-code %d and want-body %q", tt.name, tt.wantCode,
// 				tt.wantBody)

// 			if code != tt.wantCode {
// 				t.Errorf("want %d; got %d", tt.wantCode, code)
// 			}

// 			if !bytes.Contains(body, tt.wantBody) {
// 				t.Errorf("want body to contain %q, but got %q", tt.wantBody, body)
// 			}
// 		})
// 	}

// }

// // TestSignupUser tests that signupUser handler returns appropriate status codes and error messages
// // corresponding logic of signupUser handler.
// func TestSignupUser(t *testing.T) {
// 	t.Parallel()
// 	// Create the application struct containing our mocked dependencies and
// 	// set up the test server for running an end-to-test.
// 	app := newTestApplication(t)
// 	ts := newTestServer(t, app.routes())
// 	defer ts.Close()

// 	// Make a GET /user/signup request and then extract the CSRF token from the
// 	// response body.
// 	_, _, body := ts.get(t, "/user/signup")
// 	csrfToken := extractCSRFToken(t, body)

// 	tests := []struct {
// 		name         string
// 		userName     string
// 		userEmail    string
// 		userPassword string
// 		csrfToken    string
// 		wantCode     int
// 		wantBody     []byte
// 	}{
// 		{"Valid submission", "Bob", "bob@example.com", "validPa$$word", csrfToken,
// 			http.StatusSeeOther, nil},
// 		{"Empty name", "", "bob@example.com", "validPa$$word", csrfToken, http.StatusOK,
// 			[]byte("This field cannot be blank")},
// 		{"Empty email", "Bob", "", "validPa$$word", csrfToken, http.StatusOK,
// 			[]byte("This field cannot be blank")},
// 		{"Empty password", "Bob", "bob@example.com", "", csrfToken, http.StatusOK,
// 			[]byte("This field cannot be blank")},
// 		{"Invalid email (incomplete domain)", "Bob", "bob@example.", "validPa$$word",
// 			csrfToken, http.StatusOK, []byte("This field is invalid")},
// 		{"Invalid email (missing @)", "Bob", "bobexample.com", "validPa$$word", csrfToken,
// 			http.StatusOK, []byte("This field is invalid")},
// 		{"Invalid email (missing local part)", "Bob", "@example.com", "validPa$$word",
// 			csrfToken, http.StatusOK, []byte("This field is invalid")},
// 		{"Short password", "Bob", "bob@example.com", "pa$$word", csrfToken, http.StatusOK,
// 			[]byte("This field is too short (minimum is 10 characters")},
// 		{"Duplicate email", "Bob", "dupe@example.com", "validPa$$word", csrfToken, http.StatusOK,
// 			[]byte("Address is already in use")},
// 		{"Invalid CSRF Token", "", "", "", "wrongToken", http.StatusBadRequest, nil},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			form := url.Values{}
// 			form.Add("name", tt.userName)
// 			form.Add("email", tt.userEmail)
// 			form.Add("password", tt.userPassword)
// 			form.Add("csrf_token", tt.csrfToken)

// 			code, _, body := ts.postForm(t, "/user/signup", form)
// 			t.Logf("testing %q for want-code %d and want-body %q", tt.name, tt.wantCode,
// 				tt.wantBody)

// 			if code != tt.wantCode {
// 				t.Errorf("want %d; got %d", tt.wantCode, code)
// 			}

// 			if !bytes.Contains(body, tt.wantBody) {
// 				t.Errorf("want body to contain %q, but got %q", tt.wantBody, body)
// 			}
// 		})
// 	}
// }
