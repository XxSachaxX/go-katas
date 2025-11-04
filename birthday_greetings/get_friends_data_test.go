package birthday_greetings

import (
	"reflect"
	"testing"
)

func TestGetFriendsDataReturnType(t *testing.T) {
	data := ""
	got := BuildFriends(data)
	want := reflect.TypeOf([]Friend{})

	if reflect.TypeOf(got) != want {
		t.Errorf("Return value should be a slice of friends")
	}
}

func TestBuildFriendsWithoutData(t *testing.T) {
	data := ""
	got := BuildFriends(data)
	want := []Friend{}

	if len(got) != len(want) {
		t.Errorf("Expected %d friends, got %d", len(want), len(got))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
