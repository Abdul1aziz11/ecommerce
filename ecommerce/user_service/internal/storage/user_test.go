package storage

import (
	"testing"
	"time"
	"user-service/proto"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepo(db)
	user := &proto.User{
		FirstName: "Abdullaziz",
		LastName:  "Mamadjanov",
		Email:     "abdullazizmamadjanov@example.com",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.FirstName, user.LastName, user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.CreateUser(user)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "User created successfully", resp.Message)
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepo(db)
	user := &proto.User{
		Id:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}

	mock.ExpectExec("UPDATE users").
		WithArgs(user.FirstName, user.LastName, user.Email, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.UpdateUser(user)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "User updated successfully", resp.Message)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepo(db)
	req := &proto.UserRequest{
		Id: "1",
	}

	mock.ExpectExec("UPDATE users SET deleted_at = NOW() WHERE id = ?").
		WithArgs(req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.DeleteUser(req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, "User deleted successfully", resp.Message)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewUserRepo(db)
	req := &proto.UserRequest{
		Id: "1",
	}

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "created_at", "updated_at", "deleted_at"}).
		AddRow("1", "Abdullaziz", "Mamadjanov", "abdullazizmamadjanov@example.com", time.Now(), time.Now(), nil)

	mock.ExpectQuery("SELECT id, first_name, last_name, email, created_at, updated_at, deleted_at FROM users WHERE id = ?").
		WithArgs(req.Id).
		WillReturnRows(rows)

	user, err := repo.GetUser(req)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "1", user.Id)
	require.Equal(t, "Abdullaziz", user.FirstName)
	require.Equal(t, "Mamadjanov", user.LastName)
	require.Equal(t, "abdullazizmamadjanov@example.com", user.Email)
	require.IsType(t, &timestamppb.Timestamp{}, user.CreatedAt)
	require.IsType(t, &timestamppb.Timestamp{}, user.UpdatedAt)
	require.Nil(t, user.DeletedAt)
}
