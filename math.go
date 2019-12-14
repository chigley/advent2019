package advent2019

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(x, y int) int {
	for y != 0 {
		tmp := y
		y = x % y
		x = tmp
	}
	return x
}

func LCM(x, y int, zs ...int) int {
	ret := x * y / GCD(x, y)
	for _, z := range zs {
		ret = LCM(ret, z)
	}
	return ret
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
