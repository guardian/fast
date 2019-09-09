package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var confName = ".fast"

// LogLine - log line format
type LogLine struct {
	DateTime  string
	PerfScore float64
}

func getConfig() *os.File {
	f, err := os.OpenFile(confName, os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	return f
}

// TODO use reader/writer instead of direct access
func append(dt time.Time, report Lighthouse, w io.Writer) {
	// datetime perf-score
	dtFmt, _ := dt.MarshalText()
	line := LogLine{DateTime: string(dtFmt), PerfScore: report.Categories.Performance.Score}
	lineStr := fmt.Sprintf("%s %.2f", line.DateTime, line.PerfScore)

	_, err := w.Write([]byte(lineStr))
	check(err)
}

func read(r io.Reader, b []byte) int {
	n, err := r.Read(b)
	check(err)

	return n
}
