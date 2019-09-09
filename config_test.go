package main

import (
	"bytes"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	report := Lighthouse{}
	w := bytes.NewBufferString("")
	dt, _ := time.Parse(time.RFC3339, "2019-09-09T17:56:36.830633+01:00")
	append(dt, report, w)

	want := "2019-09-09T17:56:36.830633+01:00 0.00"
	got := w.String()

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
