package db

import (
	"errors"
	"fmt"
	models "todoApi/models"

	"gorm.io/gorm"
)

func DeleteTodoFromDb(id int) error {

	err := OpenDbConnection()
	if err != nil {
		return err
	}

	err = dbConnection.First(&models.TblTodo{}, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("todo with id %d not found", id)
	}

	todo := dbConnection.Where("id = ?", id).Delete(&models.TblTodo{})

	if todo == nil {
		return errors.New("there was not todo with that id")
	}

	return nil
}
