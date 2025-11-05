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

func TestBuildFriendsWithMultipleFriendsData(t *testing.T) {
	data := "Smith,Jane,1990/05/15,jane.smith@example.com\nDoe,John,1982/10/08,john.doe@example.com"
	got, err := BuildFriends(data)
	want := []Friend{
		{FirstName: "Jane", LastName: "Smith", BirthDate: "1990/05/15", Email: "jane.smith@example.com"},
		{FirstName: "John", LastName: "Doe", BirthDate: "1982/10/08", Email: "john.doe@example.com"},
	}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestBuildBirthdayMessageWithDataLessFriend(t *testing.T) {
	friend := Friend{FirstName: "", LastName: "", BirthDate: "", Email: ""}

	_, err := friend.BuildBirthdayMessage()

	if err == nil {
		t.Errorf("Expected error to be raised but was not")
	}
}

func TestBuildBirthdayMessageWithoutFriendFirstName(t *testing.T) {
	friend := Friend{FirstName: "", LastName: "Smith", BirthDate: "1990/05/15", Email: "jane.smith@example.com"}

	_, err := friend.BuildBirthdayMessage()

	if err == nil {
		t.Errorf("Expected error to be raised but was not")
	}

	if err.Error() != "first name is empty" {
		t.Errorf("Expected error message to be 'friend's first name is empty' but got '%v'", err.Error())
	}
}

func TestBuildBirthdayMessageWithoutFriendLastName(t *testing.T) {
	friend := Friend{FirstName: "Jane", LastName: "", BirthDate: "1990/05/15", Email: "jane.smith@example.com"}

	_, err := friend.BuildBirthdayMessage()

	if err == nil {
		t.Errorf("Expected error to be raised but was not")
	}

	if err.Error() != "last name is empty" {
		t.Errorf("Expected error message to be 'friend's first name is empty' but got '%v'", err.Error())
	}
}

func TestBuildBirthdayMessageWithoutFriendBirthdate(t *testing.T) {
	friend := Friend{FirstName: "Jane", LastName: "Doe", BirthDate: "", Email: "jane.smith@example.com"}

	_, err := friend.BuildBirthdayMessage()

	if err == nil {
		t.Errorf("Expected error to be raised but was not")
	}

	if err.Error() != "birth date is empty" {
		t.Errorf("Expected error message to be 'friend's first name is empty' but got '%v'", err.Error())
	}
}

func TestBuildBirthdayMessageWithoutFriendEmail(t *testing.T) {
	friend := Friend{FirstName: "Jane", LastName: "Doe", BirthDate: "1990/05/15", Email: ""}

	_, err := friend.BuildBirthdayMessage()

	if err == nil {
		t.Errorf("Expected error to be raised but was not")
	}

	if err.Error() != "email is empty" {
		t.Errorf("Expected error message to be 'friend's first name is empty' but got '%v'", err.Error())
	}
}

func TestBuildBirthdayMessageWithAFriendWithAllData(t *testing.T) {
	friend := Friend{FirstName: "Jane", LastName: "Doe", BirthDate: "1990/05/15", Email: "jane.smith@example.com"}

	got, err := friend.BuildBirthdayMessage()
	want := BirthdayGreetings{
		title: "Happy Birthday",
		message: "Happy birthday, dear Jane Doe!",
	}

	if err != nil {
		t.Errorf("Expected no error but got '%v'", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected '%v' but got '%v'", want, got)
	}
}

func TestSendBirthdayGreetings(t *testing.T) {
	birthdayGreetings := BirthdayGreetings{title: "Happy Birthday", message: "Happy birthday, dear Jane Doe!"}
	err := birthdayGreetings.Send()

	if err != nil {
		t.Errorf("Expected no error but got '%v'", err.Error())
	}
}
