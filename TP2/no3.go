/* NIM: 1301194182
   Nama: Manuel Benedict
   Bentuk Perulangan
*/

package main
import "fmt"

func main () {
	var a,b float64
	for a<9&&b<9 {
		fmt.Print("Masukkan berat belanjaan di kedua kantong (dalam kg): ")
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Println("Berat kedua kantong (dalam kg): ", a," ",b)
	}
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}
		