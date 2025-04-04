package service

import "github.com/sada-L/pmserver/internal/model"

type groupService struct {
  GroupRepository *model.GroupReporitory
}
 
func NewGroupService(r *model.GroupReporitory) model.GroupService  {
  return &groupService{GroupRepository: r} 
}

func (s groupService) Create()  {

}
