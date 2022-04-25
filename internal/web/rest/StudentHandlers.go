package web

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentsHandlers struct {
	DAO persistence.Students
}

// swagger:operation GET /students/{id} Students StudentHandlerOne
// ---
// summary: Renvoie l'étudiant d'id spécifié
// parameters:
// - name: id
//   in: path
//   description: id de l'étudiant
//   type: int
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/studentRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (sh StudentsHandlers) StudentHandlerOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonString, err := json.Marshal(sh.DAO.Find(id_int))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation GET /students Students AllStudentHandler
// ---
// summary: Renvoie la liste de tous les étudiants
// responses:
//   "200":
//     "$ref": "#/responses/studentsRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (sh StudentsHandlers) AllStudentHandler(w http.ResponseWriter, r *http.Request) {
	jsonString, err := json.Marshal(sh.DAO.FindAll())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation POST /students Students CreatestudentHandler
// ---
// summary: Ajoute un nouvel étudiant
// parameters:
// - name: student
//   description: L'étudiant à ajouter
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "201":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (sh StudentsHandlers) CreatestudentHandler(w http.ResponseWriter, r *http.Request) {
	student := entities.NewStudent(0, "", "", 0, "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)

	isOk, err := json.Marshal(sh.DAO.Create(student))

	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, string(isOk))
}

// swagger:operation PUT /students Students UpdatestudentHandler
// ---
// summary: Modifie un étudiant
// parameters:
// - name: student
//   description: L'étudiant à modifier
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (sh StudentsHandlers) UpdatestudentHandler(w http.ResponseWriter, r *http.Request) {

	student := entities.NewStudent(0, "", "", 0, "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)

	isOk, err := json.Marshal(sh.DAO.Update(student))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(isOk))
}

// swagger:operation DELETE /students/{id} Students DelstudentHandler
// ---
// summary: Supprime l'étudiant d'id spécifié
// parameters:
// - name: id
//   in: path
//   description: id de l'étudiant
//   type: string
//   required: true
// - name: student
//   description: L'étudiant à supprimer
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (sh StudentsHandlers) DelstudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	isOk, err := json.Marshal(sh.DAO.Delete(id_int))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(isOk))
}

func (sh StudentsHandlers) InitializeStudentsRoutes(r *mux.Router) {
	r.HandleFunc("/rest/students/{id}", sh.StudentHandlerOne).Methods("GET")
	r.HandleFunc("/rest/students", sh.AllStudentHandler).Methods("GET")
	r.HandleFunc("/rest/students", sh.CreatestudentHandler).Methods("POST")
	r.HandleFunc("/rest/students", sh.UpdatestudentHandler).Methods("PUT")
	r.HandleFunc("/rest/students/{id}", sh.DelstudentHandler).Methods("DELETE")
}
