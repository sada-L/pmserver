package controller

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
	"net/http"
	"strconv"
)

type GroupController struct {
	gs model.GroupService
}

func NewGroupController(gs model.GroupService) *GroupController {
	return &GroupController{gs: gs}
}

func (gc *GroupController) CreateGroup() http.HandlerFunc {
	type Input struct {
		Title   string `json:"title"`
		Image   string `json:"image"`
		GroupId uint   `json:"group_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := utils.UserFromContext(ctx)

		input := Input{}
		if err := utils.ReadJSON(r.Body, &input); err != nil {
			utils.BadRequestError(w)
			return
		}

		group := &model.Group{
			Title:   input.Title,
			Image:   input.Image,
			GroupId: input.GroupId,
			UserId:  user.Id,
		}

		id, err := gc.gs.CreateGroup(r.Context(), group)
		if err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, id)
	}
}

func (gc *GroupController) UpdateGroup() http.HandlerFunc {
	type Input struct {
		Title   string `json:"title"`
		Image   string `json:"image"`
		GroupId uint   `json:"group_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.BadRequestError(w)
			return
		}

		input := Input{}
		if err = utils.ReadJSON(r.Body, &input); err != nil {
			utils.BadRequestError(w)
			return
		}

		group := &model.Group{
			Id:      uint(id),
			Title:   input.Title,
			Image:   input.Image,
			GroupId: input.GroupId,
		}

		if err = gc.gs.UpdateGroup(r.Context(), group); err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, nil)
	}
}

func (gc *GroupController) DeleteGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.BadRequestError(w)
			return
		}

		if err = gc.gs.DeleteGroup(r.Context(), uint(id)); err != nil {
			utils.InternalError(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, nil)
	}
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
