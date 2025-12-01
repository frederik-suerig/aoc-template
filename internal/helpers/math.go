package helpers

// SignedInteger is a constraint for signed integer types.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Abs returns the absolute value of a signed integer.
func Abs[T SignedInteger](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// IntAbs returns the absolute value of an integer (convenience function for int).
func IntAbs(x int) int {
	return Abs(x)
}

// Min returns the minimum of two comparable values.
func Min[T SignedInteger](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two comparable values.
func Max[T SignedInteger](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// MinN returns the minimum of a variable number of values.
func MinN[T SignedInteger](nums ...T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

// MaxN returns the maximum of a variable number of values.
func MaxN[T SignedInteger](nums ...T) T {
	if len(nums) == 0 {
		var zero T
		return zero
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

// Mod returns the proper modulo operation (handles negative numbers correctly).
// Unlike Go's % operator, Mod always returns a non-negative result.
func Mod[T SignedInteger](a, b T) T {
	result := a % b
	if result < 0 {
		result += b
	}
	return result
}

// Pow returns base raised to the power of exp (base^exp).
// Returns 0 if exp is negative.
func Pow[T SignedInteger](base, exp T) T {
	if exp < 0 {
		var zero T
		return zero
	}
	if exp == 0 {
		return 1
	}
	result := T(1)
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
