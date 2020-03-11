/* NIM : 1301194182
   Nama: Manuel Benedict
   Program fungsi komposisi
*/ 

package main
import "fmt"

var x float32

func f(x float32) float32 {
	f:=x*x
	return f
}

func g(x float32) float32 {
	g:=x-2
	return g
}

func h(x float32) float32 {
	h:=x+1
	return h
}

func fgh(x float32) float32 {
	fgh:=f(g((h(x))))
	return fgh
}

func main () {
	fmt.Print("Masukkan nilai x = ")
	fmt.Scan(&x)
	fmt.Println("f(", x,") =", f(x))
	fmt.Println("g(", x,") =", g(x))
	fmt.Println("h(", x,") =", h(x))
	fmt.Println("fogoh(", x,") =", fgh(x))
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}