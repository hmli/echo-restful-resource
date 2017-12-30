package echo_restful_resource

import "github.com/labstack/echo"

type Resource interface {

	// GET /foo
	Index(c echo.Context) error

	// GET /foo/create
	Create(c echo.Context) error

	// POST /foo
	Store(c echo.Context) error

	// GET /foo/{id}
	Show(c echo.Context) error

	// GET /foo/{id}/edit
	Edit(c echo.Context) error

	// PUT /foo/{id}
	Update(c echo.Context) error

	// DELETE /foo/{id}
	Destroy(c echo.Context) error
}
