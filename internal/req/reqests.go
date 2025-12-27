package req

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Url struct {
	Language string
	Normal   string
	Reversed string
}

type Urls struct {
	En Url
	De Url
	Es Url
	Fr Url
	It Url
	Hu Url
	Ru Url
}

var AvailableLanguages = Urls{
	En: Url{
		"English",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/en-sk?q={q}",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/sk-en?q={q}",
	},
	De: Url{
		"Deutsch",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/de-sk?q={q}",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/sk-de?q={q}",
	},
	Es: Url{
		"Spanish",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/es-sk?q={q}",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/sk-es?q={q}",
	},
	Fr: Url{
		"French",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/fr-sk?q={q}",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/sk-fr?q={q}",
	},
	It: Url{
		"Italian",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/it-sk?q={q}",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/sk-it?q={q}",
	},
	Hu: Url{
		"Hungarian",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/hu-sk?q={q}",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/sk-hu?q={q}",
	},
	Ru: Url{
		"Russian",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/ru-sk?q={q}",
		"https://slovnik.aktuality.sk/api/dictionary/translate/search/sk-ru?q={q}",
	},
}

var (
	ErrorGet  = errors.New("error getting data from slovnik aktuality API")
	ErrorRead = errors.New("error reading body from API call")
)

// Get retrieves the data from the specified URL with the given word.
//
// Parameters:
//
//	url string - the URL to retrieve data from
//	word string - the word to include in the query
//
// Return type(s):
//
//	[]byte - the retrieved data
//	error - any error that occurred during the retrieval process
func Get(url string, word string) ([]byte, error) {
	resp, err := http.Get(set_query(url, word))
	if err != nil {
		return nil, ErrorGet
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrorRead
	}

	return body, nil
}

// set_query takes theUrl and value as parameters and returns a string.
//
// theUrl is the URL string where the value will be set.
// value is the string to be escaped and set in the URL.
// It returns the modified URL string.
func set_query(theUrl string, value string) string {
	value = url.QueryEscape(value)
	return strings.Replace(theUrl, "{q}", value, 1)
}
