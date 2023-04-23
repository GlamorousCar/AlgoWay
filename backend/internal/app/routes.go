package app

import (
	"github.com/GlamorousCar/AlgoWay/internal/transport"
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", transport.Home)
	mux.HandleFunc("/themes/menu", transport.GetThemesMenu)
	mux.HandleFunc("/theory", transport.GetAlgorithmTheory)
	mux.HandleFunc("/task", transport.GetAlgorithmTasks)
	mux.HandleFunc("/auth/register", transport.RegisterUser)
	mux.HandleFunc("/auth/login", transport.LoginUser)

	return mux
}
