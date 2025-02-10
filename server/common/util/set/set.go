package set

type Set[T comparable] interface {
	Has(value T) bool
	Add(value T)
	Remove(value T)
}

func Has(array []string, value string) bool {
	for _, id := range array {
		if id == value {
			return true
		}
	}

	return false
}

func Add(array []string, value string) []string {
	has_value := Has(array, value)

	if has_value {
		return array
	}

	return append(array, value)
}

func Remove(array []string, value string) []string {
	has_value := Has(array, value)

	if !has_value {
		return array
	}

	i := 0

	for _, id := range array {
		if id != value {
			array[i] = id

			i++
		}
	}

	return array[:i]
}
