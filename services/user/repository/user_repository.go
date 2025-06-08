package repository

import (
	"database/sql"

	"github.com/meuapoio/services/user/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash, full_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at, is_active
	`

	return r.db.QueryRow(
		query, user.Username, user.Email, user.PasswordHash, user.FullName,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.IsActive)
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, username, email, password_hash, full_name, birth_date, 
		       phone, profile_image_url, created_at, updated_at, is_active
		FROM users 
		WHERE email = $1 AND is_active = true
	`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.FullName, &user.BirthDate, &user.Phone, &user.ProfileImageURL,
		&user.CreatedAt, &user.UpdatedAt, &user.IsActive,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, username, email, password_hash, full_name, birth_date, 
		       phone, profile_image_url, created_at, updated_at, is_active
		FROM users 
		WHERE id = $1 AND is_active = true
	`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash,
		&user.FullName, &user.BirthDate, &user.Phone, &user.ProfileImageURL,
		&user.CreatedAt, &user.UpdatedAt, &user.IsActive,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(id string, req *models.UpdateUserRequest) error {
	query := `
		UPDATE users 
		SET full_name = COALESCE($2, full_name),
		    birth_date = COALESCE($3, birth_date),
		    phone = COALESCE($4, phone),
		    profile_image_url = COALESCE($5, profile_image_url),
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND is_active = true
	`

	_, err := r.db.Exec(query, id, req.FullName, req.BirthDate, req.Phone, req.ProfileImageURL)
	return err
}

func (r *UserRepository) SoftDelete(id string) error {
	query := `UPDATE users SET is_active = false WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) EmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := r.db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

func (r *UserRepository) UsernameExists(username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := r.db.QueryRow(query, username).Scan(&exists)
	return exists, err
}
