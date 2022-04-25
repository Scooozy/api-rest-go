// PROJET API WEB REST
//
// Projet création d'une API REST en Go
//
// Terms Of Service:
//
//		Schemes: http, https
//		Host: localhost:8084
//		BasePath: /rest
//		Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
// swagger:meta
package main

import (
	"encoding/json"
	"internal/entities"
	"internal/persistence"
	"log"
	"net/http"
	"strconv"

	. "internal/web/rest"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	/*bdd := new(persistence.BoltDbb)

	bdd.dbOpen("bdd.db")
	defer bdd.dbClose()
	bdd.dbCreateBucket("languages")
	bdd.dbCreateBucket("students")

	var encoded []byte
	var encoded_err error

	leStudentDAO := new(persistence.StudentDAOBolt)

	student1 := entities.NewStudent(1, "Emma", "Loisel", 23, "js")
	student2 := entities.NewStudent(2, "test", "test", 18, "c")

	encoded, encoded_err = json.Marshal(student1)
	if encoded_err != nil {
		log.Fatal(encoded_err)
	}
	bdd.dbPut("students", strconv.Itoa(student1.Id), string(encoded))

	encoded, encoded_err = json.Marshal(student2)
	if encoded_err != nil {
		log.Fatal(encoded_err)
	}
	bdd.dbPut("students", strconv.Itoa(student2.Id), string(encoded))

	leStudentDAO.Dbms = *bdd

	leLanguageDAO := new(persistence.LanguageDAOBolt)

	language1 := entities.NewLanguage("js", "JavaScript")
	language2 := entities.NewLanguage("c", "C")

	encoded, encoded_err = json.Marshal(language1)
	if encoded_err != nil {
		log.Fatal(encoded_err)
	}
	bdd.dbPut("languages", language1.Code, string(encoded))

	encoded, encoded_err = json.Marshal(language2)
	if encoded_err != nil {
		log.Fatal(encoded_err)
	}
	bdd.dbPut("languages", language2.Code, string(encoded))

	leLanguageDAO.Dbms = *bdd

	studentsHandlers := new(StudentsHandlers)
	studentsHandlers.DAO = leStudentDAO
	studentsHandlers.InitializeStudentsRoutes(r)
	LanguagesHandlers := new(LanguagesHandlers)
	LanguagesHandlers.DAO = leLanguageDAO
	LanguagesHandlers.InitializeLanguagesRoutes(r)*/

	fs := http.FileServer(http.Dir("./swagger"))
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
