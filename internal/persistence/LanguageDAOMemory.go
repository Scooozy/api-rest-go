package persistence

import (
	. "internal/entities"
)

type LanguageDAOMemory struct{}

var _ LanguagesDAO = (*LanguageDAOMemory)(nil)

var mapLanguage = make(map[string]Language)

func (l LanguageDAOMemory) FindAll() []Language {
	var tabLanguage []Language
	for _, language := range mapLanguage {
		tabLanguage = append(tabLanguage, language)
	}
	return tabLanguage
}

func (l LanguageDAOMemory) Find(code string) *Language {
	la := mapLanguage[code]
	if l.Exist(code) {
		return &la
	}
	return nil
}
func (l LanguageDAOMemory) Exist(code string) bool {
	_, ok := mapLanguage[code]
	if ok {
		return true
	} else {
		return false
	}

}

func (l LanguageDAOMemory) Delete(code string) bool {
	if l.Exist(code) {
		delete(mapLanguage, code)
		return true
	}
	return false
}

func (l LanguageDAOMemory) Create(language Language) bool {
	if !l.Exist(language.Code) {
		mapLanguage[language.Code] = language
		return true
	}
	return false
}

func (l LanguageDAOMemory) Update(language Language) bool {
	if l.Exist(language.Code) {
		mapLanguage[language.Code] = language
		return true
	}
	return false
}

func (l LanguageDAOMemory) NewLanguageDAOMemory() LanguageDAOMemory {
	return LanguageDAOMemory{}
}
