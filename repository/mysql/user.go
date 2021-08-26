package mysql

import (
	"api-gmr/repository"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo() repository.User {
	return &userRepo{
		db: getDB(),
	}
}

func (repo userRepo) selectFields() []string {
	fields := []string{
		"id",
		"username",
		"email",
		"password",
		"blok",
		"name",
	}
	return fields
}

func (repo userRepo) FindByUsername(ctx context.Context, username string) (repository.UserModel, error) {
	query := fmt.Sprintf("SELECT %s FROM users WHERE username = ?", strings.Join(repo.selectFields(), ","))
	statement, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "failed preparing query: %s", statement)
	}
	defer statement.Close()

	var u user
	err = statement.QueryRowContext(ctx, username).Scan(
		&u.id,
		&u.username,
		&u.email,
		&u.passwordHash,
		&u.blok,
		&u.name,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed scanning row")
	}
	return u, nil
}

func (repo userRepo) FindByUserID(ctx context.Context, userID int) (repository.UserModel, error) {
	query := fmt.Sprintf("SELECT %s FROM users WHERE id = ?", strings.Join(repo.selectFields(), ","))
	statement, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "failed preparing query: %s", statement)
	}
	defer statement.Close()

	var u user
	err = statement.QueryRowContext(ctx, userID).Scan(
		&u.id,
		&u.username,
		&u.email,
		&u.passwordHash,
		&u.blok,
		&u.name,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed scanning row")
	}
	return u, nil
}

func (repo userRepo) UpdateEmailandPassword(ctx context.Context, user repository.UserModel) error {
	return nil
}
