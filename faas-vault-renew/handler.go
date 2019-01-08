package function

import (
	"fmt"

	"github.com/google/uuid"
)

// Handle a serverless request
func Handle(req []byte) string {
	// Pull repository

	deptest, _ := uuid.NewUUID()

	return fmt.Sprintf("Hello, Go. You said: %s %s", string(req), deptest.String())
}
