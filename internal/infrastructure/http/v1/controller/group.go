package controller

import (
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
	"net/http"
)

type GroupController struct {
	gs model.GroupService
}

func NewGroupController(gs model.GroupService) *GroupController {
	return &GroupController{gs: gs}
}

func (gc *GroupController) GetGroupsByUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := utils.UserFromContext(ctx)
		if user == nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		groups, err := gc.gs.GroupsByUser(r.Context(), user)
		if err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, groups)
	}
}
