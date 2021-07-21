package main_tests

import (
	"testing"
)

func Test_Smoke(t *testing.T) {
	if 2+2 != 4 {
		t.Error("Some Odd Math going on yea?")
	}
}

func Test_Api_is_running(t *testing.T) {

}
