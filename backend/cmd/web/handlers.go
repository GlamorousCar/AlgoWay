package main

import (
	"context"
	"encoding/json"
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

func handleThemes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/themes" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	themesQuery := "SELECT * FROM theme"
	rows, err := conn.Query(context.Background(), themesQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	themes := make([]theme, 0)
	for rows.Next() {
		theme := theme{}
		err := rows.Scan(&theme.Id, &theme.Title)
		if err == nil {
			themes = append(themes, theme)
		}
	}
	json.NewEncoder(w).Encode(themes)
}
