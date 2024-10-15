package services

import "strconv"

func Is_valid(cc_number string) bool {
	//odd numbers sum
	odd_sum := 0

	for i := 1; i <= 15; i += 2 {
		value, _ := strconv.ParseInt(string(cc_number[i]), 10, 0)
		odd_sum += int(value)
	}

	//even numbers sum
	even_sum := 0

	for i := 0; i <= 14; i += 2 {
		value, _ := strconv.ParseInt(string(cc_number[i]), 10, 0)

		double := int(value) * 2

		if double > 9 {
			double -= 9
		}

		even_sum += double
	}

	//return result
	return (odd_sum+even_sum)%10 == 0
}
