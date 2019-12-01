package shortener

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server alive!"))
}

func (a *App) Shorten(w http.ResponseWriter, r *http.Request) {
	var body Url

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	url := body.URL

	if !isValidURL(url) {
		respondWithError(w, http.StatusBadRequest, "Invalid url")
		return
	}

	model := Url{
		URL:      url,
		SHORT_ID: shortId(),
	}
	if err := model.findByUrl(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			res, err := model.create(a.DB)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			model.ID, _ = res.LastInsertId()
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	model.findById(a.DB)

	sendResponse(w, http.StatusOK, model)
}

//Redirect route
func (a *App) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	model := Url{
		SHORT_ID: vars["shortId"],
	}

	if err := model.findByShortid(a.DB); err != nil {
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}

	http.Redirect(w, r, model.URL, http.StatusSeeOther)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	sendResponse(w, code, map[string]string{"error": message})
}

func sendResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
