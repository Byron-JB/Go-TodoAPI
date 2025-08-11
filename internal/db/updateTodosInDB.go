package db

import (
	models "todoApi/models"
)

func UpdateTodosInDB(todos []models.TodoDto) ([]models.TodoDto, error) {
	err := OpenDbConnection()

	var updatedTodos []models.TodoDto

	if err != nil {
		return []models.TodoDto{}, err
	}

	for _, todoFromRequest := range todos {

		mappedTodo, err := mapDBObjectFromTodoDto(&todoFromRequest)

		if err != nil {
			return []models.TodoDto{}, err
		}

		dbConnection.Save(&mappedTodo)

		updateTodo := mapTodoDtoFromDBObject(&mappedTodo)

		updatedTodos = append(updatedTodos, updateTodo)
	}

	return updatedTodos, nil
}
