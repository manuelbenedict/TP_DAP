/* NIM : 1301194182
   Nama: Manuel Benedict
   Soal nomor 4
*/

package main
import "fmt"


var N,suku int
var total,penyebut float64

func main () {
	fmt.Print("Masukkan N suku pertama: ")
	fmt.Scan(&N)
	suku=1
	penyebut=1
	for suku<=N {
		if suku%2==0 {
			total=total-4/penyebut
		} else if suku%2==1 {
			total=total+4/penyebut
		}
		penyebut=penyebut+2
		suku++
	}
	fmt.Println("Nilai pi: ", total)
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}