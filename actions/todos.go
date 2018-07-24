package actions

import (
	"github.com/arschles/hd/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Todo)
// DB Table: Plural (todos)
// Resource: Plural (Todos)
// Path: Plural (/todos)
// View Template Folder: Plural (/templates/todos/)

// TodosResource is the resource for the Todo model
type TodosResource struct {
	buffalo.Resource
}

// List gets all Todos. This function is mapped to the path
// GET /todos
func (v TodosResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	todos := &models.Todos{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Todos from the DB
	if err := q.All(todos); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, todos))
}

// Show gets the data for one Todo. This function is mapped to
// the path GET /todos/{todo_id}
func (v TodosResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	// To find the Todo the parameter todo_id is used.
	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, todo))
}

// New renders the form for creating a new Todo.
// This function is mapped to the path GET /todos/new
func (v TodosResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Todo{}))
}

// Create adds a Todo to the DB. This function is mapped to the
// path POST /todos
func (v TodosResource) Create(c buffalo.Context) error {
	// Allocate an empty Todo
	todo := &models.Todo{}

	// Bind todo to the html form elements
	if err := c.Bind(todo); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(todo)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, todo))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Todo was created successfully")

	// and redirect to the todos index page
	return c.Render(201, r.Auto(c, todo))
}

// Edit renders a edit form for a Todo. This function is
// mapped to the path GET /todos/{todo_id}/edit
func (v TodosResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, todo))
}

// Update changes a Todo in the DB. This function is mapped to
// the path PUT /todos/{todo_id}
func (v TodosResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Todo to the html form elements
	if err := c.Bind(todo); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(todo)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, todo))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Todo was updated successfully")

	// and redirect to the todos index page
	return c.Render(200, r.Auto(c, todo))
}

// Destroy deletes a Todo from the DB. This function is mapped
// to the path DELETE /todos/{todo_id}
func (v TodosResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	// To find the Todo the parameter todo_id is used.
	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(todo); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Todo was destroyed successfully")

	// Redirect to the todos index page
	return c.Render(200, r.Auto(c, todo))
}
