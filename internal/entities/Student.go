package entities

type Student struct {
	Id           int
	FirstName    string
	LastName     string
	Age          int
	LanguageCode string
}

func NewStudent(Id int, FirstName string, LastName string, Age int, LanguageCode string) Student {
	return Student{
		Id:           Id,
		FirstName:    FirstName,
		LastName:     LastName,
		Age:          Age,
		LanguageCode: LanguageCode,
	}
}

func String(s Student) string {
	return "{" +
		"Etudiant n° : " + string(s.Id) +
		"FirstName : " + s.FirstName +
		"LastName :" + s.LastName +
		"Age : " + string(s.Age) +
		"LanguageCode : " + s.LanguageCode +
		"}"
}
