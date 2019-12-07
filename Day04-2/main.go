package main

import (
	"strconv"
)

func main() {

}

func validatePassword(password string, lowerRange int, upperRange int) (bool, string) {
	if len(password) != 6 {
		return false, "length"
	}

	passwordNumber, _ := strconv.Atoi(password)

	if !(lowerRange <= passwordNumber && passwordNumber <= upperRange) {
		return false, "range"
	}

	if !(password[0] <= password[1] && password[1] <= password[2] && password[2] <= password[3] && password[3] <= password[4] && password[4] <= password[5]) {
		return false, "increasing"
	}

	// if !(password[0] == password[1] || password[1] == password[2] || password[2] == password[3] || password[3] == password[4] || password[4] == password[5]) {
	// 	return false, "adjacent "
	// }

	if !hasAtLeastOneGroupOfOnlyTwo(password) {
		return false, "more than 2 adjacent "
	}

	return true, ""
}

func hasAtLeastOneGroupOfOnlyTwo(password string) bool {
	//"111122"

	groups := []string{}
	var lastChecked byte = 0
	temp := ""
	for i := 0; i < 6; i++ {
		if lastChecked != password[i] {
			//start new group
			if temp != "" {
				groups = append(groups, temp)
				temp = ""
			}
			lastChecked = password[i]
		}
		temp += string(password[i])
	}
	if temp != "" {
		groups = append(groups, temp)
		temp = ""
	}

	for _, group := range groups {
		if len(group) == 2 {
			return true
		}
	}
	return false
}

func generatePasswords(lowerRange int, upperRange int) []string {
	passwords := []string{}
	for i := lowerRange; i <= upperRange; i++ {
		password := strconv.Itoa(i)

		isValid, _ := validatePassword(password, lowerRange, upperRange)
		if isValid {
			passwords = append(passwords, password)
		}
	}

	return passwords
}
