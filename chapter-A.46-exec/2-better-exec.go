package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var output []byte
	var err error

	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", "git config user.name").Output()
	} else {
		output, err = exec.Command("bash", "-c", "git config user.name").Output()
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(output))
}
