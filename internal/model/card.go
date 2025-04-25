package model

import "context"

type Card struct {
	Id         uint   `json:"id"`
	UserId     string `json:"-"`
	GroupId    uint   `json:"group_id"`
	Title      string `json:"title"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Website    string `json:"website"`
	Notes      string `json:"notes"`
	Image      string `json:"image"`
	IsFavorite bool   `json:"is_favorite"`
}

type CardRepository interface {
	ByUser(context.Context, *User) (*[]Card, error)
}

type CardService interface {
	CardsByUser(context.Context, *User) (*[]Card, error)
}
