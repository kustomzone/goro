package main

import (
	"context"
	"log"
	"os"

	"git.atonline.com/tristantech/gophp/core"
	_ "git.atonline.com/tristantech/gophp/ext/standard"
)

func main() {
	// by default, run script test.php
	p := core.NewProcess()
	ctx := core.NewGlobal(context.Background(), p)
	if err := ctx.RunFile("test.php"); err != nil {
		log.Printf("failed to run test file: %s", err)
		os.Exit(1)
	}
}
