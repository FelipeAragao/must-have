package database

import (
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	entityPkg "github.com/FelipeAragao/must-have/pkg/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) Update(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) FindByLogin(login string) (*entity.User, error) {
	id := entityPkg.NewID()
	hash, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	return &entity.User{
		ID:        id,
		Name:      "John Doe",
		Email:     "john.doe@vibbra.com.br",
		Login:     "john.doe",
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Location: entity.Location{
			Lat:     -23.5506507,
			Lng:     -46.6333824,
			Address: "Rua Vergueiro, 3185",
			City:    "São Paulo",
			State:   "SP",
			ZipCode: 65000000,
		},
	}, nil
}
