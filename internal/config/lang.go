package config

import (
	"encoding/json"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func InitI18n() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	_, err := bundle.LoadMessageFile("locales/en.json")
	if err != nil {
		log.Println("Warning: Could not load en.json:", err)
	}

	_, err = bundle.LoadMessageFile("locales/pt-br.json")
	if err != nil {
		log.Println("Warning: Could not load pt-br.json:", err)
	}
}

func GetLanguage() string {
	lang := GetConfig("language")
	if lang == "" {
		lang = "en"
		SetConfig("language", lang)
	}
	return lang
}

func Translate(key string) string {
	lang := GetLanguage()
	localizer := i18n.NewLocalizer(bundle, lang)
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: key})
}
