package search_suggestions

import (
	"sort"
	"strings"
)



/*
 * Complete the 'searchSuggestions' function below.
 *
 * The function is expected to return a 2D_STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY repository
 *  2. STRING customerQuery
 */

func searchSuggestions(repository []string, customerQuery string) [][]string {
	result := [][]string{}

	if strings.Trim(customerQuery, " ") == "" || len(repository) == 0 {
		return [][]string{}
	}

	sort.Strings(repository)
	customerQuery = strings.ToLower(customerQuery)

	for i := 1; i < len(customerQuery); i ++ {
		chars := customerQuery[0:i+1]
		var values [] string
		for _, value := range repository {
			value = strings.ToLower(value)
			if strings.Contains(value, chars) {
				values = append(values, value)
			}
			if len(values) == 3 {
				break
			}
		}

		result = append(result, values)
	}


	return result
}