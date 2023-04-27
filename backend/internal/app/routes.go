package app

import (
	"github.com/GlamorousCar/AlgoWay/internal/database"
	"github.com/GlamorousCar/AlgoWay/internal/transport"
	"net/http"
)

func routes(db database.DB) *http.ServeMux {
	mux := http.NewServeMux()
	handlers := transport.MakeMainHandler(db)

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/themes/menu", handlers.GetThemesMenu)
	mux.HandleFunc("/theory", handlers.GetAlgorithmTheory)
	mux.HandleFunc("/task", handlers.GetAlgorithmTasks)
	mux.HandleFunc("/auth/register", handlers.RegisterUser)
	mux.HandleFunc("/auth/login", handlers.LoginUser)

	return mux
}
