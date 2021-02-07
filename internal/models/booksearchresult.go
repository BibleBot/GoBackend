package models

// BookSearchResult stores the name of the book and the index of the last word in the tokenized string.
type BookSearchResult struct {
	Name  string
	Index int
}
