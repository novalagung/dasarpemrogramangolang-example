package main

func KategorikanNilai(nilai int) string {
	switch {
	case nilai >= 90:
		return "A"
	case nilai >= 80:
		return "B"
	case nilai >= 70:
		return "C"
	case nilai >= 60:
		return "D"
	default:
		return "E"
	}
}
