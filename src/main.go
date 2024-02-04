package main

import (
	"fmt"
	"os"
	Color "wcagChecker/color"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("error: not enough arguments. (want #hexcol #hexcol)\n")
		os.Exit(1)
	}

	first, err := Color.FromHex(os.Args[1])
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
	last, err := Color.FromHex(os.Args[2])
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(Color.ComplianceView(first, last))
	// score, contrast := Color.Compliance(first, last)

	// contrastStr := fmt.Sprintf("%f", contrast)
	// index := strings.IndexRune(contrastStr, '.')
	// contrastStr = contrastStr[0 : index+2]

	// fmt.Printf("Score~     %s\n", score.ToString())
	// fmt.Printf("Contrast~  %s:1\n", contrastStr)

}
