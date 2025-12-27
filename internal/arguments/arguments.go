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
//
//	a []string - the input arguments
//
// Return type(s):
//
//	int - the limit
//	string - the language
//	string - the joined string
func ParseArguments(a []string) (int, string, string) {
	var language string
	var laguageAllowed = map[string]struct{}{
		"EN": {},
		"DE": {},
		"ES": {},
		"IT": {},
		"FR": {},
		"HU": {},
		"RU": {},
	}

	flag.StringVar(&language, "language", "EN", "[optinal] Enter language - EN, DE, ES, IT, FR, HU, RU")
	flag.StringVar(&language, "l", "EN", "Alias for -language")
	limit := flag.Int("limit", 10, "[optional] Enter number - limit of results")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] WORD\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	language = strings.ToUpper(language)
	_, languageOk := laguageAllowed[language]
	if !languageOk {
		flag.Usage()
		return 0, "", ""
	}

	if *limit < 0 {
		flag.Usage()
		return 0, "", ""
	}

	rest := flag.Args()
	if len(rest) == 0 {
		flag.Usage()
		return 0, "", ""
	}

	restJoined := strings.Join(rest, "")
	restJoined = strings.Trim(restJoined, " ")
	if restJoined == "" {
		flag.Usage()
		return 0, "", ""
	}

	return *limit, language, restJoined
}
