package mars_rover

import (
	"errors"
	"strings"
)

type Map struct {
	rows []string
	rover *Rover
	obstacles []*Obstacle
}

type MapConfig struct {
	width int
	height int
}

type RoverConfig struct {
	position Position
	direction rune
}

type ObstaclesConfigs struct {
	obstacles []Obstacle
}

type Position struct {
	x int
	y int
	SetPosition func(mapRows []string, x, y int, symbol rune) error
}

type ObstacleConfig struct {
	position Position
}

type Obstacle struct {
	position Position
}

type Rover struct {
	position Position
	direction rune
}

func (m *Map) SetRover(x, y int, direction rune) error {
	if x <= 0 || y <= 0 {
		return errors.New("negative coordinates are not allowed")
	}

	if x > len(m.rows[0]) || y > len(m.rows) {
		return errors.New("coordinates out of bounds")
	}

	row := m.rows[y - 1]
	runes := []rune(row)
	runes[x - 1] = direction
	m.rows[y - 1] = string(runes)

	return nil
}

func NewObstacle(obstacleConfig ObstacleConfig) (*Obstacle, error) {
	if obstacleConfig.position.x <= 0 || obstacleConfig.position.y <= 0 {
		return nil, errors.New("invalid coordinates")
	}

	return &Obstacle {
		position: obstacleConfig.position,
	}, nil
}

func (obstacle *Obstacle) SetPosition(mapRows []string, x, y int, symbol rune) error {
	if x > len(mapRows[0]) || y > len(mapRows) {
		return errors.New("coordinates out of bounds")
	}

	row := mapRows[y - 1]
	runes := []rune(row)
	runes[x - 1] = symbol
	mapRows[y - 1] = string(runes)

	return nil
}

func NewRover(roverConfig *RoverConfig) (Rover, error) {
	if roverConfig.position.x <= 0 || roverConfig.position.y <= 0 {
		return Rover{}, errors.New("negative coordinates are not allowed")
	}

	if !isValidDirection(roverConfig.direction) {
		return Rover{}, errors.New("invalid direction")
	}

	return Rover {
		position: Position{
			x: roverConfig.position.x,
			y: roverConfig.position.y,
		},
		direction: roverConfig.direction,
	}, nil
}

func isValidDirection(direction rune) bool {
	if direction == 'N' || direction == 'S' || direction == 'E' || direction == 'W' {
		return true
	}
	return false
}

func (r *Rover) SetPosition(mapRows []string, x, y int, direction rune) error {
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

func (m *Map) MoveRover(command string) error {
	if !m.rover.isValidCommand(command) {
		return errors.New("invalid command")
	}

	if command == "turn_left" {
		m.rover.TurnLeft()
	}

	if command == "turn_right" {
		m.rover.TurnRight()
	}

	if command == "move_forward" {
		err := m.rover.MoveForward(m.rows)
		if err != nil {
			return err
		}
	}

	m.rover.SetPosition(m.rows, m.rover.position.x, m.rover.position.y, m.rover.direction)
	return nil
}

func (r *Rover) isValidCommand(command string) bool {
	if command == "turn_left" || command == "turn_right" || command == "move_forward" {
		return true
	}
	return false
}

func (r *Rover) TurnLeft() error {
	switch r.direction {
	case 'N':
		r.direction = 'W'
	case 'W':
		r.direction = 'S'
	case 'S':
		r.direction = 'E'
	case 'E':
		r.direction = 'N'
	}

	return nil
}

func (r *Rover) TurnRight() error {
	switch r.direction {
	case 'N':
		r.direction = 'E'
	case 'E':
		r.direction = 'S'
	case 'S':
		r.direction = 'W'
	case 'W':
		r.direction = 'N'
	}

	return nil
}

func (r *Rover) MoveForward(mapRows []string) error {
	switch r.direction {
	case 'N':
		if r.position.y + 1 > len(mapRows) {
			return errors.New("coordinates out of bounds")
		}


		newX := r.position.x
		newY := r.position.y + 1
		if mapRows[newY - 1][newX - 1] == 'O' {
			return errors.New("obstacle detected")
		}

		r.position.y++
	case 'S':
		if r.position.y - 1 <= 0 {
			return errors.New("coordinates out of bounds")
		}
		newX := r.position.x
		newY := r.position.y - 1
		if mapRows[newY - 1][newX - 1] == 'O' {
			return errors.New("obstacle detected")
		}

		r.position.y--
	case 'E':
		if r.position.x + 1 > len(mapRows[0]) {
			return errors.New("coordinates out of bounds")
		}

		newX := r.position.x + 1
		newY := r.position.y
		if mapRows[newY - 1][newX - 1] == 'O' {
			return errors.New("obstacle detected")
		}

		r.position.x++
	case 'W':
		if r.position.x - 1 <= 0 {
			return errors.New("coordinates out of bounds")
		}

		newX := r.position.x - 1
		newY := r.position.y
		if mapRows[newY - 1][newX - 1] == 'O' {
			return errors.New("obstacle detected")
		}

		r.position.x--
	}

	return nil
}

func CreateMap(mapConfig *MapConfig, roverConfig *RoverConfig, obstacleConfigs ObstaclesConfigs) (*Map, error) {
	width, height := mapConfig.width, mapConfig.height

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

	if roverConfig.position.x > len(mapRows[0]) || roverConfig.position.y > len(mapRows) {
		return nil, errors.New("coordinates out of bounds")
	}

	if roverConfig.position.x <= 0 || roverConfig.position.y <= 0 {
		return nil, errors.New("coordinates out of bounds")
	}


	rover, _ := NewRover(roverConfig)
	obstacles := make([]*Obstacle, len(obstacleConfigs.obstacles))
	for i, obstacle := range obstacleConfigs.obstacles {
		config := ObstacleConfig{position: obstacle.position}
		newObstacle, err := NewObstacle(config)
		if err != nil {
			return nil, err
		}
		obstacles[i] = newObstacle
	}

	rover.SetPosition(mapRows, rover.position.x, rover.position.y, rover.direction)
	for _, obstacle := range obstacles {
		if obstacle.position.x > len(mapRows[0]) || obstacle.position.y > len(mapRows) {
			return nil, errors.New("coordinates out of bounds")
		}

		if obstacle.position.x > 0 && obstacle.position.y > 0 {
			obstacle.SetPosition(mapRows, obstacle.position.x, obstacle.position.y, 'O')
		}
	}

	return &Map{
		rows: mapRows,
		rover: &rover,
		obstacles: obstacles,
	}, nil
}
