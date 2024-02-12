package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var pesanan []string
var hargaPesanan []int

func main() {
	for {
		fmt.Println("\n=== Warung Tok Abah ====")
		fmt.Println("1. Pesan")
		fmt.Println("2. Lihat Pesanan Saya")
		fmt.Println("3. Bayar")
		fmt.Println("4. Pulang")

		var pilihan int
		fmt.Print("PILIH : ")
		fmt.Scanln(&pilihan)
		if pilihan == 1 {
			clearScreen()
			pesanMenu()
		} else if pilihan == 2 {
			clearScreen()
			lihatPesanan()
		} else if pilihan == 3 {
			clearScreen()
			bayarPesanan()
		} else if pilihan == 4 {
			clearScreen()
			if len(pesanan) > 0 {
				fmt.Print("Harap selesainkan pembayaran anda terlebih dahulu")
			} else {
				fmt.Print("Terima Kasih Telah Berkunjung 😊😊")
				break
			}
		} else {
			clearScreen()
			fmt.Print("Menu Tidak Tersedia!")
		}
	}
}

func pesanMenu() {
	menu := []string{"Jus Apel", "jus Jeruk", "Sanger", "Teh Tarik", "Indomie", "Nasi Goreng"}
	harga := []int{10000, 10000, 20000, 12000, 13000, 15000}
	fmt.Println("\n==== M E N U ====")
	for i := 0; i < len(menu); i++ {
		fmt.Println(i+1, ".", string(menu[i]))
	}
	for {
		var menuSelected int
		fmt.Print("\nPilih Menu(1-6 | 0 Stop) : ")
		fmt.Scanln(&menuSelected)
		if menuSelected <= len(menu) && menuSelected > 0 {
			for i := 0; i < len(menu); i++ {
				if menuSelected == i+1 {
					pesanan = append(pesanan, menu[i])
					hargaPesanan = append(hargaPesanan, harga[i])
				}
			}
		} else if menuSelected == 0 {
			clearScreen()
			break
		} else {
			fmt.Println("Pilihan Tidak Tersedia")
		}

	}
}

func lihatPesanan() {
	if len(pesanan) <= 0 {
		fmt.Println("Anda belum memesan apapun, lakukan pemesanan dahulu")
	} else {
		fmt.Println("\nPesanan Anda (", len(pesanan), ")")
		fmt.Println("No | Pesanan     | Harga")
		fmt.Println("----+-------------+-----------------")
		for i := 0; i < len(pesanan); i++ {
			fmt.Printf("%-2d | %-11s | %-15s\n", i+1, pesanan[i], formatRupiah(hargaPesanan[i]))
		}
	}
}

func bayarPesanan() {
	if len(pesanan) <= 0 {
		fmt.Println("Anda belum memesan apapun, lakukan pemesanan dahulu")
	} else {

		sekarang := time.Now()
		tanggalSekarang := sekarang.Format("2006-01-02") // Format tanggal dalam format YYYY-MM-DD
		var dibayar int
		var jumlahPesanan = len(pesanan)
		var subTotal = totalCalc(hargaPesanan)
		var pajak = pajakCalc(len(pesanan))
		var total = totalCalc(hargaPesanan) + pajakCalc(len(pesanan))

		fmt.Println("\n\nWARUNG TOK ABAH")
		fmt.Println("Jl. Kampung Durian Runtuh")

		fmt.Println("No Pesanan: ", rand.Intn(101))
		fmt.Println("Tanggal: ", tanggalSekarang, "\n\n")

		fmt.Println("-----------------------------------------")
		for i := 0; i < len(pesanan); i++ {
			fmt.Printf("%-2d %-15s | %-15s\n", i+1, pesanan[i], formatRupiah(hargaPesanan[i]))

		}
		fmt.Println("-----------------------------------------")
		fmt.Printf("%-12s %-15s %-10d\n", " ", "Jumlah Pesanan", jumlahPesanan)
		fmt.Printf("%-12s %-15s %-10s\n", " ", "Sub Total", formatRupiah(subTotal))
		fmt.Printf("%-12s %-15s %-10s\n", " ", "Pajak", formatRupiah(pajak))
		fmt.Println("-----------------------------------------")
		fmt.Printf("%-12s %-15s %-10s\n", " ", "Total", formatRupiah(total))
		fmt.Println("-----------------------------------------")
		fmt.Printf("%-12s %-16s", " ", "Bayar")
		fmt.Scanln(&dibayar)
		fmt.Println("-----------------------------------------")
		fmt.Printf("%-12s %-15s %-10s\n", " ", "Kembalian", formatRupiah(dibayar-total))
		fmt.Println("-----------------------------------------")
		fmt.Println("Terima kasih sudah datang ke toko kami😊\n\n")

		pesanan = pesanan[:0]
		hargaPesanan = hargaPesanan[:0]
	}

}

func totalCalc(prices []int) int {
	sum := 0
	for _, num := range prices {
		sum += num
	}
	return sum
}

func pajakCalc(qty int) int {
	var pajak = 2500
	return qty * pajak
}

// Fungsi untuk memformat angka menjadi format rupiah
func formatRupiah(angka int) string {
	// Mengubah angka menjadi string
	strAngka := strconv.Itoa(angka)

	// Memisahkan string angka menjadi bagian-bagian dengan tanda ","
	parts := []string{}
	for len(strAngka) > 3 {
		parts = append(parts, strAngka[len(strAngka)-3:])
		strAngka = strAngka[:len(strAngka)-3]
	}
	parts = append(parts, strAngka)

	// Membalikkan urutan bagian-bagian angka
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	// Menggabungkan bagian-bagian angka dengan tanda "." untuk memisahkan ribuan
	formatted := strings.Join(parts, ".")

	// Menambahkan awalan "Rp " di depan angka
	formatted = "Rp " + formatted

	return formatted
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Untuk sistem Unix/Linux dan macOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Untuk sistem Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Tidak dapat membersihkan layar. Sistem operasi tidak didukung.")
	}
}
