package req

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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

// Sets standard headers
//
// Parameters:
//
// req http.Request - request to attach the header to
func setHeader(req http.Request) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", "https://slovnik.aktuality.sk/preklad/anglicko-slovensky/?q=klb")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"153\", \"Not_A Brand\";v=\"8\", \"Chromium\";v=\"153\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/153.0.0.0 Safari/537.36")
}

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
	finalUrl := setQuery(url, word)

	req, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		return nil, err
	}

	setHeader(*req)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error response from API: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Unexpected status code: %d\nRaw Response: %s", resp.StatusCode, string(errBody))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrorRead
	}

	return body, nil
}

// setQuery takes theUrl and value as parameters and returns a string.
//
// theUrl is the URL string where the value will be set.
// value is the string to be escaped and set in the URL.
// It returns the modified URL string.
func setQuery(theUrl string, value string) string {
	value = url.QueryEscape(value)
	return strings.Replace(theUrl, "{q}", value, 1)
}
