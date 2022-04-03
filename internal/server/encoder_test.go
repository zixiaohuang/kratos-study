package server

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHTTPStruct(t *testing.T) {
	a := &HTTPError{
		Errors: make(map[string][]string),
	}
	a.Errors["body"] = []string{"can't be empty"}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Printf("%s", string(b))
}