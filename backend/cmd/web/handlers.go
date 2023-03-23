package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func getAlgorithmTheory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	var res = algorithmTheory{}
	err = conn.QueryRow(context.Background(), "SELECT id,content FROM theory WHERE algorithm_id=$1", id).Scan(&res.ID, &res.Content)
	if res.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	jsonResp, err := json.Marshal(res)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
