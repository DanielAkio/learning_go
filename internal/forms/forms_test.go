package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	f := New(r.PostForm)

	isValid := f.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows invalid when we have all fields")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("name", "Daniel Akio")
	postedData.Add("last_name", "T")
	r.PostForm = postedData

	form := New(r.PostForm)
	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows valid for a field that does not exist")
	}

	err_msg := form.Errors.Get("x")
	if err_msg == "" {
		t.Error("should have an error, but did not get one")
	}

	form = New(r.PostForm)
	form.MinLength("name", 6)
	if !form.Valid() {
		t.Error("form shows invalid when min length is valid")
	}

	err_msg = form.Errors.Get("name")
	if err_msg != "" {
		t.Error("should not have an error, but got one")
	}

	form = New(r.PostForm)
	form.MinLength("last_name", 6)
	if form.Valid() {
		t.Error("form shows valid when min length is invalid")
	}

	form.Errors.Get("name")
}

// Utilizar o comando "call" ao rodar metodos no debug console
func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("name", "Daniel Akio")
	postedData.Add("email", "danielakioteixeira@gmail.com")
	r.PostForm = postedData

	form := New(r.PostForm)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid for a field that does not exist")
	}

	form = New(r.PostForm)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows invalid when email is valid")
	}

	form = New(r.PostForm)
	form.IsEmail("name")
	if form.Valid() {
		t.Error("form shows valid when email is invalid")
	}
}
