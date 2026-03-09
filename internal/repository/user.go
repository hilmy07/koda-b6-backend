package repository

import (
	"context"
	"time"

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
		`SELECT id,email,fullname,password,phone,address,profile_img,created_at,updated_at 
		FROM users WHERE email=$1`,
		email,
	)

	user := models.User{}

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Fullname,
		&user.Password,
		&user.Phone,
		&user.Address,
		&user.Profile_img,
		&user.Created_at,
		&user.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(req models.CreateUserRequest) error {

	now := time.Now()

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO users 
		(email, fullname, password, phone, address, profile_img, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		req.Email,
		req.Fullname,
		req.Password,
		req.Phone,
		req.Address,
		req.Profile_img,
		now,
		now,
	)

	return err
}