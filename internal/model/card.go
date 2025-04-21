package model

import "context"

type Card struct {
	Id       string
	Name     string
	UserName string
	Url      string
	Password string
	UserId   string
	GroupId  string
}

type CardRepository interface {
	ByUserId(context.Context, string) (*[]Card, error)
}

type CardService interface {
	CardsByUserId(context.Context, string) (*[]Card, error)
}
