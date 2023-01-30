package create_deal

import (
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"time"
)

type DealInputDTO struct {
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

type CreateDealUseCaseInterface interface {
	Execute(input *DealInputDTO) (*DealOutputDTO, error)
}

type CreateDealDealCase struct {
	DealGateway gateway.DealGateway
	UserGateway gateway.UserGateway
}

func NewCreateDealUseCase(dealGateway gateway.DealGateway, userGateway gateway.UserGateway) *CreateDealDealCase {
	return &CreateDealDealCase{
		DealGateway: dealGateway,
		UserGateway: userGateway,
	}
}

func (uc *CreateDealDealCase) Execute(input *DealInputDTO) (*DealOutputDTO, error) {

	user, err := uc.UserGateway.FindByID(input.UserID)
	if err != nil {
		return nil, err
	}

	deal, err := entity.NewDeal(
		user.ID,
		entity.DealType(input.Type),
		input.Value,
		input.Description,
		input.TradeFor,
		entity.LocationDeal{
			Lat:     input.Location.Lat,
			Lng:     input.Location.Lng,
			Address: input.Location.Address,
			City:    input.Location.City,
			State:   input.Location.State,
			ZipCode: input.Location.ZipCode,
		},
		entity.Urgency{
			Type:      entity.UrgencyType(input.Urgency.Type),
			LimitDate: input.Urgency.LimitDate,
		},
		input.Photos,
	)

	if err != nil {
		return nil, err
	}

	err = uc.DealGateway.Create(deal)
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
