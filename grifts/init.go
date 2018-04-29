package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wung-s/firebase_authentication/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
