package main

import (
	"flag"
	"fmt"
	"os"
)

func Encrypt(plaintext, key string) {

}

func Decrypt() {

}

func main() {
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println(fmt.Errorf("missing required arguments"))
		os.Exit(1)
	}

}
