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

	rows, err := r.db.Query(
		context.Background(),
		`SELECT id,email,fullname,password,phone,address,profile_img,created_at,updated_at 
		FROM users WHERE email=$1`,
		email,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.User])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(req models.CreateUserRequest) error {

	now := time.Now()

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO users 
		(fullname, email, password, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5)`,
		req.Fullname,
		req.Email,
		req.Password,
		now,
		now,
	)

	return err
}

func (r *UserRepository) UpdatePasswordByEmail(email string, password string) error {

	_, err := r.db.Exec(
		context.Background(),
		`UPDATE users
		 SET password = $1,
		     updated_at = NOW()
		 WHERE email = $2`,
		password,
		email,
	)

	return err
}

func (r *UserRepository) UpdateUserProfile(req models.CreateUserRequest) error {

	_, err := r.db.Exec(
		context.Background(),
		`UPDATE users 
		SET fullname=$1,
		    phone=$2,
		    address=$3,
		    profile_img=$4,
		    updated_at=$5
		WHERE email=$6`,
		req.Fullname,
		req.Phone,
		req.Address,
		req.Profile_img,
		time.Now(),
		req.Email,
	)

	return err
}