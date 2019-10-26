package shortener

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

//Body is the response body
type Url struct {
	ID         int64  `json:"id"`
	URL        string `json:"url"`
	SHORT_ID   string `json:"shortId"`
	CREATED_AT string `json:"createdAt"`
	UPDATED_AT string `json:"updatedAt"`
	DELETED_AT string `json:"deletedAt"`
}

//Home page
func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server alive!"))
}

//Shorten url POST method
func (a *App) Shorten(w http.ResponseWriter, r *http.Request) {
	var id int64
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

	shortId := shortId()
	a.DB.QueryRow("SELECT id, shortid FROM shortened_urls WHERE url = ?", url).Scan(&id, &shortId)
	if id == 0 {
		res, err := a.DB.Exec(`INSERT INTO shortened_urls (url, shortid, created_at, updated_at) VALUES(?, ?, now(), now())`, url, shortId)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		id, _ = res.LastInsertId()
	}

	a.DB.QueryRow("SELECT id, url, shortid, created_at, updated_at, deleted_at FROM shortened_urls WHERE id = ?", id).Scan(&body.ID, &body.URL, &body.SHORT_ID, &body.CREATED_AT, &body.UPDATED_AT, &body.DELETED_AT)

	sendResponse(w, http.StatusOK, body)
}

//Redirect route
func (a *App) Redirect(w http.ResponseWriter, r *http.Request) {
	var id int64
	var longURL string

	vars := mux.Vars(r)
	shortId := vars["shortId"]

	a.DB.QueryRow("SELECT id, url FROM shortened_urls WHERE shortid = ?", shortId).
		Scan(&id, &longURL)

	if id == 0 {
		respondWithError(w, http.StatusNotFound, "Not found")
		return
	}
	http.Redirect(w, r, longURL, http.StatusSeeOther)
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true

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
