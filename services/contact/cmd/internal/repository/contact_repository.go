package repository

import (
	"architecture_go/services/contact/cmd/internal/domain"
	"context"

	"github.com/jackc/pgx/v4"
)

type contactRepository struct {
	conn *pgx.Conn
}

func NewContactRepository(conn *pgx.Conn) ContactRepository {
	return &contactRepository{
		conn: conn,
	}
}

func (r *contactRepository) CreateContact(contact *domain.Contact) error {
	_, err := r.conn.Exec(context.Background(), "INSERT INTO contacts (id, first_name, middle_name, last_name, phone_number) VALUES ($1, $2, $3, $4, $5)", contact.ID, contact.FirstName, contact.MiddleName, contact.LastName, contact.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *contactRepository) GetContactByID(id int) (*domain.Contact, error) {
	var contact domain.Contact
	err := r.conn.QueryRow(context.Background(), "SELECT id, first_name, middle_name, last_name, phone_number FROM contacts WHERE id = $1", id).Scan(&contact.ID, &contact.FirstName, &contact.MiddleName, &contact.LastName, &contact.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactRepository) UpdateContact(contact *domain.Contact) error {
	_, err := r.conn.Exec(context.Background(), "UPDATE contacts SET first_name = $2, middle_name = $3, last_name = $4, phone_number = $5 WHERE id = $1", contact.ID, contact.FirstName, contact.MiddleName, contact.LastName, contact.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *contactRepository) DeleteContact(id int) error {
	_, err := r.conn.Exec(context.Background(), "DELETE FROM contacts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
