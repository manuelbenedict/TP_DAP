/* NIM : 1301194182
   Nama: Manuel Benedict
   Soal nomor 4 modifikasi
*/

package main
import "fmt"

var suku int
var Si,Si1,penyebut float64

func main () {
	suku=2
	penyebut=1
	for Si-Si1>0.00001||Si-Si1<(-0.00001)||Si-Si1==0 {
		if suku%2==1 {
			Si=Si-4/penyebut
			Si1=Si+4/(penyebut+2)
		} else if suku%2==0 {
			Si=Si+4/penyebut
			Si1=Si-4/(penyebut+2)
		}
		suku=suku+1
		penyebut=penyebut+2
	}
	fmt.Println("Hasil pi: ", Si)
	fmt.Println("Hasil pi: ", Si1)
	fmt.Println("Pada suku ke: ", suku)
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}