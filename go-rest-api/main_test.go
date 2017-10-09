package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPeople(t *testing.T) {
	// Create a request to pass to our handler. We don't ahve any query parameters for no so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/people", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	Router().ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":"1","firstname":"John","lastname":"Doe","address":{"city":"City X","state":"State X"}},{"id":"2","firstname":"Koko","lastname":"Doe","address":{"city":"City Y","state":"State Z"}},{"id":"3","firstname":"Francis","lastname":"Sunday"}]`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}
