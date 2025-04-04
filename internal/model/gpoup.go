package model

type Group struct {
  Id     uint
  Name   string
  Cards  []Card
  UserId uint
}

type GroupReporitory interface {
  Create(group *Group) error
  Update(group *Group) error
  Delete(id uint) error
}

type GroupService interface {
  Create() 
}

type GroupController interface {
  Create()
}
