package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	id, err := GenerateID()
	assert.Equal(t, err, nil, "err should be nil")
	fmt.Println(id)
}
