package actions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"golang.org/x/net/context"
)

// Authenticate will ensure only authenticated users gain access to protected endpoints
func Authenticate(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		// do some work before calling the next handler
		client, err := FirebaseApp.Auth(context.Background())

		idToken := c.Request().Header.Get("Authorization")
		idToken = strings.Replace(idToken, `Bearer `, "", 1)
		if ENV == "development" || ENV == "test" {
			fmt.Println("Authorization", idToken)
		}
		_, err = client.VerifyIDToken(idToken)
		if err != nil {
			fmt.Printf("error verifying ID token: %v\n", err)
			c.Response().WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(c.Response()).Encode("Missing or invalid token")
			return err
		}

		err = next(c)
		return err
	}
}
