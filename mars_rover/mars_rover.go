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
	symbol rune
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

	if !roverConfig.isValidPosition(mapRows) {
		return nil, errors.New("coordinates out of bounds")
	}

	rover, roverErr := NewRover(roverConfig)
	if roverErr != nil {
		return nil, roverErr
	}
	rover.SetPosition(mapRows, rover.position.x, rover.position.y, rover.direction)

	obstacles, obstaclesErr := createObstacles(obstacleConfigs)
	if obstaclesErr != nil {
		return nil, obstaclesErr
	}

	for _, obstacle := range obstacles {
		if obstacle.position.x > len(mapRows[0]) || obstacle.position.y > len(mapRows) {
			return nil, errors.New("coordinates out of bounds")
		}

		if mapRows[obstacle.position.y - 1][obstacle.position.x - 1] != '-' {
			return nil, errors.New("obstacle already exists at position")
		}

		if obstacle.position.x > 0 && obstacle.position.y > 0 {
			obstacle.SetPosition(mapRows, obstacle.position.x, obstacle.position.y, obstacle.symbol)
		}
	}

	return &Map{
		rows: mapRows,
		rover: rover,
		obstacles: obstacles,
	}, nil
}

func (rc *RoverConfig) isValidPosition(mapRows []string) bool {
	if rc.position.x > len(mapRows[0]) || rc.position.y > len(mapRows) {
		return false
	}

	if rc.position.x <= 0 || rc.position.y <= 0 {
		return false
	}

	return true
}

func createObstacles(obstacleConfigs ObstaclesConfigs) ([]*Obstacle, error) {
	obstacles := make([]*Obstacle, len(obstacleConfigs.obstacles))

	for i, obstacle := range obstacleConfigs.obstacles {
		config := ObstacleConfig{position: obstacle.position, symbol: obstacle.symbol}
		newObstacle, err := NewObstacle(config)
		if err != nil {
			return nil, err
		}
		obstacles[i] = newObstacle
	}

	return obstacles, nil
}
