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

func TestSetPositionWithNegativeX(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	err := SetRoverPosition(createdMap, -1, 1)

	if err == nil {
		t.Errorf("Should not allow negative x position")
	}
}

func TestSetRoverPositionWithNegativeY(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	err := SetRoverPosition(createdMap, 1, -1)

	if err == nil {
		t.Errorf("Should not allow negative y position")
	}
}

func TestRoverPositionWithOutOfBoundsX(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	err := SetRoverPosition(createdMap, 3, 1)

	if err == nil {
		t.Errorf("Should not allow x position out of bounds")
	}
}

func TestRoverPositionWithOutOfBoundsY(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	err := SetRoverPosition(createdMap, 1, 3)

	if err == nil {
		t.Errorf("Should not allow y position out of bounds")
	}
}

func TestRoverPositionWithValidCoordinates(t *testing.T) {
	createdMap, _ := MakeMap(2,2)
	err := SetRoverPosition(createdMap, 1, 1)
	expectedMap := []string{"X-", "--"}

	if err != nil {
		t.Errorf("SetRoverPosition(1, 1) returned error: %v", err)
	}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}

func TestRoverPositionWithOtherSetOfValidCoordinates(t *testing.T) {
	createdMap, _ := MakeMap(10,5)
	err := SetRoverPosition(createdMap, 5, 5)
	expectedMap := []string{"----------", "----------","----------","----------","----X-----"}

	if err != nil {
		t.Errorf("SetRoverPosition(5, 5) returned error: %v", err)
	}

	if !reflect.DeepEqual(createdMap, expectedMap) {
		t.Errorf("Expected map %v, got %v", expectedMap, createdMap)
	}
}
