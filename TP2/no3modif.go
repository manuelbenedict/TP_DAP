/* NIM: 1301194182
   Nama: Manuel Benedict
   Bentuk Perulangan - Modifikasi
*/

package main
import "fmt"

func main () {
	var a,b float64
			fmt.Print("Masukkan berat belanjaan di kedua kantong (dalam kg): ")
			fmt.Scan(&a)
			fmt.Scan(&b)
			fmt.Println("Berat kedua kantong (dalam kg): ", a," ",b)
			fmt.Println("Sepeda motor pak Andi akan oleng: ", a-b>=9||b-a>=9)
	for (a+b<=150)&&(a>0&&b>0) {
			fmt.Print("Masukkan berat belanjaan di kedua kantong (dalam kg): ")
			fmt.Scan(&a)
			fmt.Scan(&b)
			fmt.Println("Berat kedua kantong (dalam kg): ", a," ",b)
			fmt.Println("Sepeda motor pak Andi akan oleng: ", a-b>=9||b-a>=9)
	}
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}