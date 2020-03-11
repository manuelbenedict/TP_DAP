/* Nama	: Manuel Benedict
   NIM	: 1301194182
   Program mencari titik terdekat dari antara banyak titik yang diinput user (sampai user memasukkan 0.0 0.0).
*/

package main
import "fmt"
import "math"

type Point struct {
	x float64
	y float64
}

const N=1000
var points [N]Point
var numpoints int=0

func jarak(p1,p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2)+math.Pow(p1.y-p2.y, 2))
}

func bacaTitik () {
	var a,b float64
	fmt.Scanln(&a,&b)
	for a!=0||b!=0 {
		points[numpoints]=Point {
			a,
			b,
		}
		fmt.Scanln(&a,&b)
		numpoints++
	}
}

func ambilTitikTerdekat(a,b *Point) {
	var a_min,b_min int
	var jarak_min float64
	jarak_min=jarak(points[0],points[1])
	for i:=0; i<numpoints; i++ {
		for j:=0; j<numpoints; j++ {
			if jarak(points[i], points[j])<jarak_min && i!=j {
				a_min=i
				b_min=j
				jarak_min=jarak(points[i], points[j])
			}
		}
	}
	*a=points[a_min]
	*b=points[b_min]
}

func main () {
	var a,b Point
	bacaTitik()
	ambilTitikTerdekat(&a,&b)
	fmt.Printf("Titik terdekat adalah (%v, %v) dan (%v, %v) dengan jarak %v. \n", a.x, a.y, b.x, b.y, jarak(a,b))
	fmt.Println("Nama: Manuel Benedict")
	fmt.Println("NIM: 1301194182")
}