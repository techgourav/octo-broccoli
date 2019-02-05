package db

import (
	"encoding/json"
	models "octo-broccoli/server/src/models"

	scribble "github.com/nanobox-io/golang-scribble"
)

var db *scribble.Driver

// Init : Initialize a new db
func Init() error {
	if db != nil {
		return nil
	}

	_db, err := scribble.New("db/store", nil)
	if err != nil {
		return err
	}
	db = _db
	return nil
}

// Create : Create a todo to the db
func Create(todo string) error {
	err := db.Write("todo", todo, models.Todo{Text: todo})
	if err != nil {
		return err
	}
	return nil
}

// Read : Read a todo from db
func Read(todo string) (*models.Todo, error) {
	var result models.Todo
	err := db.Read("todo", todo, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ReadAll : Read all todos
func ReadAll() (*[]models.Todo, error) {
	todosDBList, err := db.ReadAll("todo")
	if err != nil {
		return nil, err
	}

	todos := make([]models.Todo, 0)
	for _, todo := range todosDBList {
		newTodo := models.Todo{}
		if err = json.Unmarshal([]byte(todo), &newTodo); err != nil {
			return nil, err
		}
		todos = append(todos, newTodo)
	}

	return &todos, nil
}

// Delete : Delete a particular todo
func Delete(todo string) error {
	err := db.Delete("todo", todo)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAll : Delete all the records
func DeleteAll() error {
	err := db.Delete("todo", "")
	if err != nil {
		return err
	}
	return nil
}
