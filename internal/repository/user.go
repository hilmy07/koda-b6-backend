package repository

import (
	"context"

	"backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {

	row := r.db.QueryRow(
		context.Background(),
		"SELECT id,email,password FROM users WHERE email=$1",
		email,
	)

	user := models.User{}

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(email string, password string) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO users (email, password) VALUES ($1,$2)`,
		email,
		password,
	)

	return err
}