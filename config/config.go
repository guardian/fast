package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/guardian/fast/lighthouse"
)

var confName = ".fast"

// LogLine - log line format
type LogLine struct {
	Branch    string
	DateTime  string
	PerfScore float64
	JSBytes   float64
	TTI       float64
}

func Exists() bool {
	_, err := os.Stat(confName)
	return err == nil
}

func Get() *os.File {
	f, err := os.OpenFile(confName, os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	return f
}

func Format(dt time.Time, branch string, report lighthouse.Lighthouse) string {
	// datetime perf-score
	dtFmt := dt.UTC().Format(time.RFC3339)

	var jsBytes float64
	for _, item := range report.Audits.ResourceSummary.Details.Items {
		if item.ResourceType == "script" {
			jsBytes = item.Size
		}
	}

	line := LogLine{
		Branch:    branch,
		DateTime:  string(dtFmt),
		PerfScore: report.Categories.Performance.Score,
		JSBytes:   jsBytes,
		TTI:       report.Audits.Interactive.NumericValue,
	}

	return fmt.Sprintf(
		"%-10s %s %-4.2f %-5.2f %-8.f\n",
		line.Branch[:9],
		line.DateTime,
		line.PerfScore,
		line.TTI/1000, // convert to seconds
		line.JSBytes,
	)
}

// TODO use reader/writer instead of direct access
func Append(dt time.Time, branch string, report lighthouse.Lighthouse, w io.Writer) {
	lineStr := Format(dt, branch, report)
	_, err := w.Write([]byte(lineStr))
	check(err)
}

func Create() {
	header := fmt.Sprint("Format: [branch] [timestamp] [perf-score] [js size] [tti]\n\n")

	conf, err := os.Create(confName)
	check(err)

	_, err = conf.Write([]byte(header))
	check(err)
}

func read(r io.Reader, b []byte) int {
	n, err := r.Read(b)
	check(err)

	return n
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
