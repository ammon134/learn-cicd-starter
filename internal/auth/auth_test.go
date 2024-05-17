package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	httpReq := http.Request{Header: http.Header{}}
	httpReq.Header.Set("Authorization", "ApiKey fakeHash")
	got, err := GetAPIKey(httpReq.Header)
	if err != nil {
		t.Fatal(err)
	}
	want := "fakeHash"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %s, got: %s", want, got)
	}
}
