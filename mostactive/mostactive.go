package mostactive

import "sort"

func mostActive(customers []string) []string {
	customerCount := make(map[string]int)

	for _, customer := range customers {
		customerCount[customer]++
	}

	var result []string

	for customer, count := range customerCount {
		totalOrders := len(customers)
		activityPercentage := float64(count) / float64(totalOrders) * 100

		if activityPercentage >= 5 {
			result = append(result, customer)
		}
	}

	sort.Strings(result)
	return result
}
