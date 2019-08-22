package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
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
	out, err := runCmd("lighthouse https://www.theguardian.com/society/2019/aug/22/deaths-on-the-rise-in-10-of-britains-toughest-prisons?dcr --output json")
	if err != nil {
		log.Fatal(err)
	}

	var lh Lighthouse
	err = json.Unmarshal(out, &lh)
	fmt.Printf("Performance score is %v\n", lh.Categories.Performance.Score)
}

func runCmd(cmd string) ([]byte, error) {
	args := strings.Split(cmd, " ")
	c := exec.Command(args[0], args[1:]...)

	return c.Output()
}
