package unmarsh

import (
	"encoding/json"
	"errors"
)

var (
	ErrUnmar = errors.New("error unmarshalling")
)

type Data struct {
	StatusCode   int
	ErrorMessage string
	Data         struct {
		Regular struct {
			Total int
			Items []map[string]string
		}
		Switched   bool
		IsStrict   bool
		IsWildcard bool
		IsFuzzy    bool
		IsFulltext bool
	}
}

// UnMarsh is a Go function that unmarshals the provided byte slice into a Data struct.
//
// It takes a byte slice as a parameter and returns a Data struct and an error.
func UnMarsh(body []byte) (Data, error) {
	data := Data{}

	if err := json.Unmarshal(body, &data); err != nil {
		return Data{}, ErrUnmar
	}

	return data, nil
}
