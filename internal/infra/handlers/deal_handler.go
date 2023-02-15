package handlers

import (
	"encoding/json"
	"github.com/FelipeAragao/must-have/internal/usecase/deal/create_deal"
	"github.com/FelipeAragao/must-have/internal/usecase/deal/find_by_id"
	"github.com/FelipeAragao/must-have/internal/usecase/deal/update_deal"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type DealHandler struct {
	ucCreateDeal create_deal.CreateDealUseCaseInterface
	ucUpdateDeal update_deal.UpdateDealUseCaseInterface
	FindById     find_by_id.FindByIDUseCaseInterface
}

func NewDealHandler(ucCreateDeal create_deal.CreateDealUseCaseInterface, ucUpdateDeal update_deal.UpdateDealUseCaseInterface, findById find_by_id.FindByIDUseCaseInterface) *DealHandler {
	return &DealHandler{ucCreateDeal: ucCreateDeal, ucUpdateDeal: ucUpdateDeal, FindById: findById}
}

// CreateDeal godoc
// @Summary      Create deal
// @Description  Create deal
// @Tags         deals
// @Accept       json
// @Produce      json
// @Param        request     body      create_deal.DealInputDTO  true  "deal request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /deals [post]
func (h *DealHandler) CreateDeal(w http.ResponseWriter, r *http.Request) {
	var dto create_deal.DealInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.ucCreateDeal.Execute(r.Context(), &dto)
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

// UpdateDeal godoc
// @Summary      Update a deal
// @Description  Update a deal
// @Tags         deals
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "deal ID" Format(uuid)
// @Param        request     body      update_deal.DealInputDTO  true  "deal request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /deals/{id} [put]
// @Security ApiKeyAuth
func (h *DealHandler) UpdateDeal(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto update_deal.DealInputDTO
	dto.ID = id

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.ucUpdateDeal.Execute(r.Context(), &dto)
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

// GetDealById godoc
// @Summary      Get a deal
// @Description  Get a deal
// @Tags         deals
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "deal ID" Format(uuid)
// @Success      200  {object}  entity.Deal
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /deals/{id} [get]
// @Security ApiKeyAuth
func (h *DealHandler) GetDealById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto find_by_id.DealInputDTO
	dto.ID = id

	output, err := h.FindById.Execute(r.Context(), &dto)
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
