package fynelanguage

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type LanguagesMap map[string]string
type Items map[string]*Languages

var CurrentLang string = "EN"

type Languages struct {
	Binded   binding.String
	Language LanguagesMap
}

var LanguageStorage Items

func init() {
	LanguageStorage = make(Items)
}

func (l Items) Add(key string, stored binding.String, lang string, value string) {
	_, ok := l[key]
	if !ok {
		l[key] = &Languages{Binded: stored}
		l[key].Language = make(LanguagesMap)
	}
	l[key].Language[lang] = value
	err := l[key].Binded.Set(l[key].Language[CurrentLang])
	if err != nil {
		fyne.LogError("Error getting current data value: ", err)
	}
}
func (l Items) AddLangMap(key string, stored binding.String, langmap LanguagesMap) {
	l[key] = &Languages{Binded: stored}
	l[key].Language = langmap
	err := l[key].Binded.Set(l[key].Language[CurrentLang])
	if err != nil {
		fyne.LogError("Error getting current data value: ", err)
	}

}
func (l Items) Get(key string) binding.String {
	if val, ok := l[key]; ok {
		return val.Binded
	}
	binded := binding.NewString()
	err := binded.Set(key)
	if err != nil {
		fyne.LogError("Error getting current data value: ", err)
	}
	l.Add(key, binded, CurrentLang, key)
	return binded
}
func (l Items) SetLanguage(lang string) {
	CurrentLang = lang
	for _, Langs := range l {
		if val, ok := Langs.Language[CurrentLang]; ok {
			err := Langs.Binded.Set(val)
			if err != nil {
				fyne.LogError("Error getting current data value: ", err)
			}
		}

	}
}
