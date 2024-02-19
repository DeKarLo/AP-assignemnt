package domain

import (
	"errors"
	"strconv"
)

type Contact struct {
	ID          int
	FirstName   string
	MiddleName  string
	LastName    string
	PhoneNumber int
}

func NewContact(id int, firstName string, middleName string, lastName string, phoneNumber int) (*Contact, error) {
	if !isPhoneNumberValid(phoneNumber) {
		return nil, errors.New("invalid phone number")
	}

	contact := &Contact{
		ID:          id,
		FirstName:   firstName,
		MiddleName:  middleName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}

	return contact, nil
}

func (c *Contact) FullName() string {
	return c.FirstName + " " + c.MiddleName + " " + c.LastName
}

func isPhoneNumberValid(phoneNumber int) bool {
	phoneNumberStr := strconv.Itoa(phoneNumber)
	_, err := strconv.Atoi(phoneNumberStr)
	return err == nil
}
