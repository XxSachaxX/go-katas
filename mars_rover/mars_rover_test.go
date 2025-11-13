package mars_rover

import (
	"reflect"
	"testing"
)

func TestMakeMapWithEqualWidthAndLength(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}
	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}
	obstacle := Obstacle{
		position: Position{
			x: 2,
			y: 2,
		},
		symbol: 'O',
	}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)

	expectedMapRows := []string{"N-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected rows %v, got %v", expectedMapRows, createdMap.rows)
	}

	if err != nil {
		t.Errorf("Error when creating map: %v", err)
	}

	if createdMap == nil {
		t.Errorf("Map is empty")
	}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Map rows are not as expected: returned %v, expected %v", createdMap.rows, expectedMapRows)
	}
}

func TestMakeMapWithZeroWidth(t *testing.T) {
	mapConfig := MapConfig{
		width:  0,
		height: 2,
	}
	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}
	obstacle := Obstacle{
		position: Position{
			x: 2,
			y: 2,
		},
		symbol: 'O',
	}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	_, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)

	if err == nil {
		t.Errorf("Should not allow zero width")
	}
}

func TestMakeMapWithNegativeWidth(t *testing.T) {
	mapConfig := MapConfig{
		width:  -1,
		height: 2,
	}
	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}
	obstacle := Obstacle{
		position: Position{
			x: 2,
			y: 2,
		},
		symbol: 'O',
	}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	_, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)

	if err == nil {
		t.Errorf("Should not allow negative width")
	}
}

func TestMakeMapWithZeroHeight(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 0,
	}
	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}
	obstacle := Obstacle{
		position: Position{
			x: 2,
			y: 2,
		},
		symbol: 'O',
	}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	_, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)


	if err == nil {
		t.Errorf("Should not allow zero height")
	}
}

func TestMakeMapWithNegativeHeight(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: -1,
	}
	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}
	obstacle := Obstacle{
		position: Position{
			x: 2,
			y: 2,
		},
		symbol: 'O',
	}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	_, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)


	if err == nil {
		t.Errorf("Should not allow negative height")
	}
}

func TestSetRoverOnMapWithValidCoordinates(t *testing.T) {
	mapConfig := MapConfig{
		width:  3,
		height: 3,
	}
	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}
	obstacle := Obstacle{
		position: Position{
			x: 2,
			y: 2,
		},
		symbol: 'O',
	}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	expectedRows := []string{"N--", "-O-", "---"}

	if !reflect.DeepEqual(createdMap.rows, expectedRows) {
		t.Errorf("Expected rows %v, got %v", expectedRows, createdMap.rows)
	}
}

func TestCreateRoverWithNegativeX(t *testing.T) {
	roverConfig := RoverConfig{
		position: Position{
			x: -1,
			y: 1,
		},
		direction: 'N',
	}

	_, err := NewRover(&roverConfig)


	if err == nil {
		t.Errorf("Should not allow negative x position")
	}
}

func TestCreateRoverWithNegativeY(t *testing.T) {

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: -1,
		},
		direction: 'N',
	}

	_, err := NewRover(&roverConfig)

	if err == nil {
		t.Errorf("Should not allow negative y position")
	}
}

func TestSetRoverOnMapWithOutOfBoundsX(t *testing.T) {
	mapConfig := MapConfig{
		width:  3,
		height: 3,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 0,
			y: 1,
		},
		direction: 'N',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	_, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)


	if err == nil {
		t.Errorf("Should not allow x position out of bounds")
	}
}

func TestSetRoverOnMapWithOutOfBoundsY(t *testing.T) {
	mapConfig := MapConfig{
		width:  3,
		height: 3,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 0,
		},
		direction: 'N',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	_, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)


	if err == nil {
		t.Errorf("Should not allow x position out of bounds")
	}
}

func TestRoverCreationWithInvalidDirection(t *testing.T) {
	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'X',
	}

	_, err := NewRover(&roverConfig)

	if err == nil {
		t.Errorf("Should not allow invalid direction")
	}
}

func TestMoveRoverWithInvalidCommand(t *testing.T) {
	mapConfig := MapConfig{
		width:  3,
		height: 3,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("jump")

	if err == nil {
		t.Errorf("Expected error, invalid move")
	}
}


func TestTurnRoverPointingNorthToTheLeft(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("turn_left")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMapRows := []string{"W-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}

func TestTurnRoverPointingNorthToTheRight(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("turn_right")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMapRows := []string{"E-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}

func TestTurnRoverPointingSouthToTheLeft(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'S',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("turn_left")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMapRows := []string{"E-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}

func TestTurnRoverPointingSouthToTheRight(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'S',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("turn_right")

	if err != nil {
		t.Errorf("TurnRoverRight() returned error: %v", err)
	}

	expectedMapRows := []string{"W-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}

func TestTurnRoverPointingEastToTheLeft(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'E',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("turn_left")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMapRows := []string{"N-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}

func TestTurnRoverPointingEastToTheRight(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'E',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("turn_right")

	if err != nil {
		t.Errorf("TurnRoverRight() returned error: %v", err)
	}

	expectedMapRows := []string{"S-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}

func TestTurnRoverPointingWestToTheLeft(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'W',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("turn_left")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMapRows := []string{"S-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}

func TestTurnRoverPointingWestToTheRight(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'W',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("turn_right")

	if err != nil {
		t.Errorf("TurnRoverRight() returned error: %v", err)
	}

	expectedMapRows := []string{"N-", "-O"}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}

func TestTryMoveRoverForwardOutOfMapOnXAxis(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'W',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("move_forward")

	if err == nil {
		t.Errorf("Should not be able to move rover forward out of map")
	}
}

func TestTryMoveRoverForwardOutOfMapOnYAxis(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 1,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 1}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("move_forward")

	if err == nil {
		t.Errorf("Should not be able to move rover forward out of map")
	}
}

func TestTryMoveRoverForwardIntoObstacleFromEast(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 1,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}

	obstacle := Obstacle{position: Position{x: 2, y: 1}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("move_forward")

	if err == nil {
		t.Errorf("Should not be able to move rover forward into obstacle")
	}
}

func TestTryMoveRoverForwardIntoObstacleFromWest(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 2,
			y: 1,
		},
		direction: 'W',
	}

	obstacle := Obstacle{position: Position{x: 1, y: 1}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("move_forward")

	if err == nil {
		t.Errorf("Should not be able to move rover forward into obstacle")
	}
}

func TestTryMoveRoverForwardIntoObstacleFromNorth(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'N',
	}

	obstacle := Obstacle{position: Position{x: 1, y: 2}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("move_forward")

	if err == nil {
		t.Errorf("Should not be able to move rover forward into obstacle")
	}
}

func TestTryMoveRoverForwardIntoObstacleFromSouth(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 2,
		},
		direction: 'S',
	}

	obstacle := Obstacle{position: Position{x: 1, y: 1}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	err := createdMap.MoveRover("move_forward")

	if err == nil {
		t.Errorf("Should not be able to move rover forward into obstacle")
	}
}

func TestCreateObstacleWithNegativeX(t *testing.T) {
	obstacleConfig := ObstacleConfig{position: Position{x: -1, y: 1}, symbol: 'O'}
	_, err := NewObstacle(obstacleConfig)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestCreateObstacleWithNegativeY(t *testing.T) {
	obstacleConfig := ObstacleConfig{position: Position{x: 1, y: -1}, symbol: 'O'}
	_, err := NewObstacle(obstacleConfig)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPositionObstacleOnMapWithOutOfBoundsX(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'S',
	}

	obstacle := Obstacle{position: Position{x: 3, y: 1}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	_, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)

	if err == nil {
		t.Errorf("out of bound x for obstacle should return an error")
	}
}

func TestPositionObstacleOnMapWithOutOfBoundsY(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 1,
			y: 1,
		},
		direction: 'S',
	}

	obstacle := Obstacle{position: Position{x: 1, y: 3}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	_, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)

	if err == nil {
		t.Errorf("out of bound y for obstacle should return an error")
	}
}

func TestPositionObstacleWithValidCoordinates(t *testing.T) {
	mapConfig := MapConfig{
		width:  2,
		height: 2,
	}

	roverConfig := RoverConfig{
		position: Position{
			x: 2,
			y: 2,
		},
		direction: 'S',
	}

	obstacle := Obstacle{position: Position{x: 1, y: 1}, symbol: 'O'}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, err := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)


	expectedMapRows := []string{"O-", "-S"}

	if err != nil {
		t.Errorf("SetObstaclePosition(1, 1) returned error: %v", err)
	}

	if !reflect.DeepEqual(createdMap.rows, expectedMapRows) {
		t.Errorf("Expected map %v, got %v", expectedMapRows, createdMap.rows)
	}
}
