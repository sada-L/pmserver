package controller

import "github.com/sada-L/pmserver/internal/model"

type groupController struct {
  GroupService *model.GroupService 
}

func NewGroupController(s *model.GroupService) model.GroupController {
  return &groupController{GroupService: s} 
}

func (c groupController) Create()  {
  
}
