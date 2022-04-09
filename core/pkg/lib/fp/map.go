package fp

func Map[A any, B any](f func(A) B, list []A) []B {
	if f == nil {
		return []B{}
	}
	newList := make([]B, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
