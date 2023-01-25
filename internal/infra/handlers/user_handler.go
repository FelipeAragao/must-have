package handlers

import (
	"encoding/json"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/internal/usecase/user/change_password"
	"github.com/FelipeAragao/must-have/internal/usecase/user/create_user"
	"github.com/FelipeAragao/must-have/internal/usecase/user/find_by_id_user"
	"github.com/FelipeAragao/must-have/internal/usecase/user/update_user"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserGateway gateway.UserGateway
}

func NewUserHandler(userGateway gateway.UserGateway) *UserHandler {
	return &UserHandler{UserGateway: userGateway}
}

// CreateUser godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      create_user.UserInputDTO  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto create_user.UserInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := create_user.NewCreateUserUseCase(h.UserGateway)
	output, err := usecase.Execute(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Update a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "user ID" Format(uuid)
// @Param        request     body      update_user.UserInputDTO  true  "user request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /users/{id} [put]
// @Security ApiKeyAuth
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto update_user.UserInputDTO
	dto.ID = id

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := update_user.NewUpdateUserUseCase(h.UserGateway)
	output, err := usecase.Execute(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetUserById godoc
// @Summary      Get a user
// @Description  Get a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "user ID" Format(uuid)
// @Success      200  {object}  entity.User
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /users/{id} [get]
// @Security ApiKeyAuth
func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto find_by_id_user.UserInputDTO
	dto.ID = id

	usecase := find_by_id_user.NewFindByIdUserUseCase(h.UserGateway)
	output, err := usecase.Execute(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// ChangePassword godoc
// @Summary      Change password a user
// @Description  Change password a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "user ID" Format(uuid)
// @Param        request     body      change_password.UserInputDTO  true  "user request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /users/{id}/change-password [put]
// @Security ApiKeyAuth
func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto change_password.UserInputDTO
	dto.ID = id

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := change_password.NewChangePasswordUserUseCase(h.UserGateway)
	output, err := usecase.Execute(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
