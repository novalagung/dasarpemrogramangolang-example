package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "10")
	err := cmd.Run()
	if err != nil {
		fmt.Println("error:", err)
		// error: signal: killed (jika timeout tercapai)
	}
}
