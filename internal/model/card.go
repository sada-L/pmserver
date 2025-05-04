package model

import "context"

type Card struct {
	Id         uint   `json:"id"`
	UserId     uint   `json:"-"`
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
	Create(context.Context, *Card) (uint, error)
	Update(context.Context, *Card) error
	Delete(context.Context, uint) error
}

type CardService interface {
	CardsByUser(context.Context, *User) (*[]Card, error)
	CreateCard(context.Context, *Card) (uint, error)
	UpdateCard(context.Context, *Card) error
	DeleteCard(context.Context, uint) error
}
