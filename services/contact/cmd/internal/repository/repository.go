package repository

import (
	"architecture_go/services/contact/cmd/internal/domain"
	"context"
)

type ContactRepository interface {
	CreateContact(ctx context.Context, contact *domain.Contact) error
	GetContactByID(ctx context.Context, id int) (*domain.Contact, error)
	UpdateContact(ctx context.Context, contact *domain.Contact) error
	DeleteContact(ctx context.Context, id int) error
}

type GroupRepository interface {
	CreateGroup(ctx context.Context, group *domain.Group) error
	GetGroupByID(ctx context.Context, id int) (*domain.Group, error)
	AddContactToGroup(ctx context.Context, contactID, groupID int) error
}
