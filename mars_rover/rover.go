package mars_rover

import "errors"

type Rover struct {
	position Position
	direction rune
}

func NewRover(roverConfig *RoverConfig) (*Rover, error) {
	if roverConfig.position.x <= 0 || roverConfig.position.y <= 0 {
		return nil, errors.New("negative coordinates are not allowed")
	}

	if !isValidDirection(roverConfig.direction) {
		return nil, errors.New("invalid direction")
	}

	return &Rover {
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

func (r *Rover) SetPosition(mapRows []string, position Position, direction rune) error {
	if position.x <= 0 || position.y <= 0 {
		return errors.New("negative coordinates are not allowed")
	}

	if position.x > len(mapRows[0]) || position.y > len(mapRows) {
		return errors.New("coordinates out of bounds")
	}

	row := mapRows[position.y - 1]
	runes := []rune(row)
	runes[position.x - 1] = direction
	mapRows[position.y - 1] = string(runes)

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
