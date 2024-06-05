package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ronymmoura/goliath-national-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		ID:             uuid.New(),
		HashedPassword: util.RandomString(64),
		FullName:       util.RandomString(10),
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.PasswordChangedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)

	user, err := testStore.GetUser(context.Background(), createdUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.HashedPassword, user.HashedPassword)
	require.Equal(t, createdUser.FullName, user.FullName)
	require.Equal(t, createdUser.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.PasswordChangedAt)
	require.WithinDuration(t, createdUser.CreatedAt, user.CreatedAt, time.Second)
}

func TestGetUserByEmail(t *testing.T) {
	createdUser := createRandomUser(t)

	user, err := testStore.GetUserByEmail(context.Background(), createdUser.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, createdUser.ID, user.ID)
	require.Equal(t, createdUser.HashedPassword, user.HashedPassword)
	require.Equal(t, createdUser.FullName, user.FullName)
	require.Equal(t, createdUser.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.PasswordChangedAt)
	require.WithinDuration(t, createdUser.CreatedAt, user.CreatedAt, time.Second)
}

func TestEditUser(t *testing.T) {
	createdUser := createRandomUser(t)

	arg := UpdateUserParams{
		ID:       createdUser.ID,
		FullName: createdUser.FullName,
		Email:    createdUser.Email,
	}

	user, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, createdUser.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.PasswordChangedAt)
	require.WithinDuration(t, createdUser.CreatedAt, user.CreatedAt, time.Second)
}

func TestEditUserPassword(t *testing.T) {
	createdUser := createRandomUser(t)

	arg := UpdateUserPasswordParams{
		ID:                createdUser.ID,
		HashedPassword:    util.RandomString(64),
		PasswordChangedAt: time.Now(),
	}

	user, err := testStore.UpdateUserPassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.PasswordChangedAt)
	require.WithinDuration(t, arg.PasswordChangedAt, user.PasswordChangedAt, time.Second)
}

func TestDeleteuser(t *testing.T) {
	createdUser := createRandomUser(t)

	err := testStore.DeleteUser(context.Background(), createdUser.ID)
	require.NoError(t, err)

	user, err := testStore.GetUser(context.Background(), createdUser.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, user)
}
