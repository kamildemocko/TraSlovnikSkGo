package req

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	EN_SK = "https://slovnik.aktuality.sk/api/dictionary/translate/search/en-sk?q={q}"
	SK_EN = "https://slovnik.aktuality.sk/api/dictionary/translate/search/sk-en?q={q}"
)

var (
	ErrorGet = errors.New("error getting data from slovnik aktuality API")
	ErrorRead = errors.New("error reading body from API call")
	)

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

func set_query(theUrl string, value string) string {
	value = url.QueryEscape(value)
	return strings.Replace(theUrl, "{q}", value, 1)
}
