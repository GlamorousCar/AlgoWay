package app

import (
	"github.com/GlamorousCar/AlgoWay/internal/controllers"
	"github.com/GlamorousCar/AlgoWay/internal/repository"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func newServeMux(
	algorithmTheoryUseCase *usecase.AlgorithmTheoryUseCase,
	homeUseCase *usecase.HomeUseCase,
	userUseCase *usecase.UserUseCase,
	themeMenuUseCase *usecase.ThemeMenuUseCase,
	taskUseCase *usecase.TaskUseCase,
	checkSystemUseCase *usecase.CheckSystemUseCase,
) *http.ServeMux {
	mux := http.NewServeMux()

	algoTheoryHandler := controllers.NewAlgoTheoryHandler(algorithmTheoryUseCase)
	homeHandler := controllers.NewHomeHandler(homeUseCase)
	authHandler := controllers.NewAuthHandler(userUseCase)
	themeMenuHandler := controllers.NewThemeMenuHandler(themeMenuUseCase)
	taskHandler := controllers.NewTaskHandler(taskUseCase)
	checkSystemHandler := controllers.NewCheckSystemHandler(checkSystemUseCase)

	mux.HandleFunc("/", homeHandler.Home)
	mux.HandleFunc("/themes/menu", themeMenuHandler.GetThemeMenu)
	mux.HandleFunc("/theory", algoTheoryHandler.GetAlgorithmTheory)
	mux.HandleFunc("/task", taskHandler.GetAlgorithmTasks)
	mux.HandleFunc("/auth/register", authHandler.RegisterUser)
	mux.HandleFunc("/auth/login", authHandler.LoginUser)
	mux.HandleFunc("/check_task", checkSystemHandler.CheckTask)

	return mux
}

func initServeMux(conn *pgx.Conn) *http.ServeMux {
	algoTheoryRepository := repository.NewAlgoTheoryRepositoryPostgres(conn)
	algorithmTheoryUseCase := usecase.NewAlgorithmTheoryUseCase(algoTheoryRepository)

	homeRepository := repository.NewHomeRepositoryPostgres(conn)
	homeUseCase := usecase.NewHomeUseCase(homeRepository)

	userRepository := repository.NewUserRepositoryPostgres(conn)
	userUseCase := usecase.NewUserUseCase(userRepository)

	themeMenuRepository := repository.NewThemeMenuRepositoryPostgres(conn)
	themeMenuUseCase := usecase.NewThemeMenuUseCase(themeMenuRepository)

	taskRepository := repository.NewTaskRepositoryPostgres(conn)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)

	checkSystemRepository := repository.NewCheckSystemRepositoryPostgres(conn)
	checkSystemUseCase := usecase.NewCheckSystemUseCase(checkSystemRepository)

	return newServeMux(
		algorithmTheoryUseCase,
		homeUseCase,
		userUseCase,
		themeMenuUseCase,
		taskUseCase,
		checkSystemUseCase,
	)
}
