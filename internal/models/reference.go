package models

import "fmt"

// Reference is a type for a verse reference, formulated after parsing.
type Reference struct {
	Book            string
	StartingChapter int
	StartingVerse   int
	EndingChapter   int
	EndingVerse     int
	Version         Version

	IsOT  bool
	IsNT  bool
	IsDEU bool
}

// ToString returns the reference in a string form.
func (ref Reference) ToString() string {
	result := fmt.Sprintf("%s %d:%d", ref.Book, ref.StartingChapter, ref.StartingVerse)

	if ref.EndingChapter > 0 && ref.EndingChapter != ref.StartingChapter {
		result += fmt.Sprintf("-%d:%d", ref.EndingChapter, ref.EndingVerse)
	} else if ref.EndingVerse > 0 && ref.EndingVerse != ref.StartingVerse {
		result += fmt.Sprintf("-%d", ref.EndingVerse)
	} else if ref.EndingChapter > 0 && ref.EndingVerse == 0 {
		result += "-"
	}

	return result
}
