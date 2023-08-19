package main

import (
	"github.com/rzeAkbari/GoGo/gin/cmd"
	"os"
)

func main() {
	err := cmd.Run()

	if err != nil {
		os.Exit(1)
	}
}
