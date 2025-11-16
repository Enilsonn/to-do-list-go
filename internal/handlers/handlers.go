package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"to-do-list/internal/models"
	"to-do-list/internal/repository"
	"to-do-list/internal/utils"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Repo repository.TaskRepository
}

func NewHandler(repo repository.TaskRepository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var t models.Tarefa
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid body request")
		return
	}

	ctx := r.Context()

	id, err := h.Repo.Create(ctx, &t)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError,
			fmt.Sprintf("error to create task: %v", err),
		)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, map[string]any{
		"id": id,
	})
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	var list []*models.Tarefa

	ctx := r.Context()

	list, err := h.Repo.List(ctx)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError,
			fmt.Sprintf("error to list tasks: %v", err),
		)
		return
	}

	utils.RespondJSON(w, http.StatusOK, list)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParamFromCtx(ctx, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest,
			fmt.Sprintf("invalid id: %v", err),
		)
		return
	}

	var t *models.Tarefa
	t, err = h.Repo.GetByID(ctx, id)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest,
			fmt.Sprintf("error to get task by id: %v", err),
		)
		return
	}

	utils.RespondJSON(w, http.StatusOK, t)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParamFromCtx(ctx, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest,
			fmt.Sprintf("invalid id: %v", err),
		)
		return
	}

	var t *models.Tarefa
	if err = json.NewDecoder(r.Body).Decode(&t); err != nil {
		utils.RespondError(w, http.StatusBadRequest,
			fmt.Sprintf("invalid body request: %v", err),
		)
		return
	}

	if err = h.Repo.Update(ctx, id, t); err != nil {
		utils.RespondError(w, http.StatusInternalServerError,
			fmt.Sprintf("error to update: %v", err),
		)
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"id": t.ID,
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParamFromCtx(ctx, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest,
			fmt.Sprintf("invalid id: %v", err),
		)
		return
	}

	if err = h.Repo.Delete(ctx, id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError,
			fmt.Sprintf("error to delete: %v", err),
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
