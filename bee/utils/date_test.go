package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CurrentTime(t *testing.T) {
	t.Run("should return a non-zero timestamp", func(t *testing.T) {
		var result int64 = CurrentTime()
		assert.Greater(t, result, int64(0))
	})
}
