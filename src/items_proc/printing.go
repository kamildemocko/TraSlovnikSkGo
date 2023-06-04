package items_proc

import (
	"fmt"
	"strings"
)

const (
	colorPurple = "\033[35m"
	colorYellow = "\033[33m"
    colorReset = "\033[0m"
)

func PrintResult(inputWord string, data Parsed) {
	ensk := data.Ensk
	sken := data.Sken

	if ensk != nil {
		fmt.Printf("# English -> Slovak (%s)\n", inputWord)
		for key, value := range ensk {
			printKeyValue(key, value)
		}
	}

	if sken != nil {
		if ensk != nil {
			fmt.Println()
		}

		fmt.Printf("# Slovak -> English (%s)\n", inputWord)
		for key, value := range sken {
			printKeyValue(key, value)
		}
	}
}

func printKeyValue(k string, v []string) {
	out := strings.Join(v, ", ")
	fmt.Printf("%s%s%s: %s%s%s\n", colorPurple, k, colorReset, colorYellow, out, colorReset)
}
