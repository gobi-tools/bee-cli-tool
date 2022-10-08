package program

import (
	"bee/bbee/models"
	"bee/bbee/utils"
	"sort"

	"github.com/samber/lo"
)

// The Search Controller contains data & exposes functions related
// to the search through all of the aliases, functions and scripts
// a user has saved
type SearchController struct {
	elems        []models.IndexItem
	results      []SearchResult
	currentIndex int
	totalLen     int
}

func NewEmptySearchController() *SearchController {
	items := []models.IndexItem{}
	return NewSearchController(items)
}

func NewSearchController(elems []models.IndexItem) *SearchController {
	controller := new(SearchController)
	controller.elems = elems
	controller.results = controller.formResults(elems)
	controller.totalLen = len(elems)
	controller.resetCurrentIndex()
	return controller
}

func (c *SearchController) search(term string) {
	// filter
	var filtered = lo.Filter(c.elems, func(item models.IndexItem, i int) bool {
		var key = SearchKey{item: item}
		return key.Contains(term)
	})

	c.results = c.formResults(filtered)
	c.resetCurrentIndex()
}

func (c *SearchController) formResults(items []models.IndexItem) []SearchResult {
	var result = []SearchResult{}

	var paths = models.UniquePaths(items)

	for _, path := range paths {
		result = append(result, NewSearchCategory(path.A, path.B))
		var filtered = models.FilterByPath(items, path.B /* B = path on disk */)

		// always sort according to start line
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].StartLine < filtered[j].StartLine
		})

		for _, item := range filtered {
			result = append(result, NewSearchResult(item))
		}
	}

	return result
}

func (c *SearchController) moveDown() {
	c.currentIndex = utils.Min(c.currentIndex+1, len(c.results)-1)
}

func (c *SearchController) moveUp() {
	c.currentIndex = utils.Max(c.currentIndex-1, 0)
}

func (c *SearchController) resetCurrentIndex() {
	c.currentIndex = 0
}

func (c SearchController) getCurrentItem() SearchResult {
	if c.currentIndex >= 0 && c.currentIndex < len(c.results) {
		return c.results[c.currentIndex]
	} else {
		return NewEmptySearchResult()
	}
}

func (c SearchController) getNearestCategoryResult() SearchResult {
	if c.currentIndex >= 0 && c.currentIndex < len(c.results) {
		var i = c.currentIndex
		for i >= 0 {
			currentResult := c.results[i]
			if currentResult.resultType == SearchResultType(Category) {
				return currentResult
			}
			i -= 1
		}

	}
	return NewEmptySearchResult()
}

func (c SearchController) getNumberOfSearchResults() int {
	filtered := lo.Filter(c.results, func(result SearchResult, i int) bool {
		return result.resultType == SearchResultType(Item)
	})
	return len(filtered)
}
