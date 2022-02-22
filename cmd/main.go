package main

import (
	"fmt"
	"github.com/gobuild-tu/analyze"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("incompatible number of arguments: expected: 1 actual: ", len(os.Args)-1)
		return
	}

	mainPath := os.Args[1]
	fmt.Println("Build path: ", mainPath)

	outputFile := fmt.Sprintf("/tmp/compile_%d.json", time.Now().UnixNano())
	goBuild := fmt.Sprintf("go build -a -debug-actiongraph=%s %s", outputFile, mainPath)
	fmt.Println("Executing: ", goBuild)
	cmd := exec.Command("bash", "-c", goBuild)

	if err := cmd.Start(); err != nil {
		fmt.Println("build failed: ", err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("build failed: ", err)
		return
	}

	analyze.JSONFile(outputFile)
}
