package storage

import (
	"database/sql"
	"time"
	"user_service/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(req *proto.User) (*proto.UserResponse, error) {
	dbQuery := `
		INSERT INTO users
			(first_name, last_name, email)
		VALUES
			($1, $2, $3)
	`
	_, err := u.db.Exec(dbQuery, req.FirstName, req.LastName, req.Email)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{
		Message: "User created successfully",
	}, nil
}

func (u *UserRepo) UpdateUser(req *proto.User) (*proto.UserResponse, error) {
	dbQuery := `
		UPDATE users
		SET 
			first_name = $1, last_name = $2, email = $3, updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $4
	`
	_, err := u.db.Exec(dbQuery, req.FirstName, req.LastName, req.Email, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{
		Message: "User updated successfully",
	}, nil
}

func (u *UserRepo) DeleteUser(req *proto.UserRequest) (*proto.UserResponse, error) {
	dbQuery := `
		UPDATE users
		SET 
			deleted_at = $1
		WHERE
			id = $2
	`
	_, err := u.db.Exec(dbQuery, time.Now().Unix(), req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{
		Message: "User deleted successfully",
	}, nil
}

func (u *UserRepo) GetUser(req *proto.UserRequest) (*proto.User, error) {
	var (
		id        string
		firstName string
		lastName  string
		email     string
		createdAt time.Time
		updatedAt time.Time
		deletedAt int64
	)

	dbQuery := `
		SELECT * FROM users WHERE id = $1
	`

	row := u.db.QueryRow(dbQuery, req.Id)
	err := row.Scan(
		&id,
		&firstName,
		&lastName,
		&email,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return nil, err
	}

	user := &proto.User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
		DeletedAt: deletedAt,
	}

	return user, nil
}
