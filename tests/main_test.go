package main_tests

import (
	"testing"
)

func Test_Smoke(t *testing.T) {
	if 2+2 != 4 {
		t.Error("Some Odd Math going on yea?")
	}
}

/*
func Test_Api_is_running(t *testing.T) {
	req, err := http.NewRequest("GET", "/entries", nil)
	if err != nil {
		t.Fatal(err)
	}
}
*/
func Test_Api_returns_card(t *testing.T) {

}
