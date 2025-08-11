package db

import (
	"errors"
	"time"
	models "todoApi/models"
)

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	//TODO implement me
	panic("implement me")
}

const dateLayout = "2006-01-02"      // MySQL DATE
const timestampLayout = time.RFC3339 // e.g., 2023-10-01T12:00:00Z

func mapDBObjectFromTodoDto(dto *models.TodoDto) (models.TblTodo, error) {

	//log.Printf(dto.CompletedAt)
	layout := "2006-01-02" // matches DATE in MySQL

	var completedAt *time.Time
	if dto.CompletedAt != "" {
		t, err := time.Parse(layout, dto.CompletedAt)
		if err != nil {
			return models.TblTodo{}, errors.New("completed Date is not valid")
		}
		completedAt = &t
	}

	// Parse DueDate if present (treat as optional; make it required if your domain needs it)
	var dueDate time.Time
	var err error
	if dto.DueDate != "" {
		dueDate, err = time.Parse(dateLayout, dto.DueDate)
		if err != nil {
			return models.TblTodo{}, errors.New("due Date is not valid")
		}
	} else {
		// If DueDate is required, return an error instead:
		return models.TblTodo{}, errors.New("due_date is required")
	}

	// Prefer parsing CreatedAt if provided; otherwise fallback to now
	createdAt := time.Now()
	if dto.CreatedAt != "" {
		if t, err := time.Parse(timestampLayout, dto.CreatedAt); err == nil {
			createdAt = t
		}
	}

	return models.TblTodo{
		StrTitle:       dto.Title,
		StrDescription: dto.Description,
		IPriority:      &dto.Priority,
		DtCreatedAt:    createdAt,
		DtCompleted:    completedAt,
		DtDueDate:      dueDate,
	}, nil
}

func mapTodoDtoFromDBObject(dbObj *models.TblTodo) models.TodoDto {
	// Layouts must match what the DTO expects
	dateLayout := "2006-01-02"
	timestampLayout := "2006-01-02 15:04:05"

	dto := models.TodoDto{
		ID:          int(dbObj.ID),
		Title:       dbObj.StrTitle,
		Description: dbObj.StrDescription,
		Priority:    *dbObj.IPriority,
		CreatedAt:   dbObj.DtCreatedAt.Format(timestampLayout),
		DueDate:     dbObj.DtDueDate.Format(dateLayout),
	}

	// Handle nullable CompletedAt
	if dbObj.DtCompleted != nil {
		dto.CompletedAt = dbObj.DtCompleted.Format(dateLayout)
	} else {
		dto.CompletedAt = ""
	}

	return dto
}
