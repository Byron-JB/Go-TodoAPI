package models

type FetchTodoResponseDTO struct {
	MetaData MetaData
	Todos    []TodoDto
}

type MetaData struct {
	Skip     int
	Take     int
	NextPage string
}
