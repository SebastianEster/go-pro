package main

import "testing"

func TestNewRational(t *testing.T) {
	simpleRational := newRational(1, 2)
	if simpleRational.numerator != 1 {
		t.Error("Numerator should be 1")
	}
	if simpleRational.denominator != 2 {
		t.Error("Denominator should be 2")
	}

	complexRational := newRational(15, 12)
	if complexRational.numerator != 5 {
		t.Error("Numerator should be 5")
	}
	if complexRational.denominator != 4 {
		t.Error("Denominator should be 4")
	}
}

func TestAdd(t *testing.T) {
	firstRational := newRational(1, 2)
	secondRational := newRational(2, 3)

	result := firstRational.add(secondRational)

	if result.numerator != 7 {
		t.Error("Numerator should be 7")
	}
	if result.denominator != 6 {
		t.Error("Denominator should be 6")
	}
}

func TestMultiply(t *testing.T) {
	firstRational := newRational(1, 2)
	secondRational := newRational(2, 3)

	result := firstRational.multiply(secondRational)

	if result.numerator != 1 {
		t.Error("Numerator should be 1")
	}
	if result.denominator != 3 {
		t.Error("Denominator should be 3")
	}
}
