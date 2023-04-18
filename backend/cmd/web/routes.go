package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/themes/menu", app.getThemesMenu)
	mux.HandleFunc("/theory", app.getAlgorithmTheory)
	mux.HandleFunc("/task", app.getAlgorithmTasks)
	mux.HandleFunc("/auth/register", app.registerUser)
	mux.HandleFunc("/auth/login", app.logiUser)

	return mux
}
