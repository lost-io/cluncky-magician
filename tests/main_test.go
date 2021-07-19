package main

import (
	"testing"
)

func Main_Smoke_test(t *testing.T) {
	if 2+2 != 4 {
		t.Error("Some Odd Math going on yea?")
	}
}
