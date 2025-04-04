package repository

import (
	"database/sql"

	"github.com/sada-L/pmserver/internal/model"
)

type groupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) model.GroupReporitory {
	return &groupRepository{db: db}
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
