package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {

	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(GetUsers)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"firstname":"rajat","lastname":"mishra","email":"rajat@gmail.com"},{"firstname":"John","lastname":"Doe","email":"JohnDoe@gmail.com"}]`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(GetUser)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `{"id":1,"firstname":"rajat","lastname":"mishra","email":"rajat@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestCreateUser(t *testing.T) {
	var jsonReq = []byte(`{"firstname":"rajat","lastname":"mishra","email":"rajat@gmail.com"}`)

	req, err := http.NewRequest("POST", "/Create", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()

	control := http.HandlerFunc(CreateUser)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}
	expected := `{"id":1,"firstname":"rajat","lastname":"mishra","email":"rajat@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestUpdateUser(t *testing.T) {

	var jsonReq = []byte(`{"id":2,"firstname":"John","lastname":"Doe","email":"JohnDoe@gmail.com"}`)

	req, err := http.NewRequest("PUT", "/update/{id}", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	control := http.HandlerFunc(UpdateUser)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":2,"firstname":"John","lastname":"Doe","email":"JohnDoe@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(DeleteUser)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `{"id":1,"firstname":"rajat","lastname":"mishra","email":"rajat@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}

}
