package server

import (
	"context"
	todopb "octo-broccoli/protocol/go"
	db "octo-broccoli/server/src/db"
)

func (s *server) Create(ctx context.Context, request *todopb.CreateTodoRequest) (*todopb.CreateTodoResponse, error) {
	todo := request.GetTodo().GetText()

	err := db.Create(todo)
	if err != nil {
		return nil, err
	}

	return &todopb.CreateTodoResponse{
		Status: true,
		Todo: &todopb.Todo{
			Text: todo,
		},
	}, nil
}
