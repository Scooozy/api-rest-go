package persistence

import (
	"internal/entities"
)

type LanguagesDAO interface {
	FindAll() []entities.Language
	Find(code string) *entities.Language
	Exist(code string) bool
	Delete(code string) bool
	Create(Language entities.Language) bool
	Update(Language entities.Language) bool
	NewLanguageDAOMemory() LanguageDAOMemory
}
