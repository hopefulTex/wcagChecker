package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
)

const version string = "v1.0.0"
const helpString string = "compare two colors\nusage:\twcagChecker #hexColor #hexColor\n--file\tcompare base color to those in a file\nusage: --file filename #hexColor"

const (
	DUO_MODE = iota
	LIST_MODE
	// GENERATE_MODE
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("error: not enough arguments.\n")
		fmt.Println(helpString)
		os.Exit(1)
	}

	var mode int = DUO_MODE
	var output string
	var err error

	switch os.Args[1] {
	case "--version":
		fmt.Println(version)
		os.Exit(1)
	case "--help":
		fmt.Println(helpString)
		os.Exit(1)
	case "--file":
		if len(os.Args) < 4 {
			fmt.Print("error: no file or base color specified.\n")
			fmt.Print("(want --file filename #hexcolor)\n")
			os.Exit(1)
		}
		mode = LIST_MODE
	}

	// Output to terminal or to file/pipe?
	isTTY := isatty.IsTerminal(os.Stdout.Fd())

	switch mode {
	case DUO_MODE:
		output, err = duo(isTTY, os.Args[1], os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			fmt.Println(output)
			os.Exit(1)
		}
	case LIST_MODE:
		if isTTY {
			argIndex := 3
			// give a chance to ctrl+c out of a long list
			if os.Args[2] != "--no-wait" {
				style := lipgloss.NewStyle().
					Background(lipgloss.Color("#EE6666")).
					SetString("WARNING: outputting to console\nFor longer lists, piping to a file is suggested")

				fmt.Fprint(os.Stderr, style.String()+"\n")
				time.Sleep(3 * time.Second)
				argIndex--
			}
			output, err = list(os.Args[argIndex], os.Args[argIndex+1])
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			fmt.Println(output)
			os.Exit(1)
		}
	}

	fmt.Println(output)

}
