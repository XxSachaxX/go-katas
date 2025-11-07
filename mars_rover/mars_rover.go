package mars_rover

import (
	"errors"
	"strings"
)

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

func SetRoverPosition(mapRows []string, x, y int) error {
	if x < 0 || y < 0 {
		return errors.New("negative coordinates are not allowed")
	}

	return nil
}
