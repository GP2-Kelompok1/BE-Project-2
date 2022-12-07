package repository

import (
	"errors"
	"immersive-dashboard/features/team"

	"gorm.io/gorm"
)

type teamRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) team.RepositoryInterface {
	return &teamRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *teamRepository) Create(input team.CoreTeam) (row int, err error) {
	teamGorm := fromCore(input)
	tx := repo.db.Create(&teamGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *teamRepository) GetAll() (data []team.CoreTeam, err error) {
	var teams []Team

	tx := repo.db.Find(&teams)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(teams)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (*teamRepository) GetById(id int) (data team.CoreTeam, err error) {
	panic("unimplemented")
}
