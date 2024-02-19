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

type contactRepository struct{}

func (r *contactRepository) CreateContact(contact *domain.Contact) error {
	return nil
}

func (r *contactRepository) GetContactByID(id int) (*domain.Contact, error) {
	return nil, nil
}

func (r *contactRepository) UpdateContact(contact *domain.Contact) error {
	return nil
}

func (r *contactRepository) DeleteContact(id int) error {
	return nil
}

type groupRepository struct{}

func (r *groupRepository) CreateGroup(group *domain.Group) error {
	return nil
}

func (r *groupRepository) GetGroupByID(id int) (*domain.Group, error) {
	return nil, nil
}

func (r *groupRepository) AddContactToGroup(contactID, groupID int) error {
	return nil
}

func NewContactRepository() ContactRepository {
	return &contactRepository{}
}

func NewGroupRepository() GroupRepository {
	return &groupRepository{}
}
