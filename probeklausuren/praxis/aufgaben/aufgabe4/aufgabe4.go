package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

func MaxElements(l1, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	headMax := l1[0]
	if l2[0] > headMax {
		headMax = l2[0]
	}

	return append([]int{headMax}, MaxElements(l1[1:], l2[1:])...)
}
