package db

import (
	models "todoApi/models"
)

func (g *GormDatabase) UpdateTodosInDB(todos []models.TodoDto) ([]models.TodoDto, error) {

	var updatedTodos []models.TodoDto

	for _, todoFromRequest := range todos {

		mappedTodo, err := mapDBObjectFromTodoDto(&todoFromRequest)

		if err != nil {
			return []models.TodoDto{}, err
		}

		g.dbConnection.Save(&mappedTodo)

		updateTodo := mapTodoDtoFromDBObject(&mappedTodo)

		updatedTodos = append(updatedTodos, updateTodo)
	}

	return updatedTodos, nil
}
