package arguments

import "strings"

func ParseArguments(a []string) string {
	if len(a) == 0 || a[0] == "limit" {
		return ""
	}

	if a[0] == "-limit" || a[0] == "--limit" {
		a = a[2:]
	}

	return strings.Join(a, " ")
}
