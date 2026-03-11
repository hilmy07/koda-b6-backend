package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type ForgotPasswordRepository struct {
	db *pgx.Conn
}

func NewForgotPasswordRepository(db *pgx.Conn) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{db: db}
}

func (r *ForgotPasswordRepository) CreateForgotRequest(req models.ForgotPassword) error {
	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO forgot_password (email,code)
		 VALUES ($1,$2)`,
		req.Email,
		req.Code,
	)

	return err
}

func (r *ForgotPasswordRepository) GetDataByEmailnCode(email string, code string) (*models.ForgotPassword, error) {

	rows, err := r.db.Query(
		context.Background(),
		`SELECT id,email,code
		 FROM forgot_password
		 WHERE email=$1 AND code=$2`,
		email,
		code,
	)

	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ForgotPasswordRepository) DeleteDataByCode(code string) error {

	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM forgot_password WHERE code=$1`,
		code,
	)

	return err
}
