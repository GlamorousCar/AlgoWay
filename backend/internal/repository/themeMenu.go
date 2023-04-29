package repository

import (
	"context"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4"
)

type ThemeMenuRepository interface {
	GetMenu() (*[]models.ThemeMenu, error)
}

type themeMenuRepositoryPostgres struct {
	conn *pgx.Conn
}

func NewThemeMenuRepositoryPostgres(conn *pgx.Conn) *themeMenuRepositoryPostgres {
	return &themeMenuRepositoryPostgres{conn: conn}
}

func (repo *themeMenuRepositoryPostgres) GetMenu() (*[]models.ThemeMenu, error) {
	query := `SELECT t.id, t.title, t.position,
	a.id, a.title, a.description, a.position, a.theme_id
	FROM algorithm AS a
	JOIN theme AS t ON a.theme_id=t.id
	ORDER BY t.position, a.position`

	rows, err := repo.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	// Contains unique themes
	themes := make([]models.Theme, 0, defaultCapacity)

	// Map, where key - the position of the theme to which the algorithm relates
	// value - slice of algorithms
	algos := make(map[int][]models.Algorithm)

	// Holds position of current unique theme
	currentPosition := -1

	for rows.Next() {
		theme := models.Theme{}
		algo := models.Algorithm{}
		err = rows.Scan(
			&theme.Id, &theme.Title, &theme.Position,
			&algo.Id, &algo.Title, &algo.Description, &algo.Position, &algo.ThemeId,
		)

		if err != nil {
			return nil, err
		}

		themesSize := len(themes)
		if currentPosition == -1 || theme.Position != themes[themesSize-1].Position {
			themes = append(themes, theme)
			currentPosition = theme.Position
		}

		_, found := algos[currentPosition]
		if !found {
			algos[currentPosition] = make([]models.Algorithm, 0, defaultCapacity)
		}

		algos[currentPosition] = append(algos[currentPosition], algo)
	}

	menus := make([]models.ThemeMenu, 0)
	for _, theme := range themes {
		algo := algos[theme.Position]

		elem := models.ThemeMenu{
			Id:         theme.Id,
			Title:      theme.Title,
			Position:   theme.Position,
			Algorithms: algo,
		}
		menus = append(menus, elem)
	}

	return &menus, nil
}
