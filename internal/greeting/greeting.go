package greeting

import (
	"fmt"
	"strings"
)

// Build creates a simple greeting string used by the demo app.
func Build(name string) (string, error) {
	cleanName := strings.TrimSpace(name)
	if cleanName == "" {
		return "", fmt.Errorf("name must not be empty")
	}

	return fmt.Sprintf("Hello %s, your CI pipeline can build and test this app!", cleanName), nil
}
