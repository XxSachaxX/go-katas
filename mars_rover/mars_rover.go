package mars_rover

import (
	"errors"
	"strings"
)

type Rover struct {
	x int
	y int
	direction string
}

func NewRover(x, y int, direction string) (Rover) {
	return Rover {
		x: x,
		y: y,
		direction: direction,
	}
}

func MakeMap(width, height int) ([]string, error) {
	if width <=0 {
		return nil, errors.New("invalid width: cannot be zero or negative")
	}

	if height <=0 {
		return nil, errors.New("invalid height: cannot be zero or negative")
	}

	mapRows := make([]string, height)
	for i := range mapRows {
		mapRows[i] = strings.Repeat("-", width)
	}

	return mapRows, nil
}

func SetRoverPosition(mapRows []string, x, y int, direction rune) error {
	if !isValidDirection(string(direction)) {
		return errors.New("invalid direction")
	}

	return setPosition(mapRows, x, y, direction)
}

func SetObstaclePosition(mapRows []string, x, y int) error {
	return setPosition(mapRows, x, y, 'O')
}

func setPosition(mapRows []string, x, y int, symbol rune) error {
	if x <= 0 || y <= 0 {
		return errors.New("negative coordinates are not allowed")
	}

	if x > len(mapRows[0]) || y > len(mapRows) {
		return errors.New("coordinates out of bounds")
	}

	row := mapRows[y - 1]
	runes := []rune(row)
	runes[x - 1] = symbol
	mapRows[y - 1] = string(runes)

	return nil
}

func isValidDirection(direction string) bool {
	if direction == "N" || direction == "S" || direction == "E" || direction == "W" {
		return true
	}
	return false
}
