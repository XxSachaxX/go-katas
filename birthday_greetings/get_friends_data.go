package birthday_greetings

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Friend struct {
	LastName string
	FirstName string
	BirthDate string
	Email string
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

func (greetings BirthdayGreetings) Send() error {
	if greetings.title == "" {
		return errors.New("title is empty")
	}

	if greetings.message == "" {
		return errors.New("message is empty")
	}

	return nil
}

type BirthdayGreetings struct {
	title string
	message string
}

type FriendsRepository interface {
	GetFriends() ([]Friend, error)
}

type TextFileFriendsRepository struct {
	path string
}

func (repo TextFileFriendsRepository) GetFriends() ([]Friend, error) {
	data, err := os.Open(repo.path)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	csv := csv.NewReader(data)
	csv.TrimLeadingSpace = true
	csv.FieldsPerRecord = 4

	rows, err := csv.ReadAll()
		if err != nil {
			return nil, err
		}

	friends := make([]Friend, 0, len(rows))

	for _, rec := range rows {
			friends = append(friends, Friend{
				LastName:  strings.TrimSpace(rec[0]),
				FirstName: strings.TrimSpace(rec[1]),
				BirthDate: strings.TrimSpace(rec[2]),
				Email:     strings.TrimSpace(rec[3]),
			})
		}

	return friends, nil
}
