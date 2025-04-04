package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("mock create error: %v", err)
		}
		defer db.Close()

		inUser := &model.User{
			Name:     "Test Name",
			Email:    "test@email",
			Password: "test_password",
		}

		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

		mock.ExpectQuery(`
			INSERT INTO users .*
			VALUES \(.*,.*,.*\)
			RETURNING id`).
			WithArgs(inUser.Name, inUser.Email, inUser.Password).
			WillReturnRows(rows)

		ctx := context.Background()

		repo := repository.NewUserRepository(db)
		err = repo.Create(ctx, inUser)

		assert.NoError(t, err)
	})
}

func TestUserRepository_GetByEmail(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("mock create error: %v", err)
		}
		defer db.Close()

		exUser := &model.User{
			Id:       1,
			Name:     "Test Name",
			Email:    "test@email",
			Password: "test_password",
		}

		rows := sqlmock.
			NewRows([]string{"id", "name", "email", "password"}).
			AddRow(exUser.Id, exUser.Name, exUser.Email, exUser.Password)

		mock.ExpectQuery(`
			SELECT id, name, email 
			FROM users 
			WHERE email = $1`).
			WithArgs("test@email").
			WillReturnRows(rows)

		ctx := context.Background()

		repo := repository.NewUserRepository(db)
		user, err := repo.GetByEmail(ctx, "test@email")

		assert.NoError(t, err)
		assert.Equal(t, exUser, user)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("user not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("mock create error: %v", err)
		}
		defer db.Close()

		mock.ExpectQuery(`select id, email, name from users where email = $1`).
			WithArgs("email").
			WillReturnError(sql.ErrNoRows)

		ctx := context.Background()

		repo := repository.NewUserRepository(db)
		user, _ := repo.GetByEmail(ctx, "email")

		assert.Nil(t, user)
	})
}
