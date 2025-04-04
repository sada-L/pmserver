package model

type Card struct {
	Id       uint
	Name     string
	UserName string
	Url      string
	Password string
	UserId   uint
	GroupId  uint
}

type CardRepository interface {
	Create(card *Card) error
	Update(card *Card) error
	Delete(id uint) error
}

type CardService interface {
	Create()
}

type CardController interface {
	Create()
}
