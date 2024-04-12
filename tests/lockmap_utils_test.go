package tests

import (
	"file-service/utils"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGetLock(t *testing.T) {
	key := "user:1"
	putRes := utils.GetLock(key)
	assert.Equal(t, putRes, true)
}
