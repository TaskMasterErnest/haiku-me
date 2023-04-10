package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	//initialize a new time.Time object and pass it to the humandate function
	tm := time.Date(2022, 4, 10, 10, 15, 0, 0, time.UTC)
	hd := humanDate(tm)

	//check that the output of humandat is what is expected
	//if not use the t.Errorf() function to indicate the test has failed and log the expected and actual results
	if hd != "10 Apr 2022 at 10:15" {
		t.Errorf("got %q; want %q", hd, "10 Apr 2022 at 10:15")
	}
}
