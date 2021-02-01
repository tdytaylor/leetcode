package issue_0888

func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	setA := map[int]struct{}{}
	for _, v := range A {
		sumA += v
		setA[v] = struct{}{}
	}

	sumB := 0
	for _, v := range B {
		sumB += v
	}

	delta := (sumA - sumB) >> 1
	for k := range B {
		b := B[k]
		a := b + delta
		if _, has := setA[a]; has {
			return []int{a, b}
		}
	}
	return nil
}

func fairCandySwap2(a, b []int) []int {
	sumA := 0
	setA := map[int]struct{}{}
	for _, v := range a {
		sumA += v
		setA[v] = struct{}{}
	}
	sumB := 0
	for _, v := range b {
		sumB += v
	}
	delta := (sumA - sumB) >> 1
	for i := 0; ; i++ {
		y := b[i]
		x := y + delta
		if _, has := setA[x]; has {
			return []int{x, y}
		}
	}
}
