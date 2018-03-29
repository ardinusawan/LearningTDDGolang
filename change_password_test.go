package main

import (
	"testing"
	// "fmt"
	"net/http"
  	"net/http/httptest"
)

type testWebRequest struct {

}

func (testWebRequest) FetchBytes(url string) []byte {
	return []byte(`{"username": "ehe"}`)
}


func (testWebRequest) SendData(url string, data string) []byte {

	return []byte(`{"status":true, "username": "Ardi Nusawan"}`)
}

// func TestGetChangePasswordPage(t *testing.T){
// 	// status := GetChangePassword(LiveGetChangePasswordRequest{})
// 	status := GetChangePassword(testWebRequest{}, "")
// 	if status == "" {
// 		t.Errorf("Data is nil, want username: string")
// 	}
// }

func TestGetWrongResponseStatus(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusServiceUnavailable)
  }))
  defer ts.Close()
  serverUrl := ts.URL
  _ , err := GetChangePassword(LiveGetChangePasswordRequest{}, serverUrl)
  if err == nil {
  	t.Errorf("Should be error!: %s", err)
  }

}

func TestGetOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  	w.WriteHeader(http.StatusOK)
	  	w.Write([]byte(`{"status":true, "message":{"username": "Ardi Nusawan"}}`))
		if r.Method != "GET" {
		      t.Errorf("Expected ‘GET’ request, got ‘%s’", r.Method)
		    }
		    // if r.URL.EscapedPath() != "/pub" {
		    //   t.Errorf("Expected request to ‘/pub’, got ‘%s’", r.URL.EscapedPath())
		    // }
		    // r.ParseForm()
		    // topic := r.Form.Get("topic")
		    // if topic != "meaningful-topic" {
		    //   t.Errorf("Expected request to have ‘topic=meaningful-topic’, got: ‘%s’", topic)
		    // }
	  }))
	defer ts.Close()
	serverUrl := ts.URL
	_ , err := GetChangePassword(LiveGetChangePasswordRequest{}, serverUrl)
	if err != nil {
		t.Errorf("Publish() returned an error: %s", err)
	}
}