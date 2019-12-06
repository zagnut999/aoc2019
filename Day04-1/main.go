package main

import "strconv"

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

	if !(password[0] == password[1] || password[1] == password[2] || password[2] == password[3] || password[3] == password[4] || password[4] == password[5]) {
		return false, "adjacent "
	}

	if !(password[0] <= password[1] && password[1] <= password[2] && password[2] <= password[3] && password[3] <= password[4] && password[4] <= password[5]) {
		return false, "increasing"
	}

	return true, ""
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
