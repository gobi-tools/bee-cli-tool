package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getHomeUrl(t *testing.T) {
	t.Run("should return a non null home url", func(t *testing.T) {
		var result string = getHomeUrl()
		assert.NotEmpty(t, result)
	})
}

func Test_getDataUrl(t *testing.T) {
	t.Run("should return the correct url", func(t *testing.T) {
		var result string = getDataUrl()
		assert.Contains(t, result, ".local/bin/bee/data.json")
	})
}

func Test_getSourcesUrl(t *testing.T) {
	t.Run("should return the correct url", func(t *testing.T) {
		var result string = getSourcesUrl()
		assert.Contains(t, result, ".local/bin/bee/sources.json")
	})
}

func Test_getLastCommandUrl(t *testing.T) {
	t.Run("should return the correct url", func(t *testing.T) {
		var result string = getLastCommandUrl()
		assert.Contains(t, result, ".local/bin/bee/lastcommand")
	})
}
