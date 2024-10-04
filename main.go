package main

import (
	"fmt"
	"os"
)

// Fungsi untuk menampilkan "Hello World"
func tampilkanHelloWorld() {
	fmt.Println("Hello, World!")
}

// Fungsi untuk melakukan operasi matematika sederhana
func operasiMatematika() {
	var num1, num2 float64

	// Meminta input dari pengguna
	fmt.Println("Masukkan angka pertama: ")
	fmt.Scanf("%f\n", &num1)
	fmt.Scanln()
	fmt.Println("Masukkan angka kedua: ")
	fmt.Scanf("%f\n", &num2)
	fmt.Scanln()
	// Melakukan operasi aritmetika6

	penjumlahan := num1 + num2
	pengurangan := num1 - num2
	perkalian := num1 * num2
	var pembagian float64

	// Memeriksa apakah pembagian dengan nol terjadi
	if num2 != 0 {
		pembagian = num1 / num2
	} else {
		fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan!")
		fmt.Scanln()
		pembagian = 0
	}

	// Menampilkan hasil operasi
	fmt.Printf("\nPenjumlahan = %.2f\n", penjumlahan)
	fmt.Printf("Pengurangan = %.2f\n", pengurangan)
	fmt.Printf("Perkalian = %.2f\n", perkalian)
	fmt.Printf("Pembagian = %.2f\n", pembagian)
}

// Fungsi untuk menginput dan menampilkan data user
//func inputUserData() {

// Fungsi untuk menghitung faktorial
func hitungFaktorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * hitungFaktorial(n-1)
}

// func untuk menghitung faktorial
func tampilkanFaktorial() {
	var angka int
	fmt.Print("Masukkan angka untuk menghitung faktorial: ")
	fmt.Scan(&angka)

	if angka < 0 {
		fmt.Println("Error: Faktorial tidak terdefinisi untuk angka negatif.")
		return
	}

	hasil := hitungFaktorial(angka)
	fmt.Printf("Faktorial dari %d adalah %d\n", angka, hasil)
}

// Fungsi untuk menampilkan nama yang diinput oleh pengguna
func tampilkanNama() {
	var n int
	var input string

	fmt.Print("Berapa banyak nama yang ingin Anda masukkan? ")
	fmt.Scanln(&n)

	nama := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan nama ke-%d: ", i+1)
		fmt.Scanln(&input)
		nama[i] = input
	}

	// Menampilkan nama yang diinput
	fmt.Println("Daftar nama yang Anda masukkan:")
	for _, n := range nama {
		fmt.Println(n)
	}

}
func main() {
	for {
		fmt.Println("Hallo, selamat datang di CLI aplikasi")

		// Tampilkan menu utama
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tampilkan Hello World")
		fmt.Println("2. Operasi Matematika Sederhana")
		fmt.Println("3. Input dan Tampilkan Data User")
		fmt.Println("4. Tampilkan Faktorial")
		fmt.Println("5. Tampilkan variadic function")
		fmt.Println("6. Exit")
		fmt.Print("Pilih opsi : ")

		var opsi int
		fmt.Scan(&opsi)

		switch opsi {
		case 1:
			tampilkanHelloWorld()
		case 2:
			operasiMatematika()
		//case 3:
		//inputUserData()
		case 4:
			tampilkanFaktorial()
		case 5:
			tampilkanNama()
		case 6:
			fmt.Println("Keluar dari program...")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}

}
