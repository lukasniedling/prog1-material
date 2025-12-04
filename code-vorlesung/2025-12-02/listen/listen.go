package listen

// Reverse erwartet eine Liste und dreht
// die Reihenfolge der Elemente um.
// Liefert die umgedrehte Liste zurück.
// Die Funktion muss rekursiv sein.
func Reverse(list []int) []int {
	if len(list) == 0 {
		return []int{}
	}
	return append([]int{list[len(list)-1]}, Reverse(list[:len(list)-1])...)
}

// DuplicateElements erwartet eine Liste und dupliziert
// jedes Element darin.
// Liefert die Ergebnisliste zurück.
// Die Funktion muss rekursiv sein.
func DuplicateElements(list []int) []int {
	// TODO
	return []int{}
}
