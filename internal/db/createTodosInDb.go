package db

import (
	"todoApi/models"
)

func (g *GormDatabase) SaveTodosToDb(todos []models.TodoDto) ([]models.TodoDto, error) {

	var createdTodos []models.TodoDto

	var todosToSave []models.TblTodo

	// Insert each Todo into the database
	for _, todo := range todos {
		entryToSave, err := mapDBObjectFromTodoDto(&todo)

		if err != nil {
			return nil, err
		}

		todosToSave = append(todosToSave, entryToSave)
	}

	if err := g.dbConnection.Create(&todosToSave).Error; err != nil {
		return nil, err
	}

	for _, createdTodo := range todosToSave {
		createdTodos = append(createdTodos, mapTodoDtoFromDBObject(&createdTodo))
	}

	return createdTodos, nil
}
