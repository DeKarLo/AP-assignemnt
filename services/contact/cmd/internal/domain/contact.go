package domain

import (
	"errors"
	"strconv"
)

type Contact struct {
	id          int
	firstName   string
	middleName  string
	lastName    string
	phoneNumber int
}

func NewContact(id int, firstName string, middleName string, lastName string, phoneNumber int) (*Contact, error) {
	if !isPhoneNumberValid(phoneNumber) {
		return nil, errors.New("invalid phone number")
	}

	contact := &Contact{
		id:          id,
		firstName:   firstName,
		middleName:  middleName,
		lastName:    lastName,
		phoneNumber: phoneNumber,
	}

	return contact, nil
}

func (c *Contact) FullName() string {
	return c.firstName + " " + c.middleName + " " + c.lastName
}

func isPhoneNumberValid(phoneNumber int) bool {
	phoneNumberStr := strconv.Itoa(phoneNumber)
	_, err := strconv.Atoi(phoneNumberStr)
	return err == nil
}
