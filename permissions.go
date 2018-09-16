package keylogger

import (
	"os/user"

	"github.com/pkg/errors"
)

// Root defines required permissions to run the application.
type Root struct{}

// Ensure method checks if the application is running with root permissions.
func (r *Root) Ensure() error {
	u, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "can't read the current user")
	}

	if u.Uid != "0" {
		return errors.New("root permissions are required")
	}
	return nil
}
