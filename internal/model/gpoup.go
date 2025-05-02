package model

import (
	"context"
)

type Group struct {
	Id      uint   `json:"id"`
	UserId  uint   `json:"-"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	GroupId uint   `json:"group_id"`
}

type GroupRepository interface {
	Create(context.Context, *Group) (uint, error)
	Update(context.Context, *Group) error
	Delete(context.Context, uint) error
	ByUser(context.Context, *User) (*[]Group, error)
}

type GroupService interface {
	CreateGroup(context.Context, *Group) (uint, error)
	UpdateGroup(context.Context, *Group) error
	DeleteGroup(context.Context, uint) error
	GroupsByUser(context.Context, *User) (*[]Group, error)
}
