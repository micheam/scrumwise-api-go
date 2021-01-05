package scrumwise

import (
	"sync"
	"testing"
)

func TestEndpoint_ordinal(t *testing.T) {
	got := Endpoint("addBacklogItem")
	want := BaseURL + "/" + ApiVersion + "/" + "addBacklogItem"
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

var mux sync.Mutex

func TestEndpoint_trim_slash(t *testing.T) {
	_BaseURL, _ApiVersion := BaseURL, ApiVersion
	mux.Lock()
	t.Cleanup(func() {
		BaseURL = _BaseURL
		ApiVersion = _ApiVersion
		mux.Unlock()
	})

	BaseURL = "https://api.scrumwise.com/service/api/"
	ApiVersion = "/v9/"
	got := Endpoint("/addBacklogItem")
	want := "https://api.scrumwise.com/service/api/v9/addBacklogItem"
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
