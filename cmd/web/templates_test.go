package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	//an anonymous struct containing the test case name,
	//input to the humanDate (tm field) and the expected output (the want field)
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 4, 10, 10, 15, 0, 0, time.UTC),
			want: "10 Apr 2023 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 4, 10, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "10 Apr 2023 at 09:15",
		},
	}

	//Loop over the test cases
	for _, tt := range tests {
		//the t.Run() is used to run a sub-test for each test case.
		//first param is the name used to identify the test, second param is the anon func containing the actual test
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("got %q; want %q", hd, tt.want)
			}
		})
	}
}
