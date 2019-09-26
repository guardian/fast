package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/guardian/fast/config"
	"github.com/guardian/fast/lighthouse"
)

func main() {
	targetURL := flag.String("target-url", "", "a target URL to run Lighthouse against")
	startCmd := flag.String("start-cmd", "", "command(s) to start your service")
	stopCmd := flag.String("stop-cmd", "", "command(s) to shutdown your service")
	append := flag.Bool("append", false, "if true, appends result to .fast file")

	flag.Parse()

	if *targetURL == "" {
		log.Print("Missing or invalid arguments.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	ensureWdClean()
	ensureHasLighthouse()

	branch := strings.Replace(string(runCmd("git symbolic-ref --short -q HEAD")), "\n", "", -1)

	runCmd(*startCmd)
	res1 := runLighthouse(*targetURL)
	res2 := runLighthouse(*targetURL)
	res3 := runLighthouse(*targetURL)
	runCmd(*stopCmd)

	average := merge(res1, res2, res3)

	if *append {
		if !config.Exists() {
			config.Create()
		}

		conf := config.Get()
		config.Append(time.Now(), branch, average, conf)
	}

	fmt.Println(config.Header())
	fmt.Println(config.Format(time.Now(), branch, average))
}

func merge(a, b, c lighthouse.Lighthouse) lighthouse.Lighthouse {
	// basic merge as we only care about a few values
	avgPerfScore := (a.Categories.Performance.Score + b.Categories.Performance.Score + c.Categories.Performance.Score) / 3

	a.Categories.Performance.Score = avgPerfScore
	return a
}

func checkCmd(cmd string, data []byte, err error) {
	if err != nil {
		log.Printf("Command '%s' failed. Output was:\n%s", cmd, data)
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runLighthouse(targetURL string) lighthouse.Lighthouse {
	var lh lighthouse.Lighthouse
	data := runCmd(fmt.Sprintf("lighthouse %s --output json", targetURL))
	check(lh.Unmarshal(data))
	return lh
}

func ensureWdClean() {
	data := runCmd("git status --porcelain")
	if len(data) > 0 {
		log.Fatal("Error, working directory is not clean.")
	}
}

func ensureHasLighthouse() {
	cmd := "command -v lighthouse"
	err := exec.Command("/bin/sh", "-c", cmd).Run()
	if err != nil {
		log.Fatal("Error, Lighthouse is not installed (use 'npm install -g lighthouse' using latest LTS Node version).")
	}
}

// Warning exits early on failure
func runCmd(cmd string) []byte {
	data, err := exec.Command("/bin/sh", "-c", cmd).Output()
	checkCmd(cmd, data, err)
	return data
}
