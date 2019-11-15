package main

import "fmt"

var Pilih, x, y, Hasil float32	

func main(){
	for (Pilih != 5){
		fmt.Println("~ Menu Pilihan ~")
		fmt.Println("================")
		fmt.Println("1. X + Y")
		fmt.Println("2. X - Y")
		fmt.Println("3. X x Y")
		fmt.Println("4. X : Y")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih Menu ? ")
		fmt.Scan(&Pilih)
		fmt.Println()
		switch Pilih {
		case 1:
			fmt.Println("Menu Penjumlahan")
			fmt.Println("================")
			InputNilai()
			Hasil = x + y
			fmt.Println("Penjumlahan dari",x,"+",y,":", Hasil)
		case 2:
			fmt.Println("Menu Pengurangan")
			fmt.Println("================")
			InputNilai()
			Hasil = x + y
			fmt.Println("Pengurangan dari",x,"-",y,":", Hasil)
		case 3:
			fmt.Println("Menu Perkalian")
			fmt.Println("==============")
			InputNilai()
			Hasil = x * y
			fmt.Println("Perkalian dari",x,"x",y,":", Hasil)	
		case 4:
			fmt.Println("Menu Pembagian")
			fmt.Println("==============")
			InputNilai()
			Hasil = x / y
			fmt.Println("Pembagian dari",x,":",y,":", Hasil)
		case 5:
			fmt.Print("Terima Kasih!")
		
		default :
			fmt.Println("Maaf, yang anda masukkan salah. Ulangi!")
		}
		fmt.Println()
	}
}

func InputNilai()(float32, float32) {
	fmt.Print("Masukkan Nilai X : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan Nilai Y : ")		
	fmt.Scan(&y)
	return x, y
}

