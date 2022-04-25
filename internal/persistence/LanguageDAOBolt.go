package persistence

import (
	"encoding/json"
	"internal/entities"
	"log"
	"sort"
)

type LanguageDAOBolt struct {
	Dbms BoltDbb
}

func (l LanguageDAOBolt) FindAll() map[string]entities.Language {
	languagesStrings := l.Dbms.dbGetAll("languages")
	languages := make(map[string]entities.Language)

	for _, element := range languagesStrings {
		language := new(entities.Language)
		decode_err := json.Unmarshal([]byte(element), &language)
		if decode_err != nil {
			log.Fatal(decode_err)
		}
		languages[language.Code] = *language
	}

	keys := make([]string, 0, len(languages))
	for k := range languages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return languages
}

func (l LanguageDAOBolt) Find(code string) entities.Language {
	languageString := l.Dbms.dbGet("languages", code)
	language := new(entities.Language)
	decode_err := json.Unmarshal([]byte(languageString), &language)
	if decode_err != nil {
		log.Fatal(decode_err)
	}
	return *language
}

func (l LanguageDAOBolt) Exists(code string) bool {
	exists := l.Find(code) != entities.Language{}

	return exists
}

func (l LanguageDAOBolt) Delete(code string) bool {
	exists := l.Exists(code)
	if exists {
		l.Dbms.dbDelete("languages", code)
	}

	return exists
}

func (l LanguageDAOBolt) Create(language entities.Language) bool {
	exists := l.Exists(language.Code)
	if !exists {
		encoded, encoded_err := json.Marshal(language)
		if encoded_err != nil {
			log.Fatal(encoded_err)
		}
		l.Dbms.dbPut("languages", language.Code, string(encoded))
	}

	return !exists
}

func (l LanguageDAOBolt) Update(language entities.Language) bool {
	exists := l.Exists(language.Code)
	if exists {
		encoded, encoded_err := json.Marshal(language)
		if encoded_err != nil {
			log.Fatal(encoded_err)
		}
		l.Dbms.dbPut("languages", language.Code, string(encoded))
	}

	return exists
}
