package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kamildemocko/tra-slovnik-sk-go/src/arguments"
	"github.com/kamildemocko/tra-slovnik-sk-go/src/items_proc"
	"github.com/kamildemocko/tra-slovnik-sk-go/src/req"
	"github.com/kamildemocko/tra-slovnik-sk-go/src/unmarsh"
)

func main() {
	limit := flag.Int("limit", 10, "Enter number - limit of results\n")
	flag.Parse()

	word := arguments.ParseArguments(os.Args[1:])
	if word == "" {
		fmt.Println("Usage: APPLICATION.exe [-limit INT] INPUT WORD")
		return
	}

	switched_ensk, items_ensk, err := GetData(req.EN_SK, word, *limit)
	if err != nil {
		panic(err)
	}

	switched_sken, items_sken, err := GetData(req.SK_EN, word, *limit)
	if err != nil {
		panic(err)
	}

	all := items_proc.All {
		{"ensk", switched_ensk, items_ensk},
		{"sken", switched_sken, items_sken},
	}

	parsed := items_proc.ParseItems(all)

	items_proc.PrintResult(word, parsed)
}

func GetData(url string, word string, limit int) (bool, []map[string]string, error) {
	// returns bool switched, Data

	body, err := req.Get(url, word)
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
	return data.Data.Switched,items, nil
}
