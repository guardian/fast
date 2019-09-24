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
	dtFmt, _ := dt.MarshalText()
	line := LogLine{Branch: branch, DateTime: string(dtFmt), PerfScore: report.Categories.Performance.Score}
	return fmt.Sprintf("%-10s %s %.2f\n", line.Branch, line.DateTime, line.PerfScore)
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
