package program

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewSearchKey_formSearchQueries(t *testing.T) {
	t.Run("should create correct search quaries alias search key", func(t *testing.T) {
		var item = models.IndexItem{
			Name:       "test",
			Content:    "ls -all",
			Path:       "test.sh",
			Comments:   []string{},
			PathOnDisk: "test.sh",
			Type:       models.ScriptType(models.Alias),
			StartLine:  0,
			EndLine:    0,
		}
		var key = SearchKey{item}
		var actual = key.formSearchQueries()
		var expected = []string{"test.sh/alias test", "test.sh/test"}
		assert.Equal(t, expected, actual)
	})

	t.Run("should create correct search quaries function result", func(t *testing.T) {
		var item = models.IndexItem{
			Name:       "test",
			Content:    "ls -all",
			Path:       "test.sh",
			Comments:   []string{},
			PathOnDisk: "test.sh",
			Type:       models.ScriptType(models.Function),
			StartLine:  0,
			EndLine:    0,
		}
		var key = SearchKey{item}
		var actual = key.formSearchQueries()
		var expected = []string{"test.sh/function test", "test.sh/test"}
		assert.Equal(t, expected, actual)
	})

	t.Run("should create correct search quaries script result", func(t *testing.T) {
		var item = models.IndexItem{
			Name:       "test",
			Content:    "ls -all",
			Path:       "test.sh",
			Comments:   []string{},
			PathOnDisk: "test.sh",
			Type:       models.ScriptType(models.Script),
			StartLine:  0,
			EndLine:    2,
		}
		var key = SearchKey{item}
		var actual = key.formSearchQueries()
		var expected = []string{"test.sh/./test", "test.sh/test"}
		assert.Equal(t, expected, actual)
	})

	t.Run("should create correct search quaries export result", func(t *testing.T) {
		var item = models.IndexItem{
			Name:       "test",
			Content:    "ls -all",
			Path:       "test.sh",
			Comments:   []string{},
			PathOnDisk: "test.sh",
			Type:       models.ScriptType(models.Export),
			StartLine:  0,
			EndLine:    0,
		}
		var key = SearchKey{item}
		var actual = key.formSearchQueries()
		var expected = []string{"test.sh/export test", "test.sh/test"}
		assert.Equal(t, expected, actual)
	})
}

func Test_NewSearchKey_Contains(t *testing.T) {
	t.Run("should return true if key is contained in search term", func(t *testing.T) {
		var item = models.IndexItem{
			Name:       "test",
			Content:    "ls -all",
			Path:       "test.sh",
			Comments:   []string{},
			PathOnDisk: "test.sh",
			Type:       models.ScriptType(models.Alias),
			StartLine:  0,
			EndLine:    0,
		}
		var key = SearchKey{item}
		var actual = key.Contains("test")
		var expected = true
		assert.Equal(t, expected, actual)
	})

	t.Run("should return false if key is contained in search term", func(t *testing.T) {
		var item = models.IndexItem{
			Name:       "test",
			Content:    "ls -all",
			Path:       "test.sh",
			Comments:   []string{},
			PathOnDisk: "test.sh",
			Type:       models.ScriptType(models.Alias),
			StartLine:  0,
			EndLine:    0,
		}
		var key = SearchKey{item}
		var actual = key.Contains("xoxo")
		var expected = false
		assert.Equal(t, expected, actual)
	})
}
