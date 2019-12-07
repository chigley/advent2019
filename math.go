package advent2019

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Permutations(a []int) <-chan []int {
	if len(a) == 0 {
		return nil
	}

	c := make(chan []int)
	go func() {
		permutations(len(a), append([]int(nil), a...), c)
		close(c)
	}()
	return c
}

func permutations(k int, a []int, c chan<- []int) {
	if k == 1 {
		c <- append([]int(nil), a...)
		return
	}

	permutations(k-1, a, c)
	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}
		permutations(k-1, a, c)
	}
}
