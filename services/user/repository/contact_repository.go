package repository

import (
	"database/sql"

	"github.com/meuapoio/services/user/models"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Create(userID string, contact *models.CreateContactRequest) (*models.EmergencyContact, error) {
	newContact := &models.EmergencyContact{}

	query := `
		INSERT INTO emergency_contacts (user_id, name, phone, relationship, is_primary)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, name, phone, relationship, is_primary, created_at
	`

	err := r.db.QueryRow(
		query, userID, contact.Name, contact.Phone, contact.Relationship, contact.IsPrimary,
	).Scan(
		&newContact.ID, &newContact.UserID, &newContact.Name, &newContact.Phone,
		&newContact.Relationship, &newContact.IsPrimary, &newContact.CreatedAt,
	)

	return newContact, err
}

func (r *ContactRepository) GetByUserID(userID string) ([]*models.EmergencyContact, error) {
	query := `
		SELECT id, user_id, name, phone, relationship, is_primary, created_at
		FROM emergency_contacts 
		WHERE user_id = $1 
		ORDER BY is_primary DESC, created_at DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*models.EmergencyContact
	for rows.Next() {
		contact := &models.EmergencyContact{}
		err := rows.Scan(
			&contact.ID, &contact.UserID, &contact.Name, &contact.Phone,
			&contact.Relationship, &contact.IsPrimary, &contact.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactRepository) GetByID(id, userID string) (*models.EmergencyContact, error) {
	contact := &models.EmergencyContact{}
	query := `
		SELECT id, user_id, name, phone, relationship, is_primary, created_at
		FROM emergency_contacts 
		WHERE id = $1 AND user_id = $2
	`

	err := r.db.QueryRow(query, id, userID).Scan(
		&contact.ID, &contact.UserID, &contact.Name, &contact.Phone,
		&contact.Relationship, &contact.IsPrimary, &contact.CreatedAt,
	)

	return contact, err
}

func (r *ContactRepository) Update(id, userID string, req *models.UpdateContactRequest) error {
	query := `
		UPDATE emergency_contacts 
		SET name = COALESCE($3, name),
		    phone = COALESCE($4, phone),
		    relationship = COALESCE($5, relationship),
		    is_primary = COALESCE($6, is_primary)
		WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.Exec(query, id, userID, req.Name, req.Phone, req.Relationship, req.IsPrimary)
	return err
}

func (r *ContactRepository) Delete(id, userID string) error {
	query := `DELETE FROM emergency_contacts WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, id, userID)
	return err
}
