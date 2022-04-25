package persistence

import (
	"encoding/json"
	"internal/entities"
	"log"
	"sort"
	"strconv"
)

type StudentDAOBolt struct {
	Dbms BoltDbb
}

func (s *StudentDAOBolt) FindAll() map[int]entities.Student {
	studentsStrings := s.Dbms.dbGetAll("students")
	students := make(map[int]entities.Student)

	for _, element := range studentsStrings {
		student := new(entities.Student)
		decode_err := json.Unmarshal([]byte(element), &student)
		if decode_err != nil {
			log.Fatal(decode_err)
		}
		students[student.Id] = *student
	}

	keys := make([]int, 0, len(students))
	for k := range students {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return students
}

func (s *StudentDAOBolt) Find(id int) entities.Student {
	studentString := s.Dbms.dbGet("students", strconv.Itoa(id))
	student := new(entities.Student)
	decode_err := json.Unmarshal([]byte(studentString), &student)
	if decode_err != nil {
		//log.Fatal(decode_err)
	}
	return *student
}

func (s *StudentDAOBolt) Exists(id int) bool {
	exists := s.Find(id) != entities.Student{}

	return exists
}

func (s *StudentDAOBolt) Delete(id int) bool {
	exists := s.Exists(id)
	if exists {
		s.Dbms.dbDelete("students", strconv.Itoa(id))
	}

	return exists
}

func (s *StudentDAOBolt) Create(student entities.Student) bool {
	exists := s.Exists(student.Id)
	if !exists {
		encoded, encoded_err := json.Marshal(student)
		if encoded_err != nil {
			log.Fatal(encoded_err)
		}
		s.Dbms.dbPut("students", strconv.Itoa(student.Id), string(encoded))
	}

	return !exists
}

func (s *StudentDAOBolt) Update(student entities.Student) bool {
	exists := s.Exists(student.Id)
	if exists {
		encoded, encoded_err := json.Marshal(student)
		if encoded_err != nil {
			log.Fatal(encoded_err)
		}
		s.Dbms.dbPut("students", strconv.Itoa(student.Id), string(encoded))
	}

	return exists
}
