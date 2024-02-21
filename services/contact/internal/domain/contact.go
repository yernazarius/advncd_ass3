package domain

import (
	"fmt"
	"regexp"
	"strings"
)

type Contact struct {
	ID          int
	FirstName   string
	MiddleName  string
	LastName    string
	PhoneNum string
}

func NewContact(id int, phoneNum, firstName, middleName, lastName string) (*Contact, error) {
	if !isVaildPhoneNumber(phoneNum) {
		return nil, fmt.Errorf("Invalid phone num %s", phoneNum)
	}

	return &Contact{
		ID:          id,
		FirstName:   firstName,
		MiddleName:  middleName,
		LastName:    lastName,
		PhoneNum:    phoneNum,
	}, nil
}

func (c *Contact) FullName() string {
	return strings.Join([]string{c.FirstName, c.MiddleName, c.LastName}, " ")
}

func isVaildPhoneNumber(phoneNum string) bool {
	matched, _ := regexp.MatchString("^[0-9]+$", phoneNum)
	return matched
}
