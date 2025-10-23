package searching

func BinFind( l[]int, x int) int {
	mitte := len(l) / 2
	if len(l) == 0 {
		return -1
	}
	if x == l[mitte]{
		return mitte
	}
	if x < l[mitte] {
		return BinFind( l[:mitte], x)
	}
	return BinFind( l[mitte+1:], x) + mitte + 1

}