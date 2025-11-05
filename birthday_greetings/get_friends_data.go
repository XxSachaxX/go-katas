package birthday_greetings

import (
	"errors"
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

	friends := []Friend{}
	parts := strings.Split(data, ",")
	lastName, firstName, birthDate, email := parts[0], parts[1], parts[2], parts[3]
	friends = append(friends, Friend{LastName: lastName, FirstName: firstName, BirthDate: birthDate, Email: email})

	return friends, nil
}
