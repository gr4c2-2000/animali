package fynelanguage

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type key string

type language string

type LanguagesMap map[language]string
type Dictonery map[key]*Item

type Item struct {
	Binded    binding.String
	Languages LanguagesMap
}

type LanguagePack struct {
	dictonery   map[key]*Item
	currentLang language
}

func InitLanguagePack() *LanguagePack {
	lp := LanguagePack{}
	lp.dictonery = make(Dictonery)
	lp.currentLang = "EN"
	return &lp
}

func (d LanguagePack) Add(valueKey string, stored binding.String, langVal string, value string) {
	typeKey := key(valueKey)
	typeLang := language(langVal)
	d.add(typeKey, stored, typeLang, value)

}
func (d LanguagePack) AddLanguageMap(valueKey string, stored binding.String, langmap map[string]string) {
	typeKey := key(valueKey)
	var typeLanguageMap LanguagesMap
	for lang, value := range langmap {
		typeLanguageMap[language(lang)] = value
	}
	d.addLanguageMap(typeKey, stored, typeLanguageMap)
}

func (d LanguagePack) Get(valueKey string) binding.String {
	typeKey := key(valueKey)
	return d.get(typeKey)
}

func (lp LanguagePack) SetLanguage(lang string) {
	lp.currentLang = language(lang)
	for _, Langs := range lp.dictonery {
		if val, ok := Langs.Languages[lp.currentLang]; ok {
			err := Langs.Binded.Set(val)
			if err != nil {
				fyne.LogError("Error getting current data value: ", err)
			}
		}

	}
}
func (lp LanguagePack) get(key key) binding.String {
	if val, ok := lp.dictonery[key]; ok {
		return val.Binded
	}
	binded := binding.NewString()
	err := binded.Set(string(key))
	if err != nil {
		fyne.LogError("Error getting current data value: ", err)
	}
	lp.add(key, binded, lp.currentLang, string(key))
	return binded
}

func (lp LanguagePack) add(key key, stored binding.String, language language, value string) {
	_, ok := lp.dictonery[key]
	if !ok {
		lp.dictonery[key] = &Item{Binded: stored}
		lp.dictonery[key].Languages = make(LanguagesMap)
	}
	lp.dictonery[key].Languages[language] = value
	err := lp.dictonery[key].Binded.Set(lp.dictonery[key].Languages[lp.currentLang])
	if err != nil {
		fyne.LogError("Error getting current data value: ", err)
	}
}

func (lp LanguagePack) addLanguageMap(key key, stored binding.String, langmap LanguagesMap) {
	lp.dictonery[key] = &Item{Binded: stored}
	lp.dictonery[key].Languages = langmap
	err := lp.dictonery[key].Binded.Set(lp.dictonery[key].Languages[lp.currentLang])
	if err != nil {
		fyne.LogError("Error getting current data value: ", err)
	}

}
