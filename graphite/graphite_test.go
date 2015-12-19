package graphite

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGraphiteComesWithTimeOut(t *testing.T) {
	graphite := New("", "", "example.com")
	if graphite.Client.Timeout == 0 {
		t.Fatal("Needs to have a default timeout")
	}
}

func TestGraphiteReturnsErrorOnNon200(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("Die"), 589)
	}))
	defer ts.Close()

	graphite := New("", "", ts.URL)
	err := graphite.Publish(NewEvent("What", "Dat", "tag1", "tag2"))

	if err == nil {
		t.Fatal("Should have errored")
	}
}

func TestGraphiteSendsAuthWhenSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Auth Req", 401)
		}
	}))
	defer ts.Close()

	graphite := New("", "", ts.URL)
	err := graphite.Publish(NewEvent("What", "Dat", "tag1", "tag2"))

	if err == nil {
		t.Fatal("Should have not authed")
	}

	graphite = New("foo", "bar", ts.URL)
	err = graphite.Publish(NewEvent("What", "Dat", "tag1", "tag2"))

	if err != nil {
		t.Fatal(err)
	}
}

func TestGraphiteSendsEvents(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 527)
		}

		var event *Event
		if err == nil {
			err = json.Unmarshal(body, &event)
			if err != nil {
				http.Error(w, err.Error(), 528)
				return
			}
		}

		if err == nil {
			if event.What != "What" && event.Data != "Dat" && event.Tags != "tag1,tag2" {
				http.Error(w, fmt.Sprintf("%s", body), 529)
				return
			}
		}
	}))
	defer ts.Close()

	graphite := New("", "", ts.URL)
	err := graphite.Publish(NewEvent("What", "Dat", "tag1", "tag2"))

	if err != nil {
		t.Fatal(err)
	}

}
