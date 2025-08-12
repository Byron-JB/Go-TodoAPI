package db

import models "todoApi/models"

// FetchTodosFromDb Fetch a list of todos from the database
func (g *GormDatabase) FetchTodosFromDb(skip int, take int) ([]models.TodoDto, error) {

	var todos []models.TodoDto

	var todoList []models.TblTodo
	result := g.dbConnection.Offset(skip).Limit(take).Find(&todoList)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, todoFromDB := range todoList {
		todos = append(todos, mapTodoDtoFromDBObject(&todoFromDB))
	}

	return todos, nil
}
