package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	fmt.Println("All arguments with program name:", args)
	fmt.Println("Arguments only:", args[1:])
	if len(args) != 5 {
		usage()
	}
	pathArgs := args[1:4]
	modulePath := filepath.Join(pathArgs...)

	fmt.Println(modulePath)
	fmt.Println(args[4])
}

func usage() {
	fmt.Println("\nERR: invalid args.\nUsage: tz <SERVICE> <MODULE> <ENV> <TF OPERATION>\n")
}
