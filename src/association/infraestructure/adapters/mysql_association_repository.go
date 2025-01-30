package adapters

import (
	"association/domain"
	"gorm.io/gorm"
)

type MySQLAssociationRepository struct {
	db *gorm.DB
}

func NewMySQLAssociationRepository(db *gorm.DB) *MySQLAssociationRepository {
	return &MySQLAssociationRepository{db: db}
}

func (repo *MySQLAssociationRepository) Create(association domain.Association) error {
	return repo.db.Create(&association).Error
}

func (repo *MySQLAssociationRepository) GetByID(id int) (*domain.Association, error) {
	var association domain.Association
	if err := repo.db.First(&association, id).Error; err != nil {
		return nil, err
	}
	return &association, nil
}

func (repo *MySQLAssociationRepository) GetAll() ([]domain.Association, error) {
	var associations []domain.Association
	if err := repo.db.Find(&associations).Error; err != nil {
		return nil, err
	}
	return associations, nil
}

func (repo *MySQLAssociationRepository) Update(association domain.Association) error {
	return repo.db.Save(&association).Error
}

func (repo *MySQLAssociationRepository) Delete(id int) error {
	return repo.db.Delete(&domain.Association{}, id).Error
}
