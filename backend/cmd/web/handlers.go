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

func getThemesMenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/themes/menu" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	query := `SELECT t.id, t.title, t.position,
		a.id, a.title, a.description, a.position, a.theme_id
		FROM algorithm AS a
		JOIN theme AS t ON a.theme_id=t.id`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		return
	}

	themes := make(map[theme][]algorithm)
	for rows.Next() {
		theme := theme{}
		algo := algorithm{}
		err := rows.Scan(
			&theme.Id, &theme.Title, &theme.Position,
			&algo.Id, &algo.Title, &algo.Description, &algo.Position, &algo.ThemeId,
		)
		if err != nil {
			log.Fatal(err)
		} else {
			_, found := themes[theme]
			if !found {
				themes[theme] = make([]algorithm, 1)
				themes[theme][0] = algo
			} else {
				themes[theme] = append(themes[theme], algo)
			}
		}
	}

	themeMenus := make([]themeMenu, 0)
	for theme, algo := range themes {
		elem := themeMenu{
			Id:         theme.Id,
			Title:      theme.Title,
			Position:   theme.Position,
			Algorithms: algo,
		}
		themeMenus = append(themeMenus, elem)
	}
	jsonStr, jsonErr := json.Marshal(themeMenus)
	if jsonErr != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", jsonErr)
	}
	_, writeErr := w.Write(jsonStr)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
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
