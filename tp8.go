/* Nama		: Manuel Benedict
   NIM		: 1301194182
   Program yang menampilkan nilai terbesar dan indeks data terkecil dari sekumpulan data yang terdiri dari tiga huruf, bilangan bulat positif (sebagai indeks), dan bilangan riil
*/

package main
import "fmt"

const N=2019
type RecType struct {
 f1 string
 f2 int
 f3 float64
}
type ArrType[N] RecType
var data ArrType
var i int

func tambahArr() {
	var c1, c2, c3 int
	c1, c2, c3 = 97, 97, 97
	for i = 0; i < len(data); i++ {
		c3++
		if c3 > 122 {
			c3 = 97
			c2++
		} else if c2 > 122 {
			c2 = 97
			c1++
		}
		a := string(c1) + string(c2) + string(c3)
		data[i] = RecType{f1: a, f2: i, f3: float64(i) * 3.1}
	}
}

func rmax(data ArrType) float64{
	var maks float64
	maks= 0.0
	i=0
	for i < N {
		if data[i].f3 > maks {
			maks = data[i].f3
		}
		i++
	}
	return maks
}

func imin(data ArrType) int {
	var indeks,min int
	indeks= 0
	min= 9999999
	i=0
	for i < N {
		if data[i].f2 < min {
			indeks = i
			min = data[i].f2
		}
		i++
	}
	return indeks
}

func found(data ArrType, key string) bool {
	i= 0
	for i < N {
		if data[i].f1 == key {
			return true
		}
	}
	return false
}

func pos(data ArrType, key string) int {
	var k,k1,median int
	k= 0 
	k1= len(data)-1
	median=0
	for k <= k1 {
		median = (k1+k)/2
		if data[median].f1 != key {
			k++
		} else {
			k1--
		}
	}
	return median
}

func main(){
	tambahArr()
	fmt.Println("====================================")
	a:=0
	for a < N {
		if data[a].f2 > 0 {
		fmt.Println(data[a])
		}
		a++
	}
	fmt.Println("====================================")
	fmt.Println("Bilangan riil terbesar : ",rmax(data))
	fmt.Println("Posisi indeks bilangan terkecil : ",imin(data))
	fmt.Println("====================================")
	fmt.Println("Nama : Manuel Benedict")
	fmt.Println("NIM : 1301194182")
}