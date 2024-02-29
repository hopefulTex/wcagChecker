package main

import (
	"fmt"
	"os"
	"strings"
	Color "wcagChecker/color"
)

func duo(isTTY bool, firstHex, lastHex string) (string, error) {
	first, err := Color.FromHex(firstHex)
	if err != nil {
		return "", err
	}
	last, err := Color.FromHex(lastHex)
	if err != nil {
		return "", err
	}

	if isTTY {
		return Color.ComplianceView(first, last), nil
	} else {
		return Color.ComplianceString(first, last), nil
	}
}

// compare a given color to a list of colors
func list(fileName, baseString string) (string, error) {
	var output strings.Builder
	var colorBuf []Color.Color

	// open file
	colorStrings, err := getHexes(fileName)
	if err != nil {
		return "", err
	}

	// convert strings to colors
	base, err := Color.FromHex(baseString)
	if err != nil {
		return "", err
	}

	for _, hex := range colorStrings {
		c, err := Color.FromHex(hex)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}
		colorBuf = append(colorBuf, c)
	}

	// compare
	output.WriteString(fmt.Sprintf("Base: %s\n", base.String()))
	for _, c := range colorBuf {
		output.WriteString(Color.ComplianceString(c, base))
		output.WriteRune('\n')
	}
	return output.String(), nil
}
