package mars_rover

import (
	"errors"
	"strings"
)

type Position struct {
	x int
	y int
	SetPosition func(mapRows []string, x, y int, symbol rune) error
}

type Obstacle struct {
	position Position
}

func NewObstacle(x, y int) (Obstacle, error) {
	if x <= 0 || y <= 0 {
		return Obstacle{}, errors.New("invalid coordinates")
	}

	return Obstacle {
		position: Position{
			x: x,
			y: y,
		},
	}, nil
}

func (obstacle Obstacle) SetPosition(mapRows []string, x, y int, symbol rune) error {
	if x > len(mapRows[0]) || y > len(mapRows) {
		return errors.New("coordinates out of bounds")
	}

	row := mapRows[y - 1]
	runes := []rune(row)
	runes[x - 1] = symbol
	mapRows[y - 1] = string(runes)

	return nil
}

type Rover struct {
	position Position
	direction rune
}

func NewRover(x, y int, direction rune) (Rover, error) {

	if !isValidDirection(direction) {
		return Rover{}, errors.New("invalid direction")
	}

	return Rover {
		position: Position{
			x: x,
			y: y,
		},
		direction: direction,
	}, nil
}

func isValidDirection(direction rune) bool {
	if direction == 'N' || direction == 'S' || direction == 'E' || direction == 'W' {
		return true
	}
	return false
}

func (r Rover) SetPosition(mapRows []string, x, y int, direction rune) error {
	if x <= 0 || y <= 0 {
		return errors.New("negative coordinates are not allowed")
	}

	if x > len(mapRows[0]) || y > len(mapRows) {
		return errors.New("coordinates out of bounds")
	}

	row := mapRows[y - 1]
	runes := []rune(row)
	runes[x - 1] = direction
	mapRows[y - 1] = string(runes)

	return nil
}

func (r Rover) Move(mapRows []string, command string) error {
	if !r.isValidCommand(command) {
		return errors.New("invalid command")
	}

	if command == "turn_left" {
		switch r.direction {
		case 'N':
			r.direction = 'W'
			r.SetPosition(mapRows, r.position.x, r.position.y, r.direction)
		case 'W':
			r.direction = 'S'
			r.SetPosition(mapRows, r.position.x, r.position.y, r.direction)
		case 'S':
			r.direction = 'E'
			r.SetPosition(mapRows, r.position.x, r.position.y, r.direction)
		case 'E':
			r.direction = 'N'
			r.SetPosition(mapRows, r.position.x, r.position.y, r.direction)
		}
	}

	return nil
}

func (r Rover) isValidCommand(command string) bool {
	if command == "turn_left" || command == "turn_right" || command == "move_forward" {
		return true
	}
	return false
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
