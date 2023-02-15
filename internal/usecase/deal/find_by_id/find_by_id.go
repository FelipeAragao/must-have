package find_by_id

import (
	"context"
	"errors"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"time"
)

type DealInputDTO struct {
	ID string `json:"id"`
}

type DealOutputDTO struct {
	ID          string  `json:"id"`
	Type        int     `json:"type"`
	Value       float64 `json:"value"`
	Description string  `json:"description"`
	TradeFor    string  `json:"trade_for"`
	Location    struct {
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
		Address string  `json:"address"`
		City    string  `json:"city"`
		State   string  `json:"state"`
		ZipCode int     `json:"zip_code"`
	} `json:"location"`
	Urgency struct {
		Type      int       `json:"type"`
		LimitDate time.Time `json:"limit_date"`
	} `json:"urgency"`
	Photos    []string  `json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindByIDUseCaseInterface interface {
	Execute(context context.Context, input *DealInputDTO) (*DealOutputDTO, error)
}

type FindByIDUseCase struct {
	gateway.DealGateway
}

func NewFindByIdUseCase(dealGateway gateway.DealGateway) *FindByIDUseCase {
	return &FindByIDUseCase{DealGateway: dealGateway}
}

func (uc *FindByIDUseCase) Execute(ctx context.Context, input *DealInputDTO) (*DealOutputDTO, error) {

	if input == nil || input.ID == "" {
		return nil, errors.New("id is required")
	}

	deal, err := uc.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &DealOutputDTO{
		ID:          deal.ID.String(),
		Type:        int(deal.Type),
		Value:       deal.Value,
		Description: deal.Description,
		TradeFor:    deal.TradeFor,
		Location: struct {
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
			Address string  `json:"address"`
			City    string  `json:"city"`
			State   string  `json:"state"`
			ZipCode int     `json:"zip_code"`
		}{
			Lat:     deal.Location.Lat,
			Lng:     deal.Location.Lng,
			Address: deal.Location.Address,
			City:    deal.Location.City,
			State:   deal.Location.State,
			ZipCode: deal.Location.ZipCode,
		},
		Urgency: struct {
			Type      int       `json:"type"`
			LimitDate time.Time `json:"limit_date"`
		}{
			Type:      int(deal.Urgency.Type),
			LimitDate: deal.Urgency.LimitDate,
		},
		Photos:    deal.Photos,
		CreatedAt: deal.CreatedAt,
		UpdatedAt: deal.UpdatedAt,
	}, nil
}
