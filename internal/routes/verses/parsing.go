package verses

import (
	"fmt"
	"sort"
	"strings"

	"github.com/BibleBot/backend/internal/models"
	"github.com/BibleBot/backend/internal/utils/namefetcher"
	"github.com/BibleBot/backend/internal/utils/slices"
)

// FindBooksInString locates a book name within a string, accounting for other parameters.
func FindBooksInString(str string) (string, []models.BookSearchResult) {
	str = strings.Replace(str, "Letter of Jeremiah", "epjer", 1)

	tokens := strings.Split(str, " ")
	books := namefetcher.GetBookNames(false)

	var results []models.BookSearchResult
	var existingIndices []int

	for bookKey, valueArray := range books {
		for _, item := range valueArray {
			potentialValues := []string{strings.ToUpper(item), strings.ToLower(item), strings.ToTitle(item), item}

			for _, value := range potentialValues {
				for range tokens {
					if isTokenInString(value, str) {
						fmt.Printf("'%s' contains '%s'\n", str, value)
						overlappedNames := map[string][]string{
							"john": append(books["1john"], append(books["2john"], books["3john"]...)...),
							"ezra": append(books["1esd"], books["2esd"]...),
							"song": books["song"],
							"ps":   books["ps151"],
						}

						lastItem := strings.Split(value, " ")[len(strings.Split(value, " "))-1]
						potentialLastItems := []string{strings.ToUpper(lastItem), strings.ToLower(lastItem), strings.ToTitle(lastItem), lastItem}
						potentialIndices := slices.MapSlice(tokens, func(val string, mapIdx int) int {
							if slices.Index(potentialLastItems, val) != -1 {
								return mapIdx
							}

							return -1
						})

						_, potentialOverlapKey := overlappedNames[bookKey]
						isOverlappingBook := slices.Any(overlappedNames[bookKey], func(val string) bool {
							return isTokenInString(val, str)
						})

						if potentialOverlapKey && isOverlappingBook {
							for _, idx := range potentialIndices {
								if idx == -1 {
									continue
								}

								var bookName string

								switch bookKey {
								case "ps":
									bookName = fmt.Sprintf("%s %s", tokens[idx], tokens[idx+1])
									break
								case "song":
									bookName = fmt.Sprintf("%s %s %s", tokens[idx-2], tokens[idx-1], tokens[idx])
									break
								default:
									bookName = fmt.Sprintf("%s %s", tokens[idx-1], tokens[idx])
									break
								}

								if slices.StringInSlice(bookName, overlappedNames[bookKey]) {
									for key, value := range books {
										if slices.StringInSlice(bookName, value) {
											str = strings.Replace(str, bookName, key, 1)
											tempTokens := strings.Split(str, " ")

											results = append(results, models.BookSearchResult{
												Name:  key,
												Index: slices.Index(tempTokens, key),
											})

											existingIndices = append(existingIndices, idx)
											str = strings.Replace(str, key, "nil", 1)
										}
									}

								}
							}
						} else {
							for _, idx := range potentialIndices {
								if idx == -1 {
									continue
								}

								if !slices.IntInSlice(idx, existingIndices) {
									str = strings.Replace(str, value, bookKey, 1)
									tempTokens := strings.Split(str, " ")

									results = append(results, models.BookSearchResult{
										Name:  bookKey,
										Index: slices.Index(tempTokens, bookKey),
									})

									existingIndices = append(existingIndices, idx)
									str = strings.Replace(str, bookKey, "nil", 1)
									break
								}
							}
						}
					}
				}
			}
		}
	}

	tempTokens := strings.Split(str, " ")
	for idx, token := range tempTokens {
		if token == "epnil" {
			tempTokens[slices.Index(tempTokens, "epnil")] = "epjer"

			results = append(results, models.BookSearchResult{
				Name:  "epjer",
				Index: idx,
			})
		} else if token == "epjer" {
			results = append(results, models.BookSearchResult{
				Name:  "epjer",
				Index: idx,
			})
		}

		for _, result := range results {
			if idx == result.Index && token == "nil" {
				tempTokens[slices.Index(tempTokens, "nil")] = result.Name
			}
		}
	}
	str = strings.Join(tempTokens, " ")

	// Remove duplicates and invalid results.
	filteredResults := slices.RemoveBSRDuplicates(slices.FilterBSR(results, func(bsr models.BookSearchResult) bool {
		return bsr.Index > -1
	}))

	// Sort output to match input order.
	sort.SliceStable(filteredResults, func(i, j int) bool {
		return filteredResults[i].Index < filteredResults[j].Index
	})

	return str, filteredResults
}

func isTokenInString(token string, str string) bool {
	return strings.Contains(fmt.Sprintf(" %s ", str), fmt.Sprintf(" %s ", token))
}
