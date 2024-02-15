package main

import (
	"fmt"
	"os"
)

func main() {
	cli, err := ReadCommandLine()
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}

	csv, err := ParseCSV(cli.Input)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}

	jsonStr, err := CastToJSON(csv, cli.PrettyFormat)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}

	fmt.Println(jsonStr)
}
