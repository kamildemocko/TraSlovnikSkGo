package items_proc

type UrlType int

const (
	UrlTypeNorm UrlType = iota
	UrlTypeRev
)

type All = []struct {
	Type     UrlType
	Switched bool
	Items    []map[string]string
}

type Parsed struct {
	Normal   map[string][]string
	Reversed map[string][]string
}

// ParseItems parses the items and returns the parsed result.
//
// Parameter(s): all All
// Return type: Parsed
func ParseItems(all All) Parsed {
	parsed := Parsed{}

	for _, it := range all {
		if it.Switched {
			continue
		}

		if it.Type == UrlTypeNorm {
			parsed.Normal = processItems(it.Type, it.Items)
		} else if it.Type == UrlTypeRev {
			parsed.Reversed = processItems(it.Type, it.Items)
		}
	}

	return parsed
}

// processItems processes items based on the specified UrlType and returns a map of strings to string slices.
//
// type_ UrlType - the type of URL
// items []map[string]string - a slice of maps containing string key-value pairs
// map[string][]string - a map of strings to string slices
func processItems(urlType UrlType, items []map[string]string) map[string][]string {
	prepData := map[string][]string{}

	if urlType == UrlTypeNorm {
		for _, w := range items {
			prepData[w["translation"]] = append(prepData[w["translation"]], w["word_sk"])
		}
	} else if urlType == UrlTypeRev {
		for _, w := range items {
			prepData[w["word_sk"]] = append(prepData[w["word_sk"]], w["translation"])
		}
	}

	return prepData
}
