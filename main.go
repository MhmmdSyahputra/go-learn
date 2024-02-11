package main
import "fmt"

func main() {
	// variabel and tipe data =============================================
    fmt.Println("\n=== VARIABEL AND TIPE DATA ===")
    var nama string = "putra"
    umur := 21
    fmt.Println(nama, "berusia", umur, "tahun.")
    fmt.Println("\n=============\n")
    // var nama string (var) = "putra" -> this global scope
    // umur := 21 (:=) -> this local scope
	// ===================================================================
	
	// LIST =============================================
    fmt.Println("\n=== LIST APPEND ===")
    // Membuat slice kosong
    var angka []int
    // Menambahkan data ke dalam slice menggunakan fungsi append
    angka = append(angka, 10)
    angka = append(angka, 20)
    angka = append(angka, 30)
	
    // Menampilkan isi slice
    fmt.Println("Isi slice angka:", angka)
    fmt.Println("\n=============\n")
	// ===================================================================
	
	// INPUT DATA =============================================
    fmt.Println("\n=== INPUT ===")
    var bilangan1 int
    var bilangan2 int
    // Meminta input dari pengguna
    fmt.Print("masukan angka pertama: ")
    fmt.Scanln(&bilangan1)

    fmt.Print("masukan angka kedua: ")
    fmt.Scanln(&bilangan2)

	//method template literation
	fmt.Printf("%d * %d = %d\n", bilangan1, bilangan2, bilangan1*bilangan2)
	// ===================================================================
	
	// FOR ===============================================================
    fmt.Println("\n=== FOR ===")
	// FOR
    for i := 0; i < 5; i++ {
		fmt.Println(i)
    }
	fmt.Println("\n")
	
    // LIKE WHILE
    j := 0
    for j < 5 {
		fmt.Println(j)
        j++
    }
	fmt.Println("\n")
	
    // INFINITE LOOP
    k := 0
    for {
		fmt.Println(k)
        k++
        if k == 5 {
			break
        }
    }
	fmt.Println("\n")
	
	// ===================================================================

	// IF ===============================================================
	fmt.Println("\n=== IF ===")
	// Pengkondisian if-else-if
	num := 10
	//IF
    if num > 10 {
        fmt.Println("Num lebih besar dari 10")
    } else if num == 10 {
        fmt.Println("Num sama dengan 10")
    } else {
        fmt.Println("Num kurang dari 10")
    }
	
	// Pengkondisian switch
    switch num {
    case 1:
        fmt.Println("Satu")
    case 2:
        fmt.Println("Dua")
    case 3:
        fmt.Println("Tiga")
    default:
        fmt.Println("Lebih dari tiga")
    }

	// Penggunaan switch tanpa ekspresi
    switch {
    case num < 5:
        fmt.Println("Kurang dari lima")
    case num == 5:
        fmt.Println("Sama dengan lima")
    case num > 5:
        fmt.Println("Lebih dari lima")
    }

	// ===================================================================
}
