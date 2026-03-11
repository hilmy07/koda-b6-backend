package repository

import "github.com/jackc/pgx/v5"

type ForgotPasswordRepository struct {
	db *pgx.Conn
}

func NewForgotPasswordRepository(db *pgx.Conn) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{db: db}
}

func (r *ForgotPasswordRepository) CreateForgotPassword() {

}

func (r *ForgotPasswordRepository) GetDataByEmailnCode() {

}

func (r *ForgotPasswordRepository) DeleteDataByCode() {
	
}
