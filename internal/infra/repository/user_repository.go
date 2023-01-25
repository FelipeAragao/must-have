package repository

import (
	"context"
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	"github.com/FelipeAragao/must-have/internal/infra/repository/db"
	entityPkg "github.com/FelipeAragao/must-have/pkg/entity"
)

type UserRepository struct {
	db      *sql.DB
	cxt     context.Context
	queries *db.Queries
}

func NewUserRepository(dbSQL *sql.DB) *UserRepository {
	return &UserRepository{
		dbSQL,
		context.Background(),
		db.New(dbSQL),
	}
}

func (r *UserRepository) Create(user *entity.User) error {
	err := r.queries.CreateUser(r.cxt, db.CreateUserParams{
		ID:              user.ID.String(),
		Name:            user.Name,
		Email:           user.Email,
		Login:           user.Login,
		Password:        user.Password,
		LocationLat:     sql.NullFloat64{Float64: user.Location.Lat, Valid: true},
		LocationLng:     sql.NullFloat64{Float64: user.Location.Lng, Valid: true},
		LocationAddress: user.Location.Address,
		LocationCity:    user.Location.City,
		LocationState:   user.Location.State,
		LocationZipCode: int64(user.Location.ZipCode),
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(user *entity.User) error {
	err := r.queries.UpdateUser(r.cxt, db.UpdateUserParams{
		ID:              user.ID.String(),
		Name:            user.Name,
		Email:           user.Email,
		Login:           user.Login,
		Password:        user.Password,
		LocationLat:     sql.NullFloat64{Float64: user.Location.Lat, Valid: true},
		LocationLng:     sql.NullFloat64{Float64: user.Location.Lng, Valid: true},
		LocationAddress: user.Location.Address,
		LocationCity:    user.Location.City,
		LocationState:   user.Location.State,
		LocationZipCode: int64(user.Location.ZipCode),
		UpdatedAt:       user.UpdatedAt,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	user, err := r.queries.FindByID(r.cxt, id)
	if err != nil {
		return nil, err
	}

	ID, err := entityPkg.ParseID(user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:        ID,
		Name:      user.Name,
		Email:     user.Email,
		Login:     user.Login,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Location: entity.Location{
			Lat:     user.LocationLat.Float64,
			Lng:     user.LocationLng.Float64,
			Address: user.LocationAddress,
			City:    user.LocationCity,
			State:   user.LocationState,
			ZipCode: int(user.LocationZipCode),
		},
	}, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	user, err := r.queries.FindByEmail(r.cxt, email)
	if err != nil {
		return nil, err
	}

	ID, err := entityPkg.ParseID(user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:        ID,
		Name:      user.Name,
		Email:     user.Email,
		Login:     user.Login,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Location: entity.Location{
			Lat:     user.LocationLat.Float64,
			Lng:     user.LocationLng.Float64,
			Address: user.LocationAddress,
			City:    user.LocationCity,
			State:   user.LocationState,
			ZipCode: int(user.LocationZipCode),
		},
	}, nil
}

func (r *UserRepository) FindByLogin(login string) (*entity.User, error) {
	user, err := r.queries.FindByEmail(r.cxt, login)
	if err != nil {
		return nil, err
	}

	ID, err := entityPkg.ParseID(user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:        ID,
		Name:      user.Name,
		Email:     user.Email,
		Login:     user.Login,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Location: entity.Location{
			Lat:     user.LocationLat.Float64,
			Lng:     user.LocationLng.Float64,
			Address: user.LocationAddress,
			City:    user.LocationCity,
			State:   user.LocationState,
			ZipCode: int(user.LocationZipCode),
		},
	}, nil
}
