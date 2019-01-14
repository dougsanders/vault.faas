package function

import (
	"encoding/json"
	"fmt"
	"log"
)

type Request struct {
	Environment string
	// This can be Unseal, Access, All, a token, or many tokens
	Filter string
}

// Handle a serverless request
func Handle(bytes []byte) string {
	req := &Request{}
	if err := json.Unmarshal(bytes, req); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Hello, Go. You said: %s", string(bytes))
}
