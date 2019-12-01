package shortener

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
	url := body.Url

	if !isValidURL(url) {
		respondWithError(w, http.StatusBadRequest, "Invalid url")
		return
	}

	model := Url{
		Url:     url,
		Shortid: shortId(),
	}
	if err := model.findByUrl(a.DB); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			err := model.create(a.DB)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
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
		Shortid: vars["shortId"],
	}

	if err := model.findByShortid(a.DB); err != nil {
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}

	http.Redirect(w, r, model.Url, http.StatusSeeOther)
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
