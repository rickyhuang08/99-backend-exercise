package sql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rickyhuang08/99-backend-exercise/helpers"
	"github.com/rickyhuang08/99-backend-exercise/internal/entity"
)

// UserRepository handles DB interactions
type UserRepository struct {
	DB *sql.DB
	TimeModule helpers.TimeProvider
}

// NewUserRepository initializes the repo with dummy data
func NewUserRepository(db *sql.DB, timeModule helpers.TimeProvider) *UserRepository {
	return &UserRepository{DB: db, TimeModule: timeModule}
}

func (r *UserRepository) Create(userPayload entity.UserPayload) (*entity.User, error) {
	now := r.TimeModule.Now().UnixMicro()

	// Insert user into the database
	res, err := r.DB.Exec("INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		userPayload.Name, userPayload.Email, userPayload.Password, now, now)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %w", err)
	}
	user := entity.User{
		ID:        int(id),
		Name:     userPayload.Name,
		Email:    userPayload.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return &user, nil
}

// GetUser retrieves a user by ID or name optional parameter
// If both are empty, it returns an error.
// If ID is provided, it fetches by ID; if name is provided, it fetches by name.
func (r *UserRepository) GetUser(params entity.GetUserParam) ([]entity.User, error) {
	// Validate pagination input
	if params.Page <= 0 {
		params.Page = helpers.DefaultPage
	}
	if params.PageSize <= 0 {
		params.PageSize = helpers.DefaultPageSize
	}
	offset := (params.Page - 1) * params.PageSize

	var users []entity.User
	query := "SELECT id, name, email, created_at, updated_at FROM users"
	args := []interface{}{}

	if params.ID > 0 || params.Name != "" {
		query += " WHERE"
	}
	
	if params.ID > 0 {
		query += " id = ?"
		args = append(args, params.ID)
	} else if params.Name != "" {
		query += " name = ?"
		args = append(args, params.Name)
	}

	if params.ID > 0 {
		query += " LIMIT 1"
	} else {
		query += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
		args = append(args, params.PageSize, offset)
	}
	fmt.Println("Query : ", query)
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email %s: %w", email, err)
		}
		return nil, fmt.Errorf("failed to query user by email: %w", err)
	}
	return &user, nil
}