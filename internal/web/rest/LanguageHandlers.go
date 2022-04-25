package web

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type LanguagesHandlers struct {
	persistence.LanguagesDAO
}

// swagger:operation GET /languages/{code} Languages LanguageHandlerOne
// ---
// summary: Renvoie le langage ayant le code spécifié
// parameters:
// - name: id
//   in: path
//   description: code du langage
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/languageRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (lh *LanguagesHandlers) LanguageHandlerOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]

	jsonString, err := json.Marshal(lh.Find(code))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation GET /languages Languages AllLanguageHandler
// ---
// summary: Renvoie la liste de tous les langages de prorgammation
// responses:
//   "200":
//     "$ref": "#/responses/languagesRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (lh *LanguagesHandlers) AllLanguageHandler(w http.ResponseWriter, r *http.Request) {
	jsonString, err := json.Marshal(lh.FindAll())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation POST /languages Languages CreateLanguageHandler
// ---
// summary: Ajoute un nouveau langage de progammation
// parameters:
// - name: language
//   description: Le langage à ajouter
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "201":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (lh *LanguagesHandlers) CreateLanguageHandler(w http.ResponseWriter, r *http.Request) {
	language := entities.NewLanguage("", "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &language)

	isOk, err := json.Marshal(lh.Create(language))

	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, string(isOk))

}

// swagger:operation PUT /languages Languages UpdateLanguageHandler
// ---
// summary: Modifie un langage de progammation
// parameters:
// - name: language
//   description: Le langage à modifier
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (lh *LanguagesHandlers) UpdateLanguageHandler(w http.ResponseWriter, r *http.Request) {

	language := entities.NewLanguage("", "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &language)

	isOk, err := json.Marshal(lh.Update(language))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(isOk))
}

// swagger:operation DELETE /languages/{code} Languages DelLanguageHandler
// ---
// summary: Supprime le langage ayant le code spécifié
// parameters:
// - name: id
//   in: path
//   description: code du langage
//   type: string
//   required: true
// - name: language
//   description: Le langage à supprimer
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (lh *LanguagesHandlers) DelLanguageHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]

	isOk, err := json.Marshal(lh.Delete(code))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(isOk))

}

func (lh *LanguagesHandlers) InitializeLanguagesRoutes(r *mux.Router) {
	r.HandleFunc("/rest/languages/{code}", lh.LanguageHandlerOne).Methods("GET")
	r.HandleFunc("/rest/languages", lh.AllLanguageHandler).Methods("GET")
	r.HandleFunc("/rest/languages", lh.CreateLanguageHandler).Methods("POST")
	r.HandleFunc("/rest/languages", lh.UpdateLanguageHandler).Methods("PUT")
	r.HandleFunc("/rest/languages/{code}", lh.DelLanguageHandler).Methods("DELETE")
}
