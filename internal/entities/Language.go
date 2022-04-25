package entities

type Language struct {
	Code string
	Name string
}

func NewLanguage(Code string, Name string) Language {
	return Language{
		Code: Code,
		Name: Name,
	}
}

func toString(l Language) string {
	return "{" +
		"Language n° : " + l.Code +
		"Name : " + l.Name +
		"}"
}
