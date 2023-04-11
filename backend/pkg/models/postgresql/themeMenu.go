package postgresql

import (
	"context"
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"github.com/jackc/pgx/v4"
	"log"
)

type ThemeMenuModel struct {
	Conn *pgx.Conn
}

func (m ThemeMenuModel) Get() (*[]models.ThemeMenu, error) {
	query := `SELECT t.id, t.title, t.position,
		a.id, a.title, a.description, a.position, a.theme_id
		FROM algorithm AS a
		JOIN theme AS t ON a.theme_id=t.id`

	rows, err := m.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	themes := make(map[models.Theme][]models.Algorithm)
	for rows.Next() {
		theme := models.Theme{}
		algo := models.Algorithm{}
		err := rows.Scan(
			&theme.Id, &theme.Title, &theme.Position,
			&algo.Id, &algo.Title, &algo.Description, &algo.Position, &algo.ThemeId,
		)
		if err != nil {
			log.Fatal(err)
		} else {
			_, found := themes[theme]
			if !found {
				themes[theme] = make([]models.Algorithm, 1)
				themes[theme][0] = algo
			} else {
				themes[theme] = append(themes[theme], algo)
			}
		}
	}

	menus := make([]models.ThemeMenu, 0)
	for theme, algo := range themes {
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
