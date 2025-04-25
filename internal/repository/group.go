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

func (r groupRepository) Create(group *model.Group) error {
	return nil
}

func (r groupRepository) Update(group *model.Group) error {
	return nil
}

func (r groupRepository) Delete(id uint) error {
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
