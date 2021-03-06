package verses

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/models"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/bookmap"
	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/namefetcher"
)

// FindBooksInString locates a book name within a string, accounting for other parameters.
func FindBooksInString(str string) (string, []models.BookSearchResult) {
	books := namefetcher.GetBookNames()
	defaultBooks := namefetcher.GetDefaultBookNames()

	var results []models.BookSearchResult

	for bookKey, valueArray := range books {
		for _, item := range valueArray {
			if isValueInString(strings.ToLower(item), str) {
				str = strings.Replace(str, strings.ToLower(item), bookKey, -1)
			}
		}
	}

	tokens := strings.Split(str, " ")
	for _, book := range defaultBooks {
		for idx, token := range tokens {
			if token == book {
				results = append(results, models.BookSearchResult{
					Name:  book,
					Index: idx,
				})
			}
		}
	}

	// Sort output to match input order.
	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Index < results[j].Index
	})

	return str, results
}

// GenerateReference creates a reference object based on a BookSearchResult and the surrounding values in a string.
func GenerateReference(db gorm.DB, str string, bookSearchResult models.BookSearchResult, version models.Version) *models.Reference {
	book := bookSearchResult.Name
	startingChapter := 0
	startingVerse := 0
	endingChapter := 0
	endingVerse := 0
	tokenIdxAfterSpan := 0

	tokens := strings.Split(str, " ")

	if bookSearchResult.Index+2 <= len(tokens) {
		relevantToken := tokens[bookSearchResult.Index+1:][0]

		if strings.Contains(relevantToken, ":") {
			tokenIdxAfterSpan = bookSearchResult.Index + 2

			colonRegex, _ := regexp.Compile(":")
			colonQuantity := len(colonRegex.FindAllStringIndex(relevantToken, -1))

			switch colonQuantity {
			case 2:
				span := strings.Split(relevantToken, "-")

				for _, pairString := range span {
					pair := strings.Split(pairString, ":")

					for idx, pairValue := range pair {
						pair[idx] = removePunctuation(pairValue)
					}

					firstNum, firstErr := strconv.Atoi(pair[0])
					secondNum, secondErr := strconv.Atoi(pair[1])

					if firstErr != nil || secondErr != nil {
						return nil
					}

					if startingChapter == 0 {
						startingChapter = firstNum
						startingVerse = secondNum
					} else {
						endingChapter = firstNum
						endingVerse = secondNum
					}
				}
			case 1:
				pair := strings.Split(relevantToken, ":")

				num, err := strconv.Atoi(pair[0])
				if err != nil {
					return nil
				}

				startingChapter = num
				endingChapter = num

				spanRegex, _ := regexp.Compile("-")
				spanQuantity := len(spanRegex.FindStringSubmatchIndex(relevantToken))

				span := strings.Split(pair[1], "-")
				for idx, pairValue := range span {
					pairValue = removePunctuation(pairValue)

					num, err := strconv.Atoi(pairValue)
					if err != nil {
						// We know that BibleGateway will extend to the end of a chapter with this syntax,
						// but for other sources this is likely not available.
						if version.Source == "bg" {
							// Instead of returning nil, we'll break out of the loop
							// in the event that the span exists to extend to the end of a chapter.
							break
						}
					}

					switch idx {
					case 0:
						startingVerse = num
					case 1:
						endingVerse = num
					default:
						return nil
					}
				}

				if endingVerse == 0 && spanQuantity == 0 {
					endingVerse = startingVerse
				}
			}

			// If the last token of the reference is a version abbreviation, utilize that version.
			if len(tokens) > tokenIdxAfterSpan {
				lastToken := strings.ToUpper(tokens[tokenIdxAfterSpan])

				var idealVersion models.Version
				db.Where(&models.Version{Abbreviation: lastToken}).First(&idealVersion)

				if idealVersion.Abbreviation == lastToken {
					version = idealVersion
				}
			}
		}
	} else {
		return nil
	}

	isOT := false
	isNT := false
	isDEU := false

	bookmapping := bookmap.GetBookmap()
	if correctBook, ok := bookmapping["ot"][book]; ok {
		isOT = true
		book = correctBook
	} else if correctBook, ok := bookmapping["nt"][book]; ok {
		isNT = true
		book = correctBook
	} else if correctBook, ok := bookmapping["deu"][book]; ok {
		isDEU = true
		book = correctBook
	}

	if startingChapter == 0 || startingVerse == 0 {
		return nil
	}

	return &models.Reference{
		Book:            book,
		StartingChapter: startingChapter,
		StartingVerse:   startingVerse,
		EndingChapter:   endingChapter,
		EndingVerse:     endingVerse,
		Version:         version,

		IsOT:  isOT,
		IsNT:  isNT,
		IsDEU: isDEU,
	}
}

// -- helper functions --

func isValueInString(value string, str string) bool {
	return strings.Contains(fmt.Sprintf(" %s ", str), fmt.Sprintf(" %s ", value))
}

func removePunctuation(str string) string {
	noPunctuationRegex, _ := regexp.Compile(`[^\w\s]|_`)
	minimizeWhitespaceRegex, _ := regexp.Compile(`\s+`)

	return minimizeWhitespaceRegex.ReplaceAllString(noPunctuationRegex.ReplaceAllString(str, ""), " ")
}
