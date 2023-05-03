package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	err = ErrorLogger.Output(2, trace)
	if err != nil {
		ErrorLogger.Println(err)
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int, errorText ...string) {
	text := http.StatusText(status)
	if len(errorText) > 0 {
		text = errorText[0]
	}
	http.Error(w, text, status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}

func Unauthorized(w http.ResponseWriter) {
	ClientError(w, http.StatusUnauthorized)
}
