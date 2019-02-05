package server

import (
	todopb "octo-broccoli/protocol/go"
	db "octo-broccoli/server/src/db"
)

func (s *server) GetAll(request *todopb.GetTodoRequest, stream todopb.TodoService_GetAllServer) error {
	todo := request.GetTodo().GetText()

	if len(todo) == 0 {
		// Stream all records
		records, err := db.ReadAll()
		if err != nil {
			return err
		}

		for _, record := range *records {
			response := &todopb.GetTodoResponse{
				Todo: &todopb.Todo{
					Text: record.Text,
				},
			}

			stream.Send(response)
		}
	} else {
		// Stream matching record
		record, err := db.Read(todo)
		if err != nil {
			return err
		}

		response := &todopb.GetTodoResponse{
			Todo: &todopb.Todo{
				Text: record.Text,
			},
		}

		stream.Send(response)
	}

	return nil
}
