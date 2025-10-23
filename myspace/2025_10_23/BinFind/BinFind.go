package searching

func binFind( l[]int, x int) int {
	mitte := len(l) / 2
	if len(l) == 0 {
		return -1
	}
	if x == l[mitte]{
		return mitte
	}
	if x < l[mitte] {
		return binFind( l[:mitte], x)
	}
	return binFind( l[mitte+1:], x) + mitte + 1	

}