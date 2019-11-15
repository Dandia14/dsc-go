package main

import "fmt"

func main(){
	const (
		Nim int = 10119033 
		Password int = 12345678
	)
	var (
		Username, KataSandi, Salah int
		Benar bool
	)
	fmt.Print("NIM      : ")
	fmt.Scan(&Username)

	fmt.Print("Password : ")
	fmt.Scan(&KataSandi)
	
	Benar = false
	Salah = 1
	for (Salah < 3 && !Benar){
		if Username == Nim && KataSandi == Password {
			Benar = true
		}else {
			if Username != Nim && KataSandi != Password{
				fmt.Println("Salah!\n")
			}
			fmt.Print("NIM      : ")
			fmt.Scan(&Username)

			fmt.Print("Password : ")
			fmt.Scan(&KataSandi)
			Salah++
		}
	}
	
	if (Benar){
		fmt.Println("Benar")
	}else{
		fmt.Println("Block!")
	}		
	
	
	
	
	// for i := 1;i < 3; i++{
	// 	if Nim == Nim && Password == Password {
	// 		fmt.Println("Benar!")	
	// 	} else{
	// 		fmt.Println("Salah!")
	// 	}
	// } 





}