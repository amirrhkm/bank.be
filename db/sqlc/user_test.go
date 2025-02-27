package db

import (
	"context"
	"testing"
	"time"

	"github.com/amirrhkm/bank.be/util"
	"github.com/stretchr/testify/require"
)

func createTestUser(username string, hashedPassword string, fullName string, email string) (Users, error) {
	arg := CreateUserParams{
		Username:       username,
		HashedPassword: hashedPassword,
		FullName:       fullName,
		Email:          email,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	return user, err
}

func TestCreateUser(t *testing.T) {
	hashedPassword, err := util.HashPassword(util.RandomString(12))
	require.NoError(t, err)

	username := util.RandomOwner()
	fullName := util.RandomOwner()
	email := username + "@test.com"

	user, err := createTestUser(username, hashedPassword, fullName, email)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, username, user.Username)
	require.Equal(t, hashedPassword, user.HashedPassword)
	require.Equal(t, fullName, user.FullName)
	require.Equal(t, email, user.Email)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
}

func TestGetUser(t *testing.T) {
	hashedPassword, err := util.HashPassword(util.RandomString(12))
	require.NoError(t, err)

	username := util.RandomOwner()
	fullName := util.RandomOwner()
	email := username + "@test.com"

	user, err := createTestUser(username, hashedPassword, fullName, email)
	response, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, response)
	require.Equal(t, user.Username, response.Username)
	require.Equal(t, user.HashedPassword, response.HashedPassword)
	require.Equal(t, user.FullName, response.FullName)
	require.Equal(t, user.Email, response.Email)
	require.WithinDuration(t, user.CreatedAt, response.CreatedAt, time.Second)
}
