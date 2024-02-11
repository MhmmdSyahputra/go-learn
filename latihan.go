package main
import "fmt"

func main(){
	fmt.Println("\n=== INPUT ===")
	var bilangan1 int
    var bilangan2 int
    // Meminta input dari pengguna
    fmt.Print("masukan angka pertama: ")
    fmt.Scanln(&bilangan1)

    fmt.Print("masukan angka kedua: ")
    fmt.Scanln(&bilangan2)

	//method template literation
	// fmt.Println(`%d * %d`,bilangan1,bilangan2)
}