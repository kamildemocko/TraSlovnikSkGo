package unmarsh

import (
	"encoding/json"
	"errors"
)

var (
	ErrUnmar = errors.New("error unmarshalling")
)

type Data struct {
	StatusCode int
	ErrorMessage string
	Data struct {
		Regular struct {
			Total int
			Items []map[string]string
		}
		_ map[string]interface{}
		_ map[string]interface{}
		Switched bool
		IsStrict bool
		IsWildcard bool
		IsFuzzy bool
		IsFulltext bool
	}
}

func UnMarsh(body []byte) (Data, error) {
	data := Data{}

	if err := json.Unmarshal(body, &data); err != nil {
		return Data{}, ErrUnmar
	}

	return data, nil
}