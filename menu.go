package main

import (
	"fmt"
	"bufio"
	"os"
)

type User struct{
	NIM string
	Password string
	Nama string
}

var Pilih int;

var ListUser = []User{}

func main(){	
	var nim, password, nama string
	LengthList := 0;	
	scanner := bufio.NewReader(os.Stdin)						
	for (Pilih != 4){
		fmt.Println("~ Menu Pilihan ~")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Cari Data")
		fmt.Println("3. Login")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih ? ")
		fmt.Scanln(&Pilih)
		fmt.Println()		
		switch Pilih {
		case 1 :
			fmt.Println("~ Menu Tambah Data ~")
			fmt.Println("====================") 			
			fmt.Print("Masukkan NIM : ")
			nim, _= scanner.ReadString('\n')
			fmt.Print("Masukkan Password : ")
			password, _= scanner.ReadString('\n')
			fmt.Print("Masukkan Nama : ")
			nama, _= scanner.ReadString('\n')			
			ListUser = append(ListUser, User{nim, password, nama})
			LengthList = LengthList + 1;			
		case 2 :
			fmt.Println("~ Cari Data ~")
			fmt.Println("====================") 						
			fmt.Print("Masukkan NIM	: ")						
			nim, _= scanner.ReadString('\n')					
			HasilCari := false
			for i := 0; i < LengthList; i++ {
				if nim == ListUser[i].NIM{
					fmt.Print("NIM  : ", ListUser[i].NIM)
					fmt.Print("Nama : ", ListUser[i].Nama)
					fmt.Println("Password : ", ListUser[i].Password)
					HasilCari = true					
				}
			}
			if HasilCari == false{
				fmt.Println("Data Tidak Ditemukan!")
			}
		case 3 :
			Login := false			
			indeks := 0
			for indeks < 3 && Login == false{
				fmt.Println("~ Login ~")
				fmt.Println("====================")
				fmt.Print("Masukkan NIM 	  : ")
				nim, _= scanner.ReadString('\n')			
				fmt.Print("Masukkan Password  : ")			
				password, _= scanner.ReadString('\n')			
				for i := 0; i < LengthList; i++{
					if nim == ListUser[i].NIM && password == ListUser[i].Password{
						Login = true
					}else{						
						Login = false						
					}
				}				
				if Login == false{
					fmt.Println("Username atau Password salah!")
					indeks++
				}							
			}			
			if Login == true{
				fmt.Println("Anda berhasil Login")
			}else if indeks <= 3 && Login == false{
				fmt.Println("anda terblokir")				
			}			
		case 4 :	
			os.Exit(1)
		}		
	}	
}