package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsHttpUrl(t *testing.T) {
	t.Run("should return true for HTTP url", func(t *testing.T) {
		var path = "http://my.examole.com"
		var result = IsHttpUrl(path)
		assert.True(t, result)
	})

	t.Run("should return true for HTTPS url", func(t *testing.T) {
		var path = "https://my.examole.com"
		var result = IsHttpUrl(path)
		assert.True(t, result)
	})

	t.Run("should return false for file path", func(t *testing.T) {
		var path = "/usr/path/to/file"
		var result = IsHttpUrl(path)
		assert.False(t, result)
	})

	t.Run("should return false empty", func(t *testing.T) {
		var path = ""
		var result = IsHttpUrl(path)
		assert.False(t, result)
	})

	t.Run("should return false for garbage", func(t *testing.T) {
		var path = "\\////safas"
		var result = IsHttpUrl(path)
		assert.False(t, result)
	})
}

func Test_FileName(t *testing.T) {
	t.Run("should return empty dot given empty input", func(t *testing.T) {
		var path = ""
		var result = FileName(path)
		var expected = "."
		assert.Equal(t, expected, result)
	})

	t.Run("should return the same input given garbage input", func(t *testing.T) {
		var path = "\\\\//aasaas;;;s;a"
		var result = FileName(path)
		var expected = "aasaas;;;s;a"
		assert.Equal(t, expected, result)
	})

	t.Run("should return a file name given valid path", func(t *testing.T) {
		var path = "/Users/test.test/my/file.text"
		var result = FileName(path)
		var expected = "file.text"
		assert.Equal(t, expected, result)
	})

	t.Run("should return a file name given valid path", func(t *testing.T) {
		var path = "https://test.com/abc?query=abc"
		var result = FileName(path)
		var expected = "abc"
		assert.Equal(t, expected, result)
	})

	t.Run("should return a file name given valid path", func(t *testing.T) {
		var path = "http://test.com/test/file.sh"
		var result = FileName(path)
		var expected = "file.sh"
		assert.Equal(t, expected, result)
	})

	t.Run("should return the original path in case it is invalid", func(t *testing.T) {
		var path = "http://\\////safas"
		var result = FileName(path)
		var expected = "http://\\////safas"
		assert.Equal(t, expected, result)
	})
}

func Test_FileNameWithoutExtTrimSuffix(t *testing.T) {
	t.Run("should trim the suffix extension of a file name", func(t *testing.T) {
		var input = "file.ext"
		var result = FileNameWithoutExtTrimSuffix(input)
		var expected = "file"
		assert.Equal(t, expected, result)
	})

	t.Run("should return same file name in case it does not have an extension", func(t *testing.T) {
		var input = "file"
		var result = FileNameWithoutExtTrimSuffix(input)
		var expected = "file"
		assert.Equal(t, expected, result)
	})

	t.Run("should return empty in case of empty file name", func(t *testing.T) {
		var input = ""
		var result = FileNameWithoutExtTrimSuffix(input)
		var expected = ""
		assert.Equal(t, expected, result)
	})

	t.Run("should only eliminate last suffix", func(t *testing.T) {
		var input = "file.ext1.ext2"
		var result = FileNameWithoutExtTrimSuffix(input)
		var expected = "file.ext1"
		assert.Equal(t, expected, result)
	})
}
