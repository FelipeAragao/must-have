package update_deal

import (
	"context"
	"errors"
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"time"
)

type DealInputDTO struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
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
	Photos []string `json:"photos"`
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

type UpdateDealUseCaseInterface interface {
	Execute(context context.Context, input *DealInputDTO) (*DealOutputDTO, error)
}

type UpdateDealDealCase struct {
	DealGateway gateway.DealGateway
}

func NewUpdateDealUseCase(dealGateway gateway.DealGateway) *UpdateDealDealCase {
	return &UpdateDealDealCase{
		DealGateway: dealGateway,
	}
}

func (uc *UpdateDealDealCase) Execute(ctx context.Context, input *DealInputDTO) (*DealOutputDTO, error) {

	deal, err := uc.DealGateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if input.UserID != deal.UserID.String() {
		return nil, errors.New("user not authorized to update this deal")
	}

	deal.Type = entity.DealType(input.Type)
	deal.Value = input.Value
	deal.Description = input.Description
	deal.TradeFor = input.TradeFor
	deal.Location = entity.LocationDeal{
		Lat:     input.Location.Lat,
		Lng:     input.Location.Lng,
		Address: input.Location.Address,
		City:    input.Location.City,
		State:   input.Location.State,
		ZipCode: input.Location.ZipCode,
	}
	deal.Urgency = entity.Urgency{
		Type:      entity.UrgencyType(input.Urgency.Type),
		LimitDate: input.Urgency.LimitDate,
	}
	deal.Photos = input.Photos

	err = deal.Modify()
	if err != nil {
		return nil, err
	}

	err = uc.DealGateway.Update(ctx, deal)
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
