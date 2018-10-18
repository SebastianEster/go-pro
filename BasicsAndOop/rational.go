package main

type Rational struct {
	numerator int
	denominator int
}

func newRational(num, den int) Rational {
	if den == 0 {
		panic("Division by zero")
	}

	gcd := func (a, b int) int {
		for b > 0 {
		a, b = b, a % b
	}
		return a
	}

	divisor := gcd(num, den)
	return Rational{num / divisor, den / divisor}
}

func (rational Rational) add(otherRational Rational) Rational {
	leftNumerator := rational.numerator * otherRational.denominator
	rightNumerator := otherRational.numerator * rational.denominator
	commonDenominator := rational.denominator* otherRational.denominator
	return newRational(leftNumerator + rightNumerator, commonDenominator)
}

func (rational Rational) multiply(otherRational Rational) Rational {
	return newRational(rational.numerator * otherRational.numerator,
		rational.denominator * otherRational.denominator)
}