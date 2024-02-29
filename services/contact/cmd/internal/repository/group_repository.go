package repository

import (
	"architecture_go/services/contact/cmd/internal/domain"
	"context"

	"github.com/jackc/pgx/v4"
)


type groupRepository struct {
	conn *pgx.Conn
}

func NewGroupRepository(conn *pgx.Conn) GroupRepository {
	return &groupRepository{
		conn: conn,
	}
}

func (r *groupRepository) CreateGroup(ctx context.Context, group *domain.Group) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO groups (id, name) VALUES ($1, $2)", group.ID, group.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *groupRepository) GetGroupByID(ctx context.Context, id int) (*domain.Group, error) {
	var group domain.Group
	err := r.conn.QueryRow(ctx, "SELECT id, name FROM groups WHERE id = $1", id).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupRepository) AddContactToGroup(ctx context.Context, contactID, groupID int) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO group_contacts (group_id, contact_id) VALUES ($1, $2)", groupID, contactID)
	if err != nil {
		return err
	}
	return nil
}
