package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)
/*
	Test For Handler Functions
 */

func TestSet(t *testing.T) {
	requestBody := map[string]string{
		"key":   "key_test",
		"value": "value_test",
	}
	bodyData, err := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/set", bytes.NewBuffer(bodyData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Set)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseMap map[string]string

	if err := json.Unmarshal([]byte(rr.Body.String()), &responseMap); err != nil {
		panic(err)
	}
	if reflect.DeepEqual(requestBody, responseMap) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseMap, requestBody)
	}
}



func TestGet(t *testing.T) {

	req, err := http.NewRequest("GET", fmt.Sprintf("/get?key=%s","key_test"),nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseMap map[string]string

	 expectedBody := map[string]string{
		"key":   "key_test",
		"value": "value_test",
	}
	if err := json.Unmarshal([]byte(rr.Body.String()), &responseMap); err != nil {
		panic(err)
	}
	if reflect.DeepEqual(expectedBody, responseMap) != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseMap, expectedBody)
	}
}

func TestSetMissingParams(t *testing.T) {
	requestBody := map[string]string{
		"key":   "key_test",
	}
	bodyData, err := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/set", bytes.NewBuffer(bodyData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Set)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestGetWrongParam(t *testing.T) {

	req, err := http.NewRequest("GET", fmt.Sprintf("/get?kesy=%s","key_test"),nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestGetNonExistsKey(t *testing.T) {

	req, err := http.NewRequest("GET", fmt.Sprintf("/get?key=%s","non_exist_key"),nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

}

func TestFlush(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/flush",nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

}