package db

import models "todoApi/models"

func FetchTodosFromDb(skip int, take int) ([]models.TodoDto, error) {

	var todos []models.TodoDto

	err := OpenDbConnection()

	if err != nil {
		return nil, err
	}

	var todoList []models.TblTodo
	result := dbConnection.Offset(skip).Limit(take).Find(&todoList)

	if result.Error != nil {
		return nil, err
	}

	for _, todoFromDB := range todoList {
		todos = append(todos, mapTodoDtoFromDBObject(&todoFromDB))
	}

	return todos, nil
}
