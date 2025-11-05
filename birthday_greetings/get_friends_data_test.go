package birthday_greetings

import (
	"reflect"
	"testing"
)

func TestBuildFriendsWithoutData(t *testing.T) {
	data := ""
	_, err := BuildFriends(data)

	if err == nil {
		t.Errorf("Expected error to be raised but was not")
	}
}

func TestBuildFriendsWithSingleFriendData(t *testing.T) {
	data := "Doe,John,1982/10/08,john.doe@example.com"
	got, err := BuildFriends(data)
	want := []Friend{{FirstName: "John", LastName: "Doe", BirthDate: "1982/10/08", Email: "john.doe@example.com"}}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestBuildFriendsWithDifferentSingleFriendData(t *testing.T) {
	data := "Smith,Jane,1990/05/15,jane.smith@example.com"
	got, err := BuildFriends(data)
	want := []Friend{{FirstName: "Jane", LastName: "Smith", BirthDate: "1990/05/15", Email: "jane.smith@example.com"}}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}
