package arguments

import (
	"flag"
	"strings"
)

// ParseArguments parses the input arguments and returns the limit, language, and joined string.
//
// Parameter(s):
//   a []string - the input arguments
// Return type(s):
//   int - the limit
//   string - the language
//   string - the joined string
func ParseArguments(a []string) (int, string, string) {
	limit := flag.Int("limit", 10, "Enter number - limit of results\n")
        var language string
	flag.StringVar(&language, "language", "EN", "Enter language - EN, DE, ES, IT, FR, HU, RU\n")
	flag.StringVar(&language, "l", "EN", "Enter language - EN, DE, ES, IT, FR, HU, RU\n")
	flag.Parse()

        rest := flag.Args()
	if len(rest) == 0 {
		return 0, "", ""
	}

	return *limit, strings.ToUpper(language), strings.Join(rest, " ")
}
