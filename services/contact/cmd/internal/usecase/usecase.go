// usecase.go
package usecase

import "architecture_go/services/contact/cmd/internal/domain"

type Usecase interface {
	CreateContact(firstName, middleName, lastName, phoneNumber string) (*domain.Contact, error)
	GetContactByID(id int) (*domain.Contact, error)
	UpdateContact(id int, firstName, middleName, lastName, phoneNumber string) (*domain.Contact, error)
	DeleteContact(id int) error

	CreateGroup(name string) (*domain.Group, error)
	GetGroupByID(id int) (*domain.Group, error)

	AddContactToGroup(contactID, groupID int) error
}
