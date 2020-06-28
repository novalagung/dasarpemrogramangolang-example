package persegi

import "fmt"

//exported function karena di awali huruf besar
func KelilingPersegi(sisi int) {
	kililing := sisi * 4
	fmt.Printf("Keliling = %d \n", kililing)
	printEnd()
}

//unexported fucntion karena diawali huruf kecil
func printEnd() {
	fmt.Println("End")
}
