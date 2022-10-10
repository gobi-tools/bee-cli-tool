package program

import (
	"bee/bbee/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewSearchResult(t *testing.T) {
	t.Run("should create new alias search result", func(t *testing.T) {
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
		var actual = NewSearchResult(item)
		var expected = NewAliasSearchResult(item)
		assert.Equal(t, expected, actual)
		assert.Equal(t, "   [#659acc]alias [#8cdbff]test", actual.mainText)
		assert.Equal(t, "", actual.secondaryText)
		assert.Equal(t, "test.sh/test", actual.previewTitle)
		assert.Equal(t, "ls -all", actual.previewContent)
		assert.Equal(t, "ls -all", actual.command)
		assert.Equal(t, "test.sh", actual.pathOnDisk)
		assert.Equal(t, 0, actual.startLine)
		assert.Equal(t, 0, actual.endLine)
		assert.Equal(t, false, actual.noHighlight)
		assert.Equal(t, Item, actual.resultType)
	})

	t.Run("should create new alias function result", func(t *testing.T) {
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
		var actual = NewSearchResult(item)
		var expected = NewFunctionSearchResult(item)
		assert.Equal(t, expected, actual)
		assert.Equal(t, "   [#fbdfb5]function [#8bdaff]test", actual.mainText)
		assert.Equal(t, "", actual.secondaryText)
		assert.Equal(t, "test.sh/test", actual.previewTitle)
		assert.Equal(t, "ls -all", actual.previewContent)
		assert.Equal(t, "ls -all\ntest", actual.command)
		assert.Equal(t, "test.sh", actual.pathOnDisk)
		assert.Equal(t, 0, actual.startLine)
		assert.Equal(t, 0, actual.endLine)
		assert.Equal(t, false, actual.noHighlight)
		assert.Equal(t, Item, actual.resultType)
	})

	t.Run("should create new alias script result", func(t *testing.T) {
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
		var actual = NewSearchResult(item)
		var expected = NewScriptSearchResult(item)
		assert.Equal(t, expected, actual)
		assert.Equal(t, "   [#e9fdac]./test", actual.mainText)
		assert.Equal(t, "", actual.secondaryText)
		assert.Equal(t, "test.sh/test", actual.previewTitle)
		assert.Equal(t, "ls -all", actual.previewContent)
		assert.Equal(t, "ls -all", actual.command)
		assert.Equal(t, "test.sh", actual.pathOnDisk)
		assert.Equal(t, 0, actual.startLine)
		assert.Equal(t, 2, actual.endLine)
		assert.Equal(t, true, actual.noHighlight)
		assert.Equal(t, Item, actual.resultType)
	})

	t.Run("should create new export script result", func(t *testing.T) {
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
		var actual = NewSearchResult(item)
		var expected = NewExportSearchResult(item)
		assert.Equal(t, expected, actual)
		assert.Equal(t, "   [#ff5e5f]export [#d8ea9f]test", actual.mainText)
		assert.Equal(t, "", actual.secondaryText)
		assert.Equal(t, "test.sh/test", actual.previewTitle)
		assert.Equal(t, "ls -all", actual.previewContent)
		assert.Equal(t, "test.sh", actual.command)
		assert.Equal(t, "test.sh", actual.pathOnDisk)
		assert.Equal(t, 0, actual.startLine)
		assert.Equal(t, 0, actual.endLine)
		assert.Equal(t, false, actual.noHighlight)
		assert.Equal(t, Item, actual.resultType)
	})
}
