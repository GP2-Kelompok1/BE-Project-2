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
func (repo *classRepository) GetById(id int) (data class.CoreClass, err error) {
	var IdClass Class
	var IdClassCore = class.CoreClass{}
	IdClass.ID = uint(id)
	tx := repo.db.First(&IdClass, IdClass.ID)
	if tx.Error != nil {
		return IdClassCore, tx.Error
	}
	IdClassCore = IdClass.toCore()
	return IdClassCore, nil
}

func (repo *classRepository) UpdateClass(datacore class.CoreClass, id int) (err error) {
	userGorm := fromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update user failed")
	}
	return nil
}

func (repo *classRepository) DeleteClass(id int) (row int, err error) {
	idClass := Class{}

	tx := repo.db.Delete(&idClass, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete class by id failed")
	}
	return int(tx.RowsAffected), nil
}
