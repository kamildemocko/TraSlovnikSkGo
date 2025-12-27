package main

import (
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

	var chosenLanguage req.Url
	if language == "EN" {
		chosenLanguage = req.AvailableLanguages.En
	} else if language == "DE" {
		chosenLanguage = req.AvailableLanguages.De
	} else if language == "ES" {
		chosenLanguage = req.AvailableLanguages.Es
	} else if language == "IT" {
		chosenLanguage = req.AvailableLanguages.It
	} else if language == "FR" {
		chosenLanguage = req.AvailableLanguages.Fr
	} else if language == "HU" {
		chosenLanguage = req.AvailableLanguages.Hu
	} else if language == "RU" {
		chosenLanguage = req.AvailableLanguages.Ru
	}

	switchedNormal, itemsNormal, err := req.GetData(chosenLanguage.Normal, word, limit)
	if err != nil {
		panic(err)
	}

	switchedReversed, itemsReversed, err := req.GetData(chosenLanguage.Reversed, word, limit)
	if err != nil {
		panic(err)
	}

	all := items_proc.All{
		{items_proc.UrlTypeNorm, switchedNormal, itemsNormal},
		{items_proc.UrlTypeRev, switchedReversed, itemsReversed},
	}

	parsed := items_proc.ParseItems(all)

	items_proc.PrintResult(word, parsed, chosenLanguage.Language)
}
