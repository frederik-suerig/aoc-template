package helpers

func RemoveElement[T any](s []T, index int) []T {
	r := make([]T, 0, len(s)-1)
	r = append(r, s[:index]...)
	if index >= len(s)-1 {
		return r 
	}

	return append(r, s[index+1:]...)
}
