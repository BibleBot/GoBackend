package textpurification

import "strings"

// PurifyVerseText removes impurities from verse texts.
func PurifyVerseText(text string) string {
	nuisances := map[string]string{
		"     ": " ",
		"    ":  " ",
		"  ":    " ",
		"“":     "\"",
		"”":     "\"",
		"\n":    " ",
		"¶ ":    "",
		" , ":   ", ",
		" .":    ".",
		"′":     "'",
	}

	for nuisance, replacement := range nuisances {
		if strings.Contains(text, nuisance) {
			text = strings.ReplaceAll(text, nuisance, replacement)
		}
	}

	if strings.Contains(text, "Selah") {
		text = strings.ReplaceAll(text, "Selah", " *(Selah)* ")
	}

	return strings.Join(strings.Fields(strings.TrimSpace(text)), " ")
}
