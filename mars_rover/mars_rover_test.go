package mars_rover

import (
	"reflect"
	"testing"
)

func TestMakeMapWithEqualWidthAndLength(t *testing.T) {
	createdMap, err := MakeMap(2, 2)
	expected := []string{"--", "--"}

	if err != nil {
		t.Errorf("MakeMap(2, 2) returned error: %v", err)
	}

	if createdMap == nil {
		t.Errorf("MakeMap(2, 2) returned empty slice")
	}

	if !reflect.DeepEqual(createdMap, expected) {
		t.Errorf("MakeMap(2, 2) returned %v, expected %v", createdMap, expected)
	}
}

func TestMakeMapWithUnequalWidthAndLength(t *testing.T) {
	createdMap, err := MakeMap(3, 2)
	expected := []string{"---", "---"}

	if err != nil {
		t.Errorf("MakeMap(3, 2) returned error: %v", err)
	}

	if createdMap == nil {
		t.Errorf("MakeMap(3, 2) returned empty slice")
	}

	if !reflect.DeepEqual(createdMap, expected) {
		t.Errorf("MakeMap(3, 2) returned %v, expected %v", createdMap, expected)
	}
}

func TestMakeMapWithZeroWidth(t *testing.T) {
	_, err := MakeMap(0, 2)

	if err == nil {
		t.Errorf("Should not allow zero width")
	}
}

func TestMakeMapWithNegativeWidth(t *testing.T) {
	_, err := MakeMap(-1, 2)

	if err == nil {
		t.Errorf("Should not allow negative width")
	}
}

func TestMakeMapWithZeroHeight(t *testing.T) {
	_, err := MakeMap(2, 0)

	if err == nil {
		t.Errorf("Should not allow zero height")
	}
}

func TestMakeMapWithNegativeHeight(t *testing.T) {
	_, err := MakeMap(2, -1)

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
	}
	obstaclesConfigs := ObstaclesConfigs{obstacles: []Obstacle{obstacle}}

	createdMap, _ := CreateMap(&mapConfig, &roverConfig, obstaclesConfigs)
	expectedRows := []string{"N--", "-O-", "---"}

	if !reflect.DeepEqual(createdMap.rows, expectedRows) {
		t.Errorf("Expected rows %v, got %v", expectedRows, createdMap.rows)
	}

}

func TestSetRoverPositionWithNegativeX(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(-1, 1, 'N')
	err := rover.SetPosition(createdMap, -1, 1, 'N')

	if err == nil {
		t.Errorf("Should not allow negative x position")
	}
}

func TestSetRoverPositionWithNegativeY(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, -1, 'N')
	err := rover.SetPosition(createdMap, 1, -1, 'N')

	if err == nil {
		t.Errorf("Should not allow negative y position")
	}
}

func TestRoverPositionWithOutOfBoundsX(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'N')
	err := rover.SetPosition(createdMap, 3, 1, 'N')

	if err == nil {
		t.Errorf("Should not allow x position out of bounds")
	}
}

func TestRoverPositionWithOutOfBoundsY(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'N')
	err := rover.SetPosition(createdMap, 1, 3, 'N')

	if err == nil {
		t.Errorf("Should not allow y position out of bounds")
	}
}

func TestRoverCreationWithInvalidDirection(t *testing.T) {
	_, err := NewRover(1, 1, 'X')

	if err == nil {
		t.Errorf("Should not allow invalid direction")
	}
}

func TestRoverPositionWithValidCoordinates(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'N')
	err := rover.SetPosition(createdMap, 1, 1, 'N')
	expectedMap := []string{"N-", "--"}

	if err != nil {
		t.Errorf("SetRoverPosition(1, 1) returned error: %v", err)
	}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestRoverPositionWithOtherSetOfValidCoordinates(t *testing.T) {
	createdMap, _ := MakeMap(10,5)
	rover, _ := NewRover(5, 5, 'N')
	err := rover.SetPosition(createdMap, 5, 5, 'N')
	expectedMap := []string{"----------", "----------","----------","----------","----N-----"}

	if err != nil {
		t.Errorf("SetRoverPosition(5, 5) returned error: %v", err)
	}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestMoveRoverWithInvalidCommand(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'N')
	err := rover.Move(createdMap, "jump")

	if err == nil {
		t.Errorf("Expected error, invalid move")
	}
}

func TestTurnRoverPointingNorthToTheLeft(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'N')
	err := rover.Move(createdMap, "turn_left")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMap := []string{"W-", "--"}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestTurnRoverPointingNorthToTheRight(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'N')
	err := rover.Move(createdMap, "turn_right")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMap := []string{"E-", "--"}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestTurnRoverPointingSouthToTheLeft(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'S')
	err := rover.Move(createdMap, "turn_left")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMap := []string{"E-", "--"}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestTurnRoverPointingSouthToTheRight(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'S')
	err := rover.Move(createdMap, "turn_right")

	if err != nil {
		t.Errorf("TurnRoverRight() returned error: %v", err)
	}

	expectedMap := []string{"W-", "--"}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestTurnRoverPointingEastToTheLeft(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'E')
	err := rover.Move(createdMap, "turn_left")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMap := []string{"N-", "--"}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestTurnRoverPointingEastToTheRight(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'E')
	err := rover.Move(createdMap, "turn_right")

	if err != nil {
		t.Errorf("TurnRoverRight() returned error: %v", err)
	}

	expectedMap := []string{"S-", "--"}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestTurnRoverPointingWestToTheLeft(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'W')
	err := rover.Move(createdMap, "turn_left")

	if err != nil {
		t.Errorf("TurnRoverLeft() returned error: %v", err)
	}

	expectedMap := []string{"S-", "--"}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestTurnRoverPointingWestToTheRight(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'W')
	err := rover.Move(createdMap, "turn_right")

	if err != nil {
		t.Errorf("TurnRoverRight() returned error: %v", err)
	}

	expectedMap := []string{"N-", "--"}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestTryMoveRoverForwardOutOfMapOnXAxis(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'W')
	err := rover.Move(createdMap, "move_forward")

	if err == nil {
		t.Errorf("Should not be able to move rover forward out of map")
	}
}

func TestTryMoveRoverForwardOutOfMapOnYAxis(t *testing.T) {
	createdMap, _ := MakeMap(2,1)
	rover, _ := NewRover(1, 1, 'N')
	err := rover.Move(createdMap, "move_forward")

	if err == nil {
		t.Errorf("Should not be able to move rover forward out of map")
	}
}

func TestTryMoveRoverForwardIntoObstacleFromEast(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'E')
	rover.SetPosition(createdMap, 1, 1, 'E')

	obstacle, _ := NewObstacle(2, 1)
	obstacle.SetPosition(createdMap, 2, 1, 'O')

	err := rover.Move(createdMap, "move_forward")
	if err == nil {
		t.Errorf("Should not be able to move rover forward into obstacle")
	}
}

func TestTryMoveRoverForwardIntoObstacleFromWest(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(2, 1, 'W')
	rover.SetPosition(createdMap, 2, 1, 'W')

	obstacle, _ := NewObstacle(1, 1)
	obstacle.SetPosition(createdMap, 1, 1, 'O')

	err := rover.Move(createdMap, "move_forward")
	if err == nil {
		t.Errorf("Should not be able to move rover forward into obstacle")
	}
}

func TestTryMoveRoverForwardIntoObstacleFromNorth(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 1, 'N')
	rover.SetPosition(createdMap, 1, 1, 'N')

	obstacle, _ := NewObstacle(1, 2)
	obstacle.SetPosition(createdMap, 1, 2, 'O')

	err := rover.Move(createdMap, "move_forward")
	if err == nil {
		t.Errorf("Should not be able to move rover forward into obstacle")
	}
}

func TestTryMoveRoverForwardIntoObstacleFromSouth(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	rover, _ := NewRover(1, 2, 'S')
	rover.SetPosition(createdMap, 1, 2, 'S')

	obstacle, _ := NewObstacle(1, 1)
	obstacle.SetPosition(createdMap, 1, 1, 'O')

	err := rover.Move(createdMap, "move_forward")
	if err == nil {
		t.Errorf("Should not be able to move rover forward into obstacle")
	}
}

func TestCreateObstacleWithNegativeX(t *testing.T) {
	_, err := NewObstacle(0, 1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestCreateObstacleWithNegativeY(t *testing.T) {
	_, err := NewObstacle(1, 0)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPositionObstacleOnMapWithOutOfBoundsX(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	obstacle, _ := NewObstacle(3, 1)
	err := obstacle.SetPosition(createdMap, 3, 1, 'O')

	if err == nil {
		t.Errorf("SetObstaclePosition(3, 1) should return an error")
	}
}

func TestPositionObstacleOnMapWithOutOfBoundsY(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	obstacle, _ := NewObstacle(1, 3)
	err := obstacle.SetPosition(createdMap, 1, 3, 'O')

	if err == nil {
		t.Errorf("SetObstaclePosition(1, 3) should return an error")
	}
}

func TestPositionObstacleWithValidCoordinates(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	obstacle, _ := NewObstacle(1, 1)
	err := obstacle.SetPosition(createdMap, 1, 1, 'O')
	expectedMap := []string{"O-", "--"}

	if err != nil {
		t.Errorf("SetObstaclePosition(1, 1) returned error: %v", err)
	}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}
