package files

import "strings"

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	// Check if found as lower or title case:
	found = strings.Contains(line, strings.ToLower(old)) || strings.Contains(line, strings.Title(old))
	if !found {
		return found, line, 0
	}

	// Lower case...
	occ = strings.Count(line, strings.ToLower(old))
	res = strings.ReplaceAll(
		line, strings.ToLower(old), strings.ToLower(new),
	)

	// Title case...
	occ += strings.Count(line, strings.Title(old))
	res = strings.ReplaceAll(
		res, strings.Title(old), strings.Title(new),
	)
	return found, res, occ
}
