package arguments

import (
	"flag"
	"fmt"
	"os"
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
  var language string

	flag.StringVar(&language, "language", "EN", "[optinal] Enter language - EN, DE, ES, IT, FR, HU, RU")
	flag.StringVar(&language, "l", "EN", "Alias for -language")
	limit := flag.Int("limit", 10, "[optional] Enter number - limit of results")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] WORD\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

  rest := flag.Args()
	if len(rest) == 0 {
		return 0, "", ""
	}

	return *limit, strings.ToUpper(language), strings.Join(rest, " ")
}
