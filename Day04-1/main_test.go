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
	expected := true

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

func TestPasswordGen1(t *testing.T) {
	lower := 111111
	upper := 111112
	expected := 2

	actual := generatePasswords(lower, upper)

	if len(actual) != expected {
		t.Errorf("There should be %d passwords, but is %d", expected, len(actual))
	}
}

func TestPasswordGen2(t *testing.T) {
	lower := 111131
	upper := 111135
	expected := 3

	actual := generatePasswords(lower, upper)

	if len(actual) != expected {
		t.Errorf("There should be %d passwords, but is %d", expected, len(actual))
	}
}

func TestFinal(t *testing.T) {
	lower := 273025
	upper := 767253
	expected := 3

	actual := generatePasswords(lower, upper)

	if len(actual) != expected {
		t.Errorf("There should be %d passwords, but is %d", expected, len(actual))
	}
}
