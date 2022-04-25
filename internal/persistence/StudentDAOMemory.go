package persistence

import (
	. "internal/entities"
)

type StudentDAOMemory struct{}

var _ Students = (*StudentDAOMemory)(nil)

var mapStudent = make(map[int]Student)

func (s StudentDAOMemory) FindAll() []Student {
	var tabStudent []Student
	for _, student := range mapStudent {
		tabStudent = append(tabStudent, student)
	}
	return tabStudent
}

func (s StudentDAOMemory) Find(id int) *Student {
	st := mapStudent[id]
	if s.Exist(id) {
		return &st
	}
	return nil
}
func (s StudentDAOMemory) Exist(id int) bool {
	_, ok := mapStudent[id]
	if ok {
		return true
	} else {
		return false
	}

}

func (s StudentDAOMemory) Delete(id int) bool {
	if s.Exist(id) {
		delete(mapStudent, id)
		return true
	}
	return false
}

func (s StudentDAOMemory) Create(student Student) bool {
	if !s.Exist(student.Id) {
		mapStudent[student.Id] = student
		return true
	}
	return false
}

func (s StudentDAOMemory) Update(student Student) bool {
	if s.Exist(student.Id) {
		mapStudent[student.Id] = student
		return true
	}
	return false
}

func (s StudentDAOMemory) NewStudentDAOMemory() StudentDAOMemory {
	return StudentDAOMemory{}
}
