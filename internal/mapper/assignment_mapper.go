package mapper

import (
	"ddl-bot/internal/entity"
	"ddl-bot/pkg/utils"
	"gorm.io/gorm"
)

type AssignmentMapper struct {
	DB *gorm.DB
}

func (mapper *AssignmentMapper) Init() error {
	err := mapper.DB.AutoMigrate(&entity.Assignment{})
	if err != nil {
		return WrapDBError(err)
	}
	return nil
}

func (mapper *AssignmentMapper) GetDB() *gorm.DB {
	return mapper.DB
}

func (mapper *AssignmentMapper) List(tx *gorm.DB) ([]entity.Assignment, error) {
	var items []entity.Assignment
	tx.Find(&items)
	if tx.Error != nil {
		return nil, WrapDBError(tx.Error)
	}
	return items, nil
}

func (mapper *AssignmentMapper) ListToDo(tx *gorm.DB) ([]entity.Assignment, error) {
	var items []entity.Assignment
	tx.Where("status=?", "待做").Find(&items)
	if tx.Error != nil {
		return nil, WrapDBError(tx.Error)
	}
	return items, nil
}

func (mapper *AssignmentMapper) Get(tx *gorm.DB, assignmentID string) (*entity.Assignment, error) {
	var item = new(entity.Assignment)
	tx.First(item)
	if tx.Error != nil {
		return nil, WrapDBError(tx.Error)
	}
	if utils.IsEmpty(item.AssignmentID) {
		return nil, nil
	}
	return item, nil
}

func (mapper *AssignmentMapper) Save(db *gorm.DB, item *entity.Assignment) error {
	fromDB, err := mapper.Get(db, item.AssignmentID)
	if err != nil {
		return err
	}
	if fromDB == nil {
		db.Create(item)
	} else {
		db.Save(item)
	}
	if db.Error != nil {
		return WrapDBError(db.Error)
	}
	return nil
}
