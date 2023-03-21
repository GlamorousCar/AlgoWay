package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	var version string
	var err error

	err = conn.QueryRow(context.Background(), "select version()").Scan(&version)
	if err != nil {
		os.Exit(1)
	}
	w.Write([]byte(fmt.Sprintf("Успешное подключение\n %s", version)))

}
