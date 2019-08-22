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
	branchOut, err := runCmd("lighthouse https://www.theguardian.com/society/2019/aug/22/deaths-on-the-rise-in-10-of-britains-toughest-prisons?dcr --output json")
	if err != nil {
		log.Fatal(err)
	}

	_, err = runCmd("git checkout master")
	if err != nil {
		log.Fatal(err)
	}

	masterOut, err := runCmd("lighthouse https://www.theguardian.com/society/2019/aug/22/deaths-on-the-rise-in-10-of-britains-toughest-prisons?dcr --output json")
	if err != nil {
		log.Fatal(err)
	}

	var masterLh, branchLh Lighthouse
	quickFail(masterLh.unmarshal(masterOut))
	quickFail(branchLh.unmarshal(branchOut))

	fmt.Printf("Branch (%.2f), Master (%.2f)", branchLh.Categories.Performance.Score, masterLh.Categories.Performance.Score)
}

func (lh *Lighthouse) unmarshal(data []byte) error {
	return json.Unmarshal(data, lh)
}

func quickFail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runCmd(cmd string) ([]byte, error) {
	args := strings.Split(cmd, " ")
	log.Printf("%v", args)
	c := exec.Command(args[0], args[1:]...)

	return c.Output()
}
