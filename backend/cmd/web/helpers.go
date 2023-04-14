package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	err = app.errorLogger.Output(2, trace)
	if err != nil {
		log.Fatal(err)
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int, errorText ...string) {
	text := http.StatusText(status)
	if len(errorText) > 0 {
		text = errorText[0]
	}
	http.Error(w, text, status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
