package config

import (
	"bytes"
	"testing"
	"time"

	"github.com/guardian/fast/lighthouse"
)

func TestAbs(t *testing.T) {
	report := lighthouse.Lighthouse{}
	w := bytes.NewBufferString("")
	dt, _ := time.Parse(time.RFC3339, "2019-09-09T17:56:36.830633+01:00")
	branch := "master"
	Append(dt, branch, report, w)

	want := "master 2019-09-09T17:56:36.830633+01:00 0.00"
	got := w.String()

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
