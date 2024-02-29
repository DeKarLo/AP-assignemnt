package repository

import "architecture_go/services/contact/cmd/internal/domain"

type ContactRepository interface {
	CreateContact(contact *domain.Contact) error
	GetContactByID(id int) (*domain.Contact, error)
	UpdateContact(contact *domain.Contact) error
	DeleteContact(id int) error
}

type GroupRepository interface {
	CreateGroup(group *domain.Group) error
	GetGroupByID(id int) (*domain.Group, error)
	AddContactToGroup(contactID, groupID int) error
}
