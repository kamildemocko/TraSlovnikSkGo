package req

import (
	"fmt"

	"github.com/kamildemocko/TraSlovnikSkGo/internal/unmarsh"
)

// GetData retrieves data from the given URL based on the provided word and limit.
//
// Parameters:
//
//	url string - the URL to retrieve data from
//	word string - the word to search for in the data
//	limit int - the maximum number of items to return
//
// Returns:
//
//	bool - if return was switched
//	[]map[string]string - the retrieved data items
//	error - an error if any occurred during the retrieval process
func GetData(url string, word string, limit int) (bool, []map[string]string, error) {
	body, err := Get(url, word)
	if err != nil {
		return false, nil, err
	}

	data, err := unmarsh.UnMarsh(body)
	if err != nil {
		return false, nil, err
	}

	statusCode := data.StatusCode
	if statusCode != 200 {
		return false, nil, fmt.Errorf("expected code 200, got %d", statusCode)
	}

	items := data.Data.Regular.Items
	if len(items) > limit {
		items = items[:limit]
	}
	return data.Data.Switched, items, nil
}
