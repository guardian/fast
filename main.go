package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Category - Type for top-level Lighthouse reports
type Category struct {
	Score float64 `json:"score"`
}

// Lighthouse output
type Lighthouse struct {
	Categories struct {
		Performance Category `json:"performance"`
	} `json:"categories"`
}

func main() {
	targetURL := flag.String("target-url", "", "a target URL to run Lighthouse against")
	startCmd := flag.String("start-cmd", "", "command(s) to start your service")
	stopCmd := flag.String("stop-cmd", "", "command(s) to shutdown your service")

	flag.Parse()

	if *targetURL == "" {
		log.Print("Missing or invalid arguments.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	runCmd(*startCmd)
	branchOut := runLighthouse(*targetURL)
	runCmd(*stopCmd)

	runCmd("git checkout master")

	runCmd(*startCmd)
	masterOut := runLighthouse(*targetURL)
	runCmd(*stopCmd)

	// TODO
	// - report against budgets (if present)
	// - run x times and average
	// - compare FCP, TTI, bundle size
	fmt.Printf(
		"Branch (%.2f), Master (%.2f)",
		branchOut.Categories.Performance.Score,
		masterOut.Categories.Performance.Score,
	)
}

func (lh *Lighthouse) unmarshal(data []byte) error {
	return json.Unmarshal(data, lh)
}

func checkCmd(cmd string, data []byte, err error) {
	if err != nil {
		log.Printf("Command '%s' failed. Output was:\n%v", cmd, data)
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runLighthouse(targetURL string) Lighthouse {
	var lh Lighthouse
	data := runCmd(fmt.Sprintf("lighthouse %s --output json", targetURL))
	check(lh.unmarshal(data))
	return lh
}

// Warning exits early on failure
func runCmd(cmd string) []byte {
	data, err := exec.Command("/bin/sh", "-c", cmd).Output()
	checkCmd(cmd, data, err)
	return data
}
