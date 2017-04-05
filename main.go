package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/topherbullock/steganogify/cmd"
)

func main() {
	args := os.Args[1:]

	var command cmd.Command

	switch args[0] {
	case "decode":
		command = &cmd.Decode{}
	case "encode":
		command = &cmd.Encode{}
	}

	if command == nil {
		os.Exit(1)
	}

	parser := flags.NewParser(command, flags.Default)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}
	command.Run()
}
