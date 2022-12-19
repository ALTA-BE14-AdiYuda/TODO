package main

import (
	"bufio"
	"fmt"
	"os"
	"todo/activity"
	"todo/config"
	"todo/user"
)

func main() {
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}
	var authActvMenu = activity.ActivityMenu{DB: conn}
	var isLoggedIn bool

	for inputMenu != 0 {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var newUser user.User
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan password : ")
			fmt.Scanln(&newUser.Password)
			res, err := authMenu.Register(newUser)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res {
				fmt.Println("Sukses mendaftarkan data")
			} else {
				fmt.Println("Gagal mendaftarn data")
			}
		}

		if inputMenu == 2 {
			var newUser user.User
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan password : ")
			fmt.Scanln(&newUser.Password)
			res, idLogged, err := authMenu.Login(newUser)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res {
				fmt.Println("Login sukses")
				isLoggedIn = true
				for isLoggedIn {
					fmt.Println("1. Aktivitas")
					fmt.Println("2. Log out")
					fmt.Println("0. Exit")
					fmt.Scanln(&inputMenu)

					switch inputMenu {
					case 1:
						var newActivity activity.Activity
						reader := bufio.NewReader(os.Stdin) //standard input
						fmt.Print("Nama kegiatan : ")
						text, _ := reader.ReadString('\n')
						newActivity.Title = text
						fmt.Print("Lokasi : ")
						location, _ := reader.ReadString('\n')
						newActivity.Location = location
						newActivity.ID = idLogged
						res, err := authActvMenu.AddActivity(newActivity)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses tambah data")
						} else {
							fmt.Println("Gagal tambah data")
						}
					case 2:
						isLoggedIn = !isLoggedIn
					case 0:
						isLoggedIn = !isLoggedIn
					}

				}

			} else {
				fmt.Println("Gagal login")
			}
		}
	}
}
