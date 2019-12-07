package main

import (
	"testing"
)

func TestPassword1(t *testing.T) {
	password := "122345"
	expected := true

	actual, message := validatePassword(password, 0, 999999)

	if actual != expected {
		t.Errorf("Password should be %t, but is %t with message '%s' '%s'", expected, actual, message, password)
	}
}

func TestPassword2(t *testing.T) {
	password := "111123"
	expected := false

	actual, message := validatePassword(password, 0, 999999)

	if actual != expected {
		t.Errorf("Password should be %t, but is %t with message '%s' '%s'", expected, actual, message, password)
	}
}

func TestPassword3(t *testing.T) {
	password := "223450"
	expected := false

	actual, message := validatePassword(password, 0, 999999)

	if actual != expected {
		t.Errorf("Password should be %t, but is %t with message '%s' '%s'", expected, actual, message, password)
	}
}

func TestPassword4(t *testing.T) {
	password := "123789"
	expected := false

	actual, message := validatePassword(password, 0, 999999)

	if actual != expected {
		t.Errorf("Password should be %t, but is %t with message '%s' '%s'", expected, actual, message, password)
	}
}

func TestPassword5(t *testing.T) {
	password := "112233"
	expected := true

	actual, message := validatePassword(password, 0, 999999)

	if actual != expected {
		t.Errorf("Password should be %t, but is %t with message '%s' '%s'", expected, actual, message, password)
	}
}

func TestPassword6(t *testing.T) {
	password := "123444"
	expected := false

	actual, message := validatePassword(password, 0, 999999)

	if actual != expected {
		t.Errorf("Password should be %t, but is %t with message '%s' '%s'", expected, actual, message, password)
	}
}

func TestPassword7(t *testing.T) {
	password := "111122"
	expected := true

	actual, message := validatePassword(password, 0, 999999)

	if actual != expected {
		t.Errorf("Password should be %t, but is %t with message '%s' '%s'", expected, actual, message, password)
	}
}

func TestPasswordGen1(t *testing.T) {
	lower := 112229
	upper := 112234
	expected := 3

	actual := generatePasswords(lower, upper)

	if len(actual) != expected {
		t.Errorf("There should be %d passwords, but is %d", expected, len(actual))
	}
}

func TestPasswordGen2(t *testing.T) {
	lower := 273025
	upper := 280000
	expected := 7

	actual := generatePasswords(lower, upper)

	if len(actual) != expected {
		t.Errorf("There should be %d passwords, but is %d", expected, len(actual))
	}
}

func TestMoreThanTwoInARow(t *testing.T) {
	password := "112233"
	if hasAtLeastOneGroupOfOnlyTwo(password) == false {
		t.Errorf("Should have passed %s", password)
	}
	password = "123444"
	if hasAtLeastOneGroupOfOnlyTwo(password) == true {
		t.Errorf("Should have failed %s", password)
	}
	password = "111122"
	if hasAtLeastOneGroupOfOnlyTwo(password) == false {
		t.Errorf("Should have passed %s", password)
	}
}

func TestFinal(t *testing.T) {
	lower := 273025
	upper := 767253
	expected := 598

	actual := generatePasswords(lower, upper)

	if len(actual) != expected {
		t.Errorf("There should be %d passwords, but is %d", expected, len(actual))
	}
}
