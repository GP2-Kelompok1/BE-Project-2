package repository

import (
	"errors"
	"immersive-dashboard/features/class"

	"gorm.io/gorm"
)

type classRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.RepositoryInterface {
	return &classRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *classRepository) Create(input class.CoreClass) (row int, err error) {
	classGorm := fromCore(input)
	tx := repo.db.Create(&classGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *classRepository) GetAll() (data []class.CoreClass, err error) {
	var classes []Class

	tx := repo.db.Find(&classes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(classes)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (*classRepository) GetById(id int) (data class.CoreClass, err error) {
	panic("unimplemented")
}
