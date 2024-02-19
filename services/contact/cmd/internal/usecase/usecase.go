// usecase.go
package usecase

import (
	"architecture_go/services/contact/cmd/internal/domain"
	"architecture_go/services/contact/cmd/internal/repository"
)

type Usecase interface {
	CreateContact(firstName, middleName, lastName, phoneNumber string) error
	GetContactByID(id int) (*domain.Contact, error)
	UpdateContact(id int, firstName, middleName, lastName, phoneNumber string) error
	DeleteContact(id int) error

	CreateGroup(name string) error
	GetGroupByID(id int) (*domain.Group, error)

	AddContactToGroup(contactID, groupID int) error
}

type ContactUsecase struct {
	contactRepo repository.ContactRepository
	groupRepo   repository.GroupRepository
}

func NewContactUsecase(contactRepo repository.ContactRepository, groupRepo repository.GroupRepository) *ContactUsecase {
	return &ContactUsecase{
		contactRepo: contactRepo,
		groupRepo:   groupRepo,
	}
}

func (uc *ContactUsecase) CreateContact(firstName, middleName, lastName string, phoneNumber int) error {

	contact := &domain.Contact{
		FirstName:   firstName,
		MiddleName:  middleName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}
	return uc.contactRepo.CreateContact(contact)
}

func (uc *ContactUsecase) GetContactByID(id int) (*domain.Contact, error) {
	return uc.contactRepo.GetContactByID(id)
}

func (uc *ContactUsecase) UpdateContact(id int, firstName, middleName, lastName string, phoneNumber int) error {
	contact, err := uc.contactRepo.GetContactByID(id)
	if err != nil {
		return err
	}

	contact.FirstName = firstName
	contact.MiddleName = middleName
	contact.LastName = lastName
	contact.PhoneNumber = phoneNumber

	return uc.contactRepo.UpdateContact(contact)
}

func (uc *ContactUsecase) DeleteContact(id int) error {
	return uc.contactRepo.DeleteContact(id)
}

func (uc *ContactUsecase) CreateGroup(name string) error {
	group := &domain.Group{
		Name: name,
	}
	return uc.groupRepo.CreateGroup(group)
}

func (uc *ContactUsecase) GetGroupByID(id int) (*domain.Group, error) {
	return uc.groupRepo.GetGroupByID(id)
}

func (uc *ContactUsecase) AddContactToGroup(contactID, groupID int) error {
	return uc.groupRepo.AddContactToGroup(contactID, groupID)
}
