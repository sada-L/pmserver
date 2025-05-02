package service

import (
	"context"
	"fmt"
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/internal/repository"
	"github.com/sada-L/pmserver/pkg/postgres"
)

type groupService struct {
	db *postgres.DB
}

func NewGroupService(db *postgres.DB) model.GroupService {
	return &groupService{db: db}
}

func (gs *groupService) CreateGroup(ctx context.Context, group *model.Group) (uint, error) {
	return 0, nil
}

func (gs *groupService) UpdateGroup(ctx context.Context, group *model.Group) error {
	return nil
}

func (gs *groupService) DeleteGroup(ctx context.Context, id uint) error {
	return nil
}

func (gs *groupService) GroupsByUser(ctx context.Context, user *model.User) (*[]model.Group, error) {
	gr := repository.NewGroupRepository(gs.db)
	groups, err := gr.ByUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("groupService - GroupsByUser - gr.ByUser: %w", err)
	}

	return groups, nil
}
