package repository

import (
	"errors"
	"immersive-dashboard/features/mentee"

	"gorm.io/gorm"
)

type menteeRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.RepositoryInterface {
	return &menteeRepository{
		db: db,
	}
}

// GetAll implements mentee.Repository
func (repo *menteeRepository) GetAll() (data []mentee.CoreMentee, err error) {
	var mentees []Mentee

	tx := repo.db.Find(&mentees)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(mentees)
	return dataCore, nil
}

// Create implements user.Repository
func (repo *menteeRepository) Create(input mentee.CoreMentee) (row int, err error) {
	menteeGorm := fromCore(input)
	tx := repo.db.Create(&menteeGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

func (repo *menteeRepository) GetById(id int) (data mentee.CoreMentee, err error) {
	var IdMentee Mentee
	var IdMenteeCore = mentee.CoreMentee{}
	IdMentee.ID = uint(id)
	tx := repo.db.First(&IdMentee, IdMentee.ID)
	if tx.Error != nil {
		return IdMenteeCore, tx.Error
	}
	IdMenteeCore = IdMentee.toCore()
	return IdMenteeCore, nil
}

func (repo *menteeRepository) UpdateMentee(datacore mentee.CoreMentee, id int) (err error) {
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

func (repo *menteeRepository) DeleteMentee(id int) (row int, err error) {
	idMentee := Mentee{}

	tx := repo.db.Delete(&idMentee, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete mentee by id failed")
	}
	return int(tx.RowsAffected), nil
}
