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
	tx, err := gs.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("groupService - CreateGroup - gs.db.BeginTx: %w", err)
	}
	defer tx.Rollback()

	gr := repository.NewGroupRepository(tx)
	id, err := gr.Create(ctx, group)
	if err != nil {
		return 0, fmt.Errorf("groupService - CreateGroup - gr.Create: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("groupService - CreateGroup - tx.Commit: %w", err)
	}

	return id, nil
}

func (gs *groupService) UpdateGroup(ctx context.Context, group *model.Group) error {
	tx, err := gs.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("groupService - CreateGroup - gs.db.BeginTx: %w", err)
	}
	defer tx.Rollback()

	gr := repository.NewGroupRepository(tx)
	if err = gr.Update(ctx, group); err != nil {
		return fmt.Errorf("groupService - CreateGroup - gr.Update: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("groupService - CreateGroup - tx.Commit: %w", err)
	}

	return nil
}

func (gs *groupService) DeleteGroup(ctx context.Context, id uint) error {
	tx, err := gs.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("groupService - DeleteGroup - gs.db.BeginTx: %w", err)
	}
	defer tx.Rollback()

	gr := repository.NewGroupRepository(tx)
	if err = gr.Delete(ctx, id); err != nil {
		return fmt.Errorf("groupService - DeleteGroup - gr.Delete: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("groupService - DeleteGroup - tx.Commit: %w", err)
	}

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
