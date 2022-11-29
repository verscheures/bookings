package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)


func  TestForm_valid(t *testing.T){
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid{
		t.Error("got invalid when shoudl have been valid")
	}

}

func TestForm_Required(t *testing.T){
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "b")
	if form.Valid(){
		t.Error("got valid when shoudl have been invalid")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid(){
		t.Error("got invalid when shoudl have been valid")
	}
}

func TestHas(t *testing.T)()  {
	postedData := url.Values{}
	form := New(postedData)

	if form.Has("whatever"){
		t.Error("Has function not flagging missing field")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)
	if !form.Has("a",){
		t.Error("Has function should allow legal request")
	}

}

func TestMinLenght(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid(){
		t.Error("MinLength not failing on missing field")
	}
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should have an error but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("some_field", "some value")
	form = New(postedData)

	form.MinLength("some_field", 100)
	if form.Valid(){
		t.Error("Shows minlength of 100 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "sabc123")
	form = New(postedData)
	form.MinLength("another_field", 1)
	if !form.Valid(){
		t.Error("shows minlength of 1 is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("Should not have an error but got one")
	}
}


func TestForm_isEmail(t *testing.T){
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid(){
		t.Error("form shows valid email for non-existent vars")
	}


	postedData = url.Values{}
	postedData.Add("email", "abc@here.com")
	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid(){
		t.Error("form shows invalid email for correct email")
	}

	postedData = url.Values{}
	postedData.Add("email", "abc123")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid(){
		t.Error("form shows valid email for invalid email")
	}
}