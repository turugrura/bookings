package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	isValid := form.Valid()
	if !isValid {
		t.Error("got valid when should have been valid")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.Has("a")
	if form.Valid() {
		t.Error("got valid when has no specific field")
	}

	isErr := form.Errors.Get("a")
	if isErr == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("a", "abcde")

	form = New(postedData)
	form.Has("a")
	if !form.Valid() {
		t.Error("form shows invalid when has specific field")
	}

	isErr = form.Errors.Get("a")
	if isErr != "" {
		t.Error("should not have an error, but get one")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form shows invalid when required fields not missing")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "abcd")
	form := New(postedData)

	form.MinLength("a", 5)
	if form.Valid() {
		t.Error("form shows valid when minlength")
	}

	postedData = url.Values{}
	postedData.Add("a", "abcde")

	form = New(postedData)
	form.MinLength("a", 5)
	if !form.Valid() {
		t.Error("form shows invalid when minlength fields")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("form shows valid when wrong email")
	}

	postedData = url.Values{}
	postedData.Add("a", "abcde@a.a")

	form = New(postedData)
	form.IsEmail("a")
	if !form.Valid() {
		t.Error("form shows invalid when correct email")
	}
}
