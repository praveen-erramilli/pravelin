package events

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pravelin/response"
	"reflect"
	"strings"
	"testing"
)

func TestEventRequestHandlerForPostEventsRequest(t *testing.T) {
	//given
	req, err := http.NewRequest("POST", "/api/v1/events",
		strings.NewReader("{\n  \"eventType\": \"copyAndPaste\",\n  \"websiteUrl\": \"https://ravelin.com\",\n  \"sessionId\": \"123123-123123-123123123\",\n  \"pasted\": true,\n  \"formId\": \"inputCardNumber\"\n}"))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	store := NewEventStore()
	handler := http.HandlerFunc(store.EventRequestHandler)

	//when
	handler.ServeHTTP(rr, req)

	//then
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var result response.Response
	json.NewDecoder(rr.Body).Decode(&result)

	expectedResult := response.Response{
		Status:     "success",
		StatusCode: 200,
		Data:       "Successfully processed the event",
	}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("handler returned wrong response: got %v want %v",
			result, expectedResult)
	}

	if origin := rr.Header().Get("Access-Control-Allow-Origin"); "*" != origin {
		t.Errorf("handler returned wrong header: got %v want %v",
			origin, "*")
	}
	if allowedHeaders := rr.Header().Get("Access-Control-Allow-Headers"); "*" != allowedHeaders {
		t.Errorf("handler returned wrong header: got %v want %v",
			allowedHeaders, "*")
	}
	if allowedMethods := rr.Header().Get("Access-Control-Allow-Methods"); "POST, OPTIONS" != allowedMethods {
		t.Errorf("handler returned wrong header: got %v want %v",
			allowedMethods, "POST, OPTIONS")
	}
	if contentType := rr.Header().Get("Content-Type"); "application/json; charset=utf-8" != contentType {
		t.Errorf("handler returned wrong header: got %v want %v",
			contentType, "application/json; charset=utf-8")
	}
}

func TestEventRequestHandlerForOptionsEventsRequest(t *testing.T) {
	//given
	req, err := http.NewRequest("OPTIONS", "/api/v1/events", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	store := NewEventStore()
	handler := http.HandlerFunc(store.EventRequestHandler)

	//when
	handler.ServeHTTP(rr, req)

	//then
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if origin := rr.Header().Get("Access-Control-Allow-Origin"); "*" != origin {
		t.Errorf("handler returned wrong header: got %v want %v",
			origin, "*")
	}
	if allowedHeaders := rr.Header().Get("Access-Control-Allow-Headers"); "*" != allowedHeaders {
		t.Errorf("handler returned wrong header: got %v want %v",
			allowedHeaders, "*")
	}
	if allowedMethods := rr.Header().Get("Access-Control-Allow-Methods"); "POST, OPTIONS" != allowedMethods {
		t.Errorf("handler returned wrong header: got %v want %v",
			allowedMethods, "POST, OPTIONS")
	}
	if contentType := rr.Header().Get("Content-Type"); "application/json; charset=utf-8" != contentType {
		t.Errorf("handler returned wrong header: got %v want %v",
			contentType, "application/json; charset=utf-8")
	}
}

func TestEventRequestHandlerForUnsupportedEventsRequest(t *testing.T) {
	//given
	req, err := http.NewRequest("PATCH", "/api/v1/events", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	store := NewEventStore()
	handler := http.HandlerFunc(store.EventRequestHandler)

	//when
	handler.ServeHTTP(rr, req)

	//then
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestEventRequestHandlerForInvalidEventsInput(t *testing.T) {
	//given
	req, err := http.NewRequest("POST", "/api/v1/events",
		strings.NewReader("{\n  \"eventType\": \"copyAndPaste\",\n  \"websiteUrl\": \"https://ravelin.com\",\n  \"sessionId\": \"123123-123123-123123123\",\n  \"pasted\": true,\n  \"formId\": \"testButton\"\n}"))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	store := NewEventStore()
	handler := http.HandlerFunc(store.EventRequestHandler)

	//when
	handler.ServeHTTP(rr, req)

	//then
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	var result response.Response
	json.NewDecoder(rr.Body).Decode(&result)

	expectedResult := response.Response{
		Status:     "failure",
		StatusCode: 400,
		Data:       "invalid value specified for formId",
	}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("handler returned wrong response: got %v want %v",
			result, expectedResult)
	}

	if origin := rr.Header().Get("Access-Control-Allow-Origin"); "*" != origin {
		t.Errorf("handler returned wrong header: got %v want %v",
			origin, "*")
	}
	if allowedHeaders := rr.Header().Get("Access-Control-Allow-Headers"); "*" != allowedHeaders {
		t.Errorf("handler returned wrong header: got %v want %v",
			allowedHeaders, "*")
	}
	if allowedMethods := rr.Header().Get("Access-Control-Allow-Methods"); "POST, OPTIONS" != allowedMethods {
		t.Errorf("handler returned wrong header: got %v want %v",
			allowedMethods, "POST, OPTIONS")
	}
	if contentType := rr.Header().Get("Content-Type"); "application/json; charset=utf-8" != contentType {
		t.Errorf("handler returned wrong header: got %v want %v",
			contentType, "application/json; charset=utf-8")
	}
}
