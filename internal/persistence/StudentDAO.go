package persistence

import (
	"internal/entities"
)

type Students interface {
	FindAll() []entities.Student
	Find(id int) *entities.Student
	Exist(id int) bool
	Delete(id int) bool
	Create(Student entities.Student) bool
	Update(Student entities.Student) bool
	NewStudentDAOMemory() StudentDAOMemory
}
