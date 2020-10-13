package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestFindAll(t *testing.T) {

	body := map[string][]map[string]interface{}{
		"data": {
			{"ID": float64(1), "name": "abhishek", "email": "abhi@gmail.com"},
			{"ID": float64(2), "name": "ajay", "email": "ajay@gmail.com"},
		},
	}
	// Grab our router
	server := SetupRouter()

	// Perform a GET request with that handler.
	w := performRequest(server, "GET", "/user")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	// var response map[string]string
	var response map[string][]map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	value, exists := response["data"]
	// Make some assertions on the correctness of the response.

	fmt.Println(value)
	fmt.Println(exists)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["data"], response["data"], "not equal")
}

func TestGetSingleUser(t *testing.T) {

	body := map[string]map[string]interface{}{
		"data": {"ID": float64(1), "name": "abhishek", "email": "abhi@gmail.com"},
	}
	// Grab our router
	server := SetupRouter()

	// Perform a GET request with that handler.
	w := performRequest(server, "GET", "/user/1")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the JSON response to a map
	// var response map[string]string
	var response map[string]map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	value, exists := response["data"]
	// Make some assertions on the correctness of the response.

	fmt.Println(value)
	fmt.Println(exists)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["data"], response["data"], "not equal")
}
