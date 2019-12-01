package shortener

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/shorten", a.Shorten).Methods("POST")
	a.Router.HandleFunc("/{shortId}", a.Redirect).Methods("GET")
}
