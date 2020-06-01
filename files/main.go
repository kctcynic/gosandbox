package main

import (
	"fmt"
	"os"
)

func main() {
	for _, path := range os.Args {
		fmt.Printf("%s becomes %s\n", path, os.ExpandEnv(path))
	}
}
