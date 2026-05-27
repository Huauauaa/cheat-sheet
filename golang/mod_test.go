package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMod(t *testing.T) {
	assert.Equal(t, 10%3, 1)
	assert.Equal(t, 10%-3, 1)
	assert.Equal(t, -10%3, -1)
	assert.Equal(t, -10%-3, -1)
}
