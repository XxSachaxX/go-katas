package birthday_greetings

import (
	"errors"
	"fmt"
	"strings"
)

type Friend struct {
	LastName string
	FirstName string
	BirthDate string
	Email string
}

func BuildFriends(data string) ([]Friend, error) {
	if data == "" {
		return []Friend{}, errors.New("data is empty")
	}

	friendsData := strings.Split(data, "\n")
	friends := []Friend{}
	for _, friend := range friendsData {
		parts := strings.Split(friend, ",")
		lastName, firstName, birthDate, email := parts[0], parts[1], parts[2], parts[3]
		friends = append(friends, Friend{LastName: lastName, FirstName: firstName, BirthDate: birthDate, Email: email})
	}

	return friends, nil
}

func (friend Friend) BuildBirthdayMessage() (BirthdayGreetings, error) {
	if friend.BirthDate == "" {
		return BirthdayGreetings{}, errors.New("birth date is empty")
	}

	if friend.Email == "" {
		return BirthdayGreetings{}, errors.New("email is empty")
	}

	if friend.FirstName == "" {
		return BirthdayGreetings{}, errors.New("first name is empty")
	}

	if friend.LastName == "" {
		return BirthdayGreetings{}, errors.New("last name is empty")
	}

	return BirthdayGreetings{
		title: "Happy Birthday",
		message: fmt.Sprintf("Happy birthday, dear %s %s!", friend.FirstName, friend.LastName),
	}, nil
}

type BirthdayGreetings struct {
	title string
	message string
}
