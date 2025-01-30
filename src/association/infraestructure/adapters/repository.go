package adapters

import (
	"association/domain"
	"association/domain/ports"
	"database/sql"

)

type RepositoryImpl struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) ports.Repository {
	return &RepositoryImpl{DB: db}
}

func (repo *RepositoryImpl) Create(association domain.Association) error {
	_, err := repo.DB.Exec("INSERT INTO associations (name, description, location, contact) VALUES (?, ?, ?, ?)",
		association.Name, association.Description, association.Location, association.Contact)
	return err
}

func (repo *RepositoryImpl) GetByID(id int) (domain.Association, error) {
	var association domain.Association
	err := repo.DB.QueryRow("SELECT id, name, description, location, contact FROM associations WHERE id = ?", id).
		Scan(&association.ID, &association.Name, &association.Description, &association.Location, &association.Contact)
	return association, err
}

func (repo *RepositoryImpl) GetAll() ([]domain.Association, error) {
	rows, err := repo.DB.Query("SELECT id, name, description, location, contact FROM associations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var associations []domain.Association
	for rows.Next() {
		var a domain.Association
		if err := rows.Scan(&a.ID, &a.Name, &a.Description, &a.Location, &a.Contact); err != nil {
			return nil, err
		}
		associations = append(associations, a)
	}
	return associations, nil
}

func (repo *RepositoryImpl) Update(association domain.Association) error {
	_, err := repo.DB.Exec("UPDATE associations SET name=?, description=?, location=?, contact=? WHERE id=?",
		association.Name, association.Description, association.Location, association.Contact, association.ID)
	return err
}

func (repo *RepositoryImpl) Delete(id int) error {
	_, err := repo.DB.Exec("DELETE FROM associations WHERE id=?", id)
	return err
}
