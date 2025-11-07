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
