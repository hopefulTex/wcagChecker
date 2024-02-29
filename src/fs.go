package main

import (
	"os"
	"strings"
)

func getHexes(path string) ([]string, error) {
	var hexList []string

	text, err := open(path)
	if err != nil {
		return hexList, err
	}

	hexList = strings.Split(text, "\n")
	for i := range hexList {
		hexList[i] = strings.TrimSpace(hexList[i])
	}

	var trimmedList []string

	for _, hex := range hexList {
		if hex != "" {
			trimmedList = append(trimmedList, hex)
		}
	}

	return trimmedList, nil
}

func open(path string) (string, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	//fmt.Printf("File Opened, Text:\n%s", string(buf))

	return string(buf), nil
}
