package mars_rover

import "errors"

type Obstacle struct {
	position Position
	symbol rune
}

func NewObstacle(obstacleConfig ObstacleConfig) (*Obstacle, error) {
	if obstacleConfig.position.x <= 0 || obstacleConfig.position.y <= 0 {
		return nil, errors.New("invalid coordinates")
	}

	return &Obstacle {
		position: obstacleConfig.position,
		symbol: obstacleConfig.symbol,
	}, nil
}

func (obstacle *Obstacle) SetPosition(mapRows []string, x, y int, symbol rune) error {
	if x > len(mapRows[0]) || y > len(mapRows) {
		return errors.New("coordinates out of bounds")
	}

	row := mapRows[y - 1]
	runes := []rune(row)
	runes[x - 1] = obstacle.symbol
	mapRows[y - 1] = string(runes)

	return nil
}
