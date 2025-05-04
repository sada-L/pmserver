package repository

import (
	"context"
	"github.com/sada-L/pmserver/pkg/postgres"

	"github.com/sada-L/pmserver/internal/model"
)

type groupRepository struct {
	q postgres.Querier
}

func NewGroupRepository(q postgres.Querier) model.GroupRepository {
	return &groupRepository{q: q}
}

func (gr *groupRepository) Create(ctx context.Context, group *model.Group) (uint, error) {
	query := `INSERT INTO groups (title, image, group_id, user_id) VALUES ($1, $2, $3, $4) RETURNING id`

	args := []interface{}{group.Title, group.Image, group.GroupId, group.UserId}
	if err := gr.q.QueryRowContext(ctx, query, args...).Scan(&group.Id); err != nil {
		return 0, err
	}

	return group.Id, nil
}

func (gr *groupRepository) Update(ctx context.Context, group *model.Group) error {
	query := `UPDATE groups SET title = $1, image = $2, group_id = $3 WHERE id = $4`

	args := []interface{}{group.Title, group.Image, group.GroupId, group.Id}
	if _, err := gr.q.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (gr *groupRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM groups WHERE id = $1`

	if err := gr.q.QueryRowContext(ctx, query, id).Err(); err != nil {
		return err
	}

	return nil
}

func (gr *groupRepository) ByUser(ctx context.Context, user *model.User) (*[]model.Group, error) {
	groupQuery := `SELECT id, title, image, group_id FROM groups WHERE user_id = $1`

	var groups []model.Group
	rows, err := gr.q.QueryContext(ctx, groupQuery, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group model.Group
		if err = rows.Scan(
			&group.Id,
			&group.Title,
			&group.Image,
			&group.GroupId); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &groups, nil
}
