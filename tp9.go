/* Nama	: Manuel Benedict
   NIM	: 1301194182
  Program pencarian dan pengurutan array 
*/

package main
import "fmt"

const MAXSIZE=20192020
type RecType struct {
	count,size int
}
type ArrType[MAXSIZE]RecType
var arr ArrType
var i,j int

func iSort(arr ArrType, nsize int) {
	var tampung int
	for i=0; i<nsize; i++ {
		for j=0; j<i+1; j++ {
			if arr[i].count<arr[j].count {
				tampung=arr[j].count
				arr[j].count=arr[i].count
				arr[i].count=tampung
			}
		}
	}
	for i=0; i<nsize; i++ {
		fmt.Print(arr[i].count, " ")
	}
}

func mSort(arr *ArrType, nsize int) {
	var maks,nMaks int
	for i=0; i<nsize; i++ {
		maks=i
		for j=i+1; j<nsize; j++ {
			if arr[j].size>arr[maks].size {
				maks=j
			}
		}
		nMaks=arr[maks].size
		arr[maks].size=arr[i].size
		arr[i].size=nMaks
	}
}

func isFound(arr ArrType, n,v int) bool {
	var kiri,tengah,kanan int
	var found bool
	kiri=0
	tengah=0
	kanan=n
	found=false
	for kiri<kanan&&!found {
		tengah=(kiri+kanan)/2
		if arr[tengah].count<v {
			kiri=tengah+1
		} else if arr[tengah].count>v {
			kanan=tengah
		} else {
			found=true
		}
	}
	return found
}

func main () {
	var cari int
	for i:=0; i<MAXSIZE; i++ {
		fmt.Scan(&arr[i].size)
	}
	for i:=0; i<MAXSIZE; i++ {
		fmt.Scan(&arr[i].count)
	}
	iSort(arr, MAXSIZE)
	mSort(&arr, MAXSIZE)
	fmt.Println()
	for i=0; i<MAXSIZE; i++ {
		fmt.Print(arr[i].size, " ")
	}
	fmt.Println()
	fmt.Print("Apa angka yang ingin Anda cari? ")
	fmt.Scan(&cari)
	fmt.Println(isFound(arr, MAXSIZE, cari))
	fmt.Println("==============================================")
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}