package username

import (
	"sort"
	"strings"
)

func possibleChange(usernames []string) []string {

	results := []string{}

	for _, username := range usernames {
		chars := strings.Split(username, "")
		sort.Strings(chars)
		rearranged := strings.Join(chars, "")

		if rearranged != username {
			results = append(results, "YES")
		} else {
			results = append(results, "NO")
		}
	}

	return results
}
