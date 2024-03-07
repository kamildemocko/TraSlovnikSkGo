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

// PrintResult prints the result of the parsed data for a given input word and language.
//
// inputWord string, data Parsed, lang string
func PrintResult(inputWord string, data Parsed, lang string) {
	n := data.Normal
	r := data.Reversed

	if n != nil {
		fmt.Printf("# %s -> Slovak (%s)\n", lang, inputWord)
		for key, value := range n {
			printKeyValue(key, value)
		}
	}

	if r != nil {
		if n != nil {
			fmt.Println()
		}

		fmt.Printf("# Slovak -> %s (%s)\n", lang, inputWord)
		for key, value := range r {
			printKeyValue(key, value)
		}
	}
}

// printKeyValue is a Go function that prints the key and its corresponding values.
//
// It takes a string key and a slice of strings as parameters and does not return any value.
func printKeyValue(k string, v []string) {
	out := strings.Join(v, ", ")
	fmt.Printf("%s%s%s: %s%s%s\n", colorPurple, k, colorReset, colorYellow, out, colorReset)
}

func PrintHelp() {
		fmt.Println("Usage: program [-limit INT] [-l EN|DE|ES|IT|FR|HU|RU] INPUT WORD")
		fmt.Println("Arguments:")
		fmt.Println("\t-limit | --limit (INT): limit results")
		fmt.Println("\t-l | --language (STR): translation language, default EN, available: EN|DE|ES|IT|FR|HU|RU")
}