package slices

import "github.com/BibleBot/backend/internal/models"

// Index finds index of string in []string, otherwise returns -1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// MapSlice maps a []string to a []int
func MapSlice(vs []string, f func(string, int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, i)
	}
	return vsm
}

// Any returns whether f returns true for any value.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// StringInSlice returns whether a is in list.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IntInSlice returns whether a is in list.
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// FilterBSR filters []BookSearchResults if function returns true for any value.
func FilterBSR(vs []models.BookSearchResult, f func(models.BookSearchResult) bool) []models.BookSearchResult {
	vsf := make([]models.BookSearchResult, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// IndexBSR returns index of BSR in a BSR slice, otherwise -1.
func IndexBSR(vs []models.BookSearchResult, t models.BookSearchResult) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// RemoveBSRDuplicates removes duplicates from []BookSearchResult and returns the new array.
func RemoveBSRDuplicates(elements []models.BookSearchResult) []models.BookSearchResult {
	encountered := map[models.BookSearchResult]bool{}
	result := []models.BookSearchResult{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}

	return result
}
