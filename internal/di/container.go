package container

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/service"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	db *pgx.Conn

	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler

	authService *service.AuthService
	authHandler *handlers.AuthHandler

	forgotRepo    *repository.ForgotPasswordRepository
	forgotService *service.ForgotPasswordService
	forgotHandler *handlers.ForgotPasswordHandler

	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handlers.ProductHandler
}

func NewContainer(db *pgx.Conn) *Container {

	c := Container{
		db: db,
	}

	c.init()

	return &c
}

func (c *Container) init() {

	c.userRepo = repository.NewUserRepository(c.db)

	c.userService = service.NewUserService(c.userRepo)

	c.userHandler = handlers.NewUserHandler(c.userService)
	

	c.authService = service.NewAuthService(c.userRepo)

	c.authHandler = handlers.NewAuthHandler(c.authService)


	c.forgotRepo = repository.NewForgotPasswordRepository(c.db)

	c.forgotService = service.NewForgotPasswordService(
		c.userRepo,
		c.forgotRepo,
	)

	c.forgotHandler = handlers.NewForgotPasswordHandler(
		c.forgotService,
	)

	c.productRepo = repository.NewProductRepository(c.db)

	c.productService = service.NewProductService(c.productRepo)

	c.productHandler = handlers.NewProductHandler(c.productService)
}

func (c *Container) AuthHandler() *handlers.AuthHandler {
	return c.authHandler
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}

func (c *Container) ProductHandler() *handlers.ProductHandler {
	return c.productHandler
}

func (c *Container) ForgotPasswordHandler() *handlers.ForgotPasswordHandler {
	return c.forgotHandler
}

func Connect() (*pgx.Conn, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/postgres?sslmode=%s",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGSSLMODE"),
	)

	conn, err := pgx.Connect(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	return conn, nil
}