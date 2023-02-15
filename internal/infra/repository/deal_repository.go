package repository

import (
	"context"
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	"github.com/FelipeAragao/must-have/internal/infra/repository/db"
	"strings"
)

type DealRepository struct {
	db      *sql.DB
	queries *db.Queries
}

func NewDealRepository(dbSQL *sql.DB) *DealRepository {
	return &DealRepository{
		dbSQL,
		db.New(dbSQL),
	}
}

func (r DealRepository) Create(ctx context.Context, deal *entity.Deal) error {
	photos := strings.Join(deal.Photos, ",")
	var err = r.queries.CreateDeal(ctx, db.CreateDealParams{
		ID:              deal.ID.String(),
		UserID:          deal.UserID.String(),
		Value:           deal.Value,
		Description:     sql.NullString{String: deal.Description, Valid: true},
		TradeFor:        sql.NullString{String: deal.TradeFor, Valid: true},
		LocationLat:     sql.NullFloat64{Float64: deal.Location.Lat, Valid: true},
		LocationLng:     sql.NullFloat64{Float64: deal.Location.Lng, Valid: true},
		LocationAddress: deal.Location.Address,
		LocationCity:    deal.Location.City,
		LocationState:   deal.Location.State,
		LocationZipCode: int64(deal.Location.ZipCode),
		Photos:          sql.NullString{String: photos, Valid: true},
		CreatedAt:       deal.CreatedAt,
		UpdatedAt:       deal.UpdatedAt,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r DealRepository) Update(ctx context.Context, deal *entity.Deal) error {
	//TODO implement me
	panic("implement me")
}

func (r DealRepository) FindByID(ctx context.Context, id string) (*entity.Deal, error) {
	//TODO implement me
	panic("implement me")
}

func (r DealRepository) FindByUser(ctx context.Context, deal_id string) (*entity.Deal, error) {
	//TODO implement me
	panic("implement me")
}
