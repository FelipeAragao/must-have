package entity

import (
	"github.com/FelipeAragao/must-have/pkg/entity"
	"github.com/FelipeAragao/must-have/pkg/validate"
	"time"
)

type DealType int
type UrgencyType int

const (
	VENDA DealType = iota + 1
	TROCA
	DESEJO
)

const (
	BAIXA UrgencyType = iota + 1
	MEDIA
	ALTA
	DATA
)

type Deal struct {
	ID          entity.ID    `json:"id"`
	UserID      entity.ID    `json:"user_id" validate:"required"`
	Type        DealType     `json:"type" validate:"required"`
	Value       float64      `json:"value" validate:"required,gte=0"`
	Description string       `json:"description" validate:"required,min=3,max=1000"`
	TradeFor    string       `json:"trade_for"`
	Location    LocationDeal `json:"location" validate:"required,dive"`
	Urgency     Urgency      `json:"urgency" validate:"required,dive"`
	Photos      []string     `json:"photos"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func NewDeal(userID entity.ID, Type DealType, value float64, description string, tradeFor string, location LocationDeal, urgency Urgency, photos []string) (*Deal, error) {
	d := &Deal{
		ID:          entity.NewID(),
		UserID:      userID,
		Type:        Type,
		Value:       value,
		Description: description,
		TradeFor:    tradeFor,
		Location:    location,
		Urgency:     urgency,
		Photos:      photos,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := d.validate()
	if err != nil {
		return nil, err
	}

	return d, nil

}

type LocationDeal struct {
	Lat     float64 `json:"lat" validate:"required,latitude"`
	Lng     float64 `json:"lng" validate:"required,longitude"`
	Address string  `json:"address" validate:"required,min=3,max=100"`
	City    string  `json:"city" validate:"required,min=3,max=100"`
	State   string  `json:"state" validate:"required,min=2,max=2"`
	ZipCode int     `json:"zip_code" validate:"required,gte=10000000,lte=99999999"`
}

type Urgency struct {
	Type      UrgencyType `json:"type" validate:"required"`
	LimitDate time.Time   `json:"limit_date"`
}

func (u *Deal) Modify() error {
	err := u.validate()
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now()
	return nil
}

func (d Deal) validate() error {
	err := validate.Struct(d)
	if err != nil {
		return err
	}
	return nil
}
