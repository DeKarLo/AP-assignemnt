// usecase.go
package usecase

import (
	"architecture_go/services/contact/cmd/internal/domain"
	"architecture_go/services/contact/cmd/internal/repository"
	"context"
)

type Usecase interface {
	CreateContact(ctx context.Context, firstName, middleName, lastName string, phoneNumber int) error
	GetContactByID(ctx context.Context, id int) (*domain.Contact, error)
	UpdateContact(ctx context.Context, id int, firstName, middleName, lastName string, phoneNumber int) error
	DeleteContact(ctx context.Context, id int) error

	CreateGroup(ctx context.Context, name string) error
	GetGroupByID(ctx context.Context, id int) (*domain.Group, error)

	AddContactToGroup(ctx context.Context, contactID, groupID int) error
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

func (uc *ContactUsecase) CreateContact(ctx context.Context, firstName, middleName, lastName string, phoneNumber int) error {
	contact := &domain.Contact{
		FirstName:   firstName,
		MiddleName:  middleName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}
	return uc.contactRepo.CreateContact(ctx, contact)
}

func (uc *ContactUsecase) GetContactByID(ctx context.Context, id int) (*domain.Contact, error) {
	return uc.contactRepo.GetContactByID(ctx, id)
}

func (uc *ContactUsecase) UpdateContact(ctx context.Context, id int, firstName, middleName, lastName string, phoneNumber int) error {
	contact, err := uc.contactRepo.GetContactByID(ctx, id)
	if err != nil {
		return err
	}

	contact.FirstName = firstName
	contact.MiddleName = middleName
	contact.LastName = lastName
	contact.PhoneNumber = phoneNumber

	return uc.contactRepo.UpdateContact(ctx, contact)
}

func (uc *ContactUsecase) DeleteContact(ctx context.Context, id int) error {
	return uc.contactRepo.DeleteContact(ctx, id)
}

func (uc *ContactUsecase) CreateGroup(ctx context.Context, name string) error {
	group := &domain.Group{
		Name: name,
	}
	return uc.groupRepo.CreateGroup(ctx, group)
}

func (uc *ContactUsecase) GetGroupByID(ctx context.Context, id int) (*domain.Group, error) {
	return uc.groupRepo.GetGroupByID(ctx, id)
}

func (uc *ContactUsecase) AddContactToGroup(ctx context.Context, contactID, groupID int) error {
	return uc.groupRepo.AddContactToGroup(ctx, contactID, groupID)
}
