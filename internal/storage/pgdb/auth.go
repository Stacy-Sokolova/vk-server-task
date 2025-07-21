package pgdb

import (
	"context"
	"fmt"
	"vk-server-task/internal/models"
	"vk-server-task/pkg/postgres"
)

const usersTable = "users"

type authStorage struct {
	pg *postgres.Postgres
}

func NewAuthStorage(pg *postgres.Postgres) *authStorage {
	return &authStorage{
		pg: pg,
	}
}

func (s *authStorage) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1", usersTable)
	row := s.pg.Pool.QueryRow(ctx, query, login)
	err := row.Scan(&user)

	return &user, err
}

func (s *authStorage) CreateUser(ctx context.Context, login, password string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password) values ($1, $2) RETURNING id", usersTable)

	row := s.pg.Pool.QueryRow(ctx, query, login, password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
