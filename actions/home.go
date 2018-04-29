package actions

import "github.com/gobuffalo/buffalo"

// OpenHandler does not require authentication
func OpenHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is an open endpoint!"}))
}

// SecureHandler requires authentication
func SecureHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is a secured endpoint !"}))
}
