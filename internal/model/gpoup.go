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
	ByUser(context.Context, *User) (*[]Group, error)
}

type GroupService interface {
	GroupsByUser(context.Context, *User) (*[]Group, error)
}
