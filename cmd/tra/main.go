package main

import (
	"log"
	"os"

	"github.com/kamildemocko/TraSlovnikSkGo/internal/arguments"
	"github.com/kamildemocko/TraSlovnikSkGo/internal/items_proc"
	"github.com/kamildemocko/TraSlovnikSkGo/internal/req"
)

func main() {
	validLanguages := map[string]bool{
		"EN": true,
		"DE": true,
		"ES": true,
		"IT": true,
		"FR": true,
		"HU": true,
		"RU": true,
	}

	limit, language, word := arguments.ParseArguments(os.Args[1:])

	if word == "" || !validLanguages[language] {
		items_proc.PrintHelp()
		return
	}

	languageMap := map[string]req.Url{
		"EN": req.AvailableLanguages.En,
		"DE": req.AvailableLanguages.De,
		"ES": req.AvailableLanguages.Es,
		"IT": req.AvailableLanguages.It,
		"FR": req.AvailableLanguages.Fr,
		"HU": req.AvailableLanguages.Hu,
		"RU": req.AvailableLanguages.Ru,
	}
	chosenLanguage := languageMap[language]

	switchedNormal, itemsNormal, err := req.GetData(chosenLanguage.Normal, word, limit)
	if err != nil {
		log.Fatal(err)
	}

	switchedReversed, itemsReversed, err := req.GetData(chosenLanguage.Reversed, word, limit)
	if err != nil {
		log.Fatal(err)
	}

	all := items_proc.All{
		{items_proc.UrlTypeNorm, switchedNormal, itemsNormal},
		{items_proc.UrlTypeRev, switchedReversed, itemsReversed},
	}

	parsed := items_proc.ParseItems(all)

	items_proc.PrintResult(word, parsed, chosenLanguage.Language)
}
