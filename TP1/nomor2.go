/* NIM: 1301194182
   Nama: Manuel Benedict
   Program ini melakukan untuk mengecek bilangan bulat yang diinput oleh user apakah termasuk tahun kabisat atau tidak.
*/

package main
import "fmt"

func main () {
	var tahun int
	var pengecekan bool
	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)
	pengecekan=tahun%400==0 ||tahun%4==0&&tahun%100!=0
	fmt.Print("Tahun kabisat: ")
	fmt.Println(pengecekan)
	
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}
	
	