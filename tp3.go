/* NIM: 1301194182
   Nama: Manuel Benedict
   Soal nomor 2c modul hal.17
*/

package main 
import "fmt"

func main () {
	var nam float64
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam>80 {
		fmt.Println("Nilai akhir mata kuliah: A")
	} else if nam>72.5 {
		fmt.Println("Nilai akhir mata kuliah: AB")
	} else if nam>65 {
		fmt.Println("Nilai akhir mata kuliah: B")
	} else if nam>57.5 {
		fmt.Println("Nilai akhir mata kuliah: BC")
	} else if nam>50 {
		fmt.Println("Nilai akhir mata kuliah: C")
	} else if nam>40 {
		fmt.Println("Nilai akhir mata kuliah: D")
	} else if nam<=0 {
		fmt.Println("Nilai akhir mata kuliah: E")
	}
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}