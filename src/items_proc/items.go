package items_proc


type All = []struct{
	Type string
	Switched bool
	Items []map[string]string
}

type Parsed struct {
	Ensk map[string][]string
	Sken map[string][]string
}

func ParseItems(all All) Parsed {
	parsed := Parsed{}

	for _, it := range all {
		if it.Switched {
			continue
		}

		if it.Type == "sken" {
			parsed.Sken = processItems(it.Type, it.Items)
		} else if it.Type == "ensk" {
			parsed.Ensk = processItems(it.Type, it.Items)
		}
	}

	return parsed
}

func processItems(type_ string, items []map[string]string) map[string][]string {
	preparedData := map[string][]string{}

	if type_ == "sken" {
		for _, w := range items {
			preparedData[w["word_sk"]] = append(preparedData[w["word_sk"]], w["translation"])
		}
	} else if type_ == "ensk" {
		for _, w := range items {
			preparedData[w["translation"]] = append(preparedData[w["translation"]], w["word_sk"])
		}
	}

	return preparedData
}
