package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"golang.org/x/term"
)

type CLI struct {
	Input        string `arg:"-i,--input"  help:"CSV content or path to .csv file"`
	PrettyFormat bool   `arg:"-p,--pretty" help:"Apply pretty format to the JSON output"`
}

func ReadCommandLine() (*CLI, error) {
	cli := CLI{}
	arg.MustParse(&cli)

	err := cli.Validate()
	if err != nil {
		return nil, err
	}

	return &cli, nil
}

func (cli *CLI) Validate() error {
	if cli.Input != "" {
		return nil
	}

	pipeInput, err := cli.readPipeInput()
	if err != nil {
		return err
	}

	if pipeInput == "" {
		return fmt.Errorf("no input provided")
	}

	cli.Input = pipeInput
	return nil
}

func (cli *CLI) readPipeInput() (string, error) {
	// Check if stdin is coming from a pipe/file or terminal.
	if term.IsTerminal(int(os.Stdin.Fd())) {
		return "", nil
	}

	input := ""

	// Read input from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input += "\n" + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("stdin read error: %w", err)
	}

	return input, nil
}
