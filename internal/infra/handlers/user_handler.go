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

type UserHandler struct {
	UserGateway gateway.UserGateway
}

func NewUserHandler(userGateway gateway.UserGateway) *UserHandler {
	return &UserHandler{UserGateway: userGateway}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto *create_user.UserInputDTO
	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := create_user.NewCreateUserUseCase(h.UserGateway)
	output, err := usecase.Execute(dto)
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
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto *update_user.UserInputDTO
	dto.ID = id

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := update_user.NewUpdateUserUseCase(h.UserGateway)
	output, err := usecase.Execute(dto)
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

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto *find_by_id_user.UserInputDTO
	dto.ID = id

	usecase := find_by_id_user.NewFindByIdUserUseCase(h.UserGateway)
	output, err := usecase.Execute(dto)

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

func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto *change_password.UserInputDTO
	dto.ID = id

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := change_password.NewChangePasswordUserUseCase(h.UserGateway)
	output, err := usecase.Execute(dto)
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
